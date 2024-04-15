package factory

//
// import (
// 	"fmt"
// 	"log"
//
// 	"github.com/gcleroux/Projet-H24/internal/game/archetypes"
// 	"github.com/gcleroux/Projet-H24/internal/game/components"
// 	"github.com/gcleroux/Projet-H24/internal/game/events"
// 	"github.com/gcleroux/Projet-H24/internal/networking"
// 	"github.com/spf13/viper"
// 	"github.com/yohamta/donburi"
// 	"github.com/yohamta/donburi/ecs"
// )
//
// func CreateConnection(ecs *ecs.ECS) *donburi.Entry {
// 	conn := archetypes.Connection.Spawn(ecs)
//
// 	components.Connection.SetValue(conn, components.ConnectionData{
// 		WebSocketClient: networking.NewWebSocketClient(),
// 	})
//
// 	// Create the connection with the server
// 	c := components.Connection.Get(conn)
// 	remote_addr := fmt.Sprintf(
// 		"ws://%s:%d%s",
// 		viper.GetString("server.address"),
// 		viper.GetInt("server.port"),
// 		viper.GetString("server.route"),
// 	)
// 	if err := c.Open(remote_addr); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	go func() {
// 		for {
// 			if pp, err := c.Read(); err == nil {
// 				events.PeerUpdateEvent.Publish(ecs.World, events.PeerUpdate{
// 					PeerPosition: pp,
// 				})
// 			}
// 		}
// 	}()
//
// 	return conn
// }
