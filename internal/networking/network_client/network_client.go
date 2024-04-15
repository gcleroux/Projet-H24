package network_client

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
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
	Conn  *co.WSConnection
	Peers map[uuid.UUID]api.PlayerPosition

	// The WebSocketClient will publish the peer position it gets from the server
	Publisher[api.PlayerPosition]

	// The WebSocketClient will subscribe to the Player inputs to send to the server
	Subscriber[api.PlayerPosition]
}

func newNetworkClient() *NetworkClient {
	n := &NetworkClient{
		Conn:  co.NewWSConnection(),
		Peers: make(map[uuid.UUID]api.PlayerPosition),

		Publisher:  NewPublisher[api.PlayerPosition](),
		Subscriber: NewSubscriber[api.PlayerPosition](Cfg.GetInt("connection.buffer_size")),
	}

	dial_addr := fmt.Sprintf(
		"wss://%s:%d%s",
		Cfg.GetString("remote_address"),
		Cfg.GetInt("port"),
		Cfg.GetString("route"),
	)

	n.Conn.Dial(dial_addr)

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

func (n *NetworkClient) Close() error {
	return n.Conn.Close()
}

func (n *NetworkClient) Send(msg api.PlayerPosition) error {
	// Headers for the message
	msg.ID = n.Conn.ID
	msg.ClientT = time.Now().UnixMilli()

	return n.Conn.Send(msg)
}

func (n *NetworkClient) Read() (api.PlayerPosition, error) {
	return n.Conn.Read()
}
