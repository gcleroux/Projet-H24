package networking

// type Connection interface {
// 	Open(string) error
// 	Close() error
//
// 	Write(api.PlayerPositionMessage) error
// 	Read() (api.PlayerPositionMessage, error)
// }

type Connection interface {
	OnConnect()
	OnDisconnect()
	OnError()
	OnMessage()
}
