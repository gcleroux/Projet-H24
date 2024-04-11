package factory

import (
	"log"

	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
	"github.com/gcleroux/Projet-H24/internal/game/components"
	"github.com/gcleroux/Projet-H24/internal/game/events"
	"github.com/gcleroux/Projet-H24/internal/networking"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateConnection(ecs *ecs.ECS) *donburi.Entry {
	conn := archetypes.Connection.Spawn(ecs)

	components.Connection.SetValue(conn, components.ConnectionData{
		WebSocketClient: networking.NewWebSocketClient(),
	})

	// Create the connection with the server
	c := components.Connection.Get(conn)
	err := c.Open("ws://localhost:8888/position")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			if pp, err := c.Read(); err == nil {
				log.Print("Message received")
				events.PeerUpdateEvent.Publish(ecs.World, events.PeerUpdate{
					PeerPosition: pp,
				})
				events.PeerUpdateEvent.ProcessEvents(ecs.World)
			}
		}
	}()

	return conn
}
