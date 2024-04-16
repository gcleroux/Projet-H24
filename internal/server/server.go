package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	// co "github.com/gcleroux/Projet-H24/internal/networking/connections"

	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
)

// subscriber represents a Player
// Messages are sent on the msgs channel and if the client
// cannot keep up with the messages, kick is called.
type subscriber struct {
	msgs chan []byte
	kick func()
}

type gameServer struct {
	// Default to 16 messages. If the buffer is full, the player will be kicked
	subscriberMessageBuffer int

	updateLimiter *rate.Limiter

	logf func(f string, v ...interface{})

	serveMux http.ServeMux

	subscribersMu sync.Mutex
	subscribers   map[*subscriber]struct{}
}

func NewGameServer() *gameServer {
	gs := &gameServer{
		subscriberMessageBuffer: 16,
		updateLimiter:           rate.NewLimiter(rate.Every(time.Millisecond*10), 8),
		logf:                    log.Printf,
		subscribers:             make(map[*subscriber]struct{}),
	}

	gs.serveMux.Handle("/join", http.HandlerFunc(gs.joinHandler))
	gs.serveMux.Handle("/update", http.HandlerFunc(gs.updateHandler))

	return gs
}

func (gs *gameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gs.serveMux.ServeHTTP(w, r)
}

func (gs *gameServer) joinHandler(w http.ResponseWriter, r *http.Request) {
	gs.logf("Received /join call from %s", r.Host)

	err := gs.join(r.Context(), w, r)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		gs.logf("%v", err)
		return
	}
}

func (gs *gameServer) join(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var mu sync.Mutex
	var c *websocket.Conn
	var closed bool
	s := &subscriber{
		msgs: make(chan []byte, gs.subscriberMessageBuffer),
		kick: func() {
			mu.Lock()
			defer mu.Unlock()
			closed = true
			if c != nil {
				c.Close(
					websocket.StatusPolicyViolation,
					"connection too slow to keep up with messages",
				)
			}
		},
	}
	gs.addSubscriber(s)
	defer gs.deleteSubscriber(s)

	c2, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return err
	}
	mu.Lock()
	if closed {
		mu.Unlock()
		return net.ErrClosed
	}
	c = c2
	mu.Unlock()
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)

	for {
		select {
		case msg := <-s.msgs:
			err := writeTimeout(ctx, time.Second*5, c, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (gs *gameServer) updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow any domain
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allowed headers

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body := http.MaxBytesReader(w, r.Body, 8192)

	msg, err := io.ReadAll(body)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusRequestEntityTooLarge),
			http.StatusRequestEntityTooLarge,
		)
		return
	}
	// Broadcast the PlayerUpdate to all subscribers
	gs.update(msg)

	w.WriteHeader(http.StatusAccepted)
}

func (gs *gameServer) update(msg []byte) {
	gs.subscribersMu.Lock()
	defer gs.subscribersMu.Unlock()

	gs.updateLimiter.Wait(context.Background())

	// It never blocks and so messages to slow subscribers
	// are dropped. If the msgs buffer is full, the subscribers
	// will be kicked
	for s := range gs.subscribers {
		select {
		case s.msgs <- msg:
		default:
			go s.kick()
		}
	}
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}

// addSubscriber registers a subscriber.
func (gs *gameServer) addSubscriber(s *subscriber) {
	gs.subscribersMu.Lock()
	gs.subscribers[s] = struct{}{}
	gs.subscribersMu.Unlock()
}

// deleteSubscriber deletes the given subscriber.
func (gs *gameServer) deleteSubscriber(s *subscriber) {
	gs.subscribersMu.Lock()
	delete(gs.subscribers, s)
	gs.subscribersMu.Unlock()
}
