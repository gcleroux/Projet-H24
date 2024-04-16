package network_client

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	api "github.com/gcleroux/Projet-H24/api/v1"
	co "github.com/gcleroux/Projet-H24/internal/networking/connections"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var config []byte

var (
	Cfg       *viper.Viper
	NetClient *NetworkClient
)

func init() {
	Cfg = viper.New()
	Cfg.SetConfigType("yaml")
	err := Cfg.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		log.Fatal(err)
	}
	NetClient = newNetworkClient()
}

type NetworkClient struct {
	ID    uuid.UUID
	Conn  co.Connection
	Peers map[uuid.UUID]api.PlayerPosition

	// The WebSocketClient will publish the peer position it gets from the server
	Publisher[api.PlayerPosition]

	// The WebSocketClient will subscribe to the Player inputs to send to the server
	Subscriber[api.PlayerPosition]
}

func newNetworkClient() *NetworkClient {
	n := &NetworkClient{
		ID:    uuid.New(),
		Conn:  co.NewWSConnection(),
		Peers: make(map[uuid.UUID]api.PlayerPosition),

		Publisher:  NewPublisher[api.PlayerPosition](),
		Subscriber: NewSubscriber[api.PlayerPosition](Cfg.GetInt("connection.buffer_size")),
	}

	dial_addr := fmt.Sprintf(
		"ws://%s:%d%s",
		Cfg.GetString("remote_address"),
		Cfg.GetInt("port"),
		Cfg.GetString("route"),
	)

	if err := n.Conn.Dial(dial_addr); err != nil {
		log.Fatal(err)
	}

	// Listen for updates from the Player input
	go n.Listen(n.Send)

	// No need for sync, at worst the client will be behind by a frame
	go func() {
		for {
			msg, err := n.Read()
			if err == nil {
				msg.TotalT = time.Now().UnixMilli()
				n.Peers[msg.ID] = msg
			}
		}
	}()

	return n
}

func (n *NetworkClient) Close() {
	n.Conn.Close()
}

func (n *NetworkClient) Send(msg api.PlayerPosition) error {
	// Populate the necessary headers or fields in the message
	msg.ID = n.ID
	msg.ClientT = time.Now().UnixMilli()

	// Marshal the PlayerPosition struct to JSON
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Create an HTTP client
	client := &http.Client{}

	// Construct the POST request
	req, err := http.NewRequest("POST", "http://localhost:8888/update", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	// Set the appropriate headers (e.g., Content-Type)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusAccepted {
		// Read the response body for more information (optional)
		respData, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned an error: %s, body: %s", resp.Status, respData)
	}

	return nil
}

func (n *NetworkClient) Read() (api.PlayerPosition, error) {
	var msg api.PlayerPosition

	data, err := n.Conn.Read()
	if err != nil {
		return api.PlayerPosition{}, err
	}

	err = json.Unmarshal(data, &msg)
	if err != nil {
		return api.PlayerPosition{}, err
	}

	return msg, nil
}
