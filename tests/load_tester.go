package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/gcleroux/Projet-H24/api/v1"
	co "github.com/gcleroux/Projet-H24/internal/networking/connections"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var config []byte

// Testing for 20 concurrents connections
const N int = 100

func init() {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	for i := 0; i < N; i++ {
		loadTester := newLoadTester()
		go loadTester.Test()
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}

type LoadTester struct {
	Conn *co.WSConnection
}

func newLoadTester() *LoadTester {
	lt := &LoadTester{
		Conn: co.NewWSConnection(),
	}

	dial_addr := fmt.Sprintf(
		"ws://%s:%d%s",
		viper.GetString("remote_address"),
		viper.GetInt("port"),
		viper.GetString("route"),
	)

	lt.Conn.Dial(dial_addr)

	go func() {
		for {
			// Ignore the message value
			lt.Read()
		}
	}()

	return lt
}

func (l *LoadTester) Close() error {
	return l.Conn.Close()
}

func (l *LoadTester) Send(msg api.PlayerPosition) error {
	// Headers for the message
	msg.ID = l.Conn.ID
	msg.ClientT = time.Now().UnixMilli()

	return l.Conn.Send(msg)
}

func (l *LoadTester) Read() (api.PlayerPosition, error) {
	return l.Conn.Read()
}

func (l *LoadTester) Test() {
	for {
		// Send a message at 60Hz
		time.Sleep(time.Second / 60)
		l.Send(api.PlayerPosition{})
	}
}
