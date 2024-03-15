package connections

type ConnectionHandler interface {
	Add(Connection)
	Remove(Connection)
	GetConns() []Connection
}
