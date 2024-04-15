package connections

import (
	"github.com/gcleroux/Projet-H24/api/v1"
)

type Connection interface {
	Dial(addr string)

	// Close the connection.
	Close() error

	Read() (api.PlayerPosition, error)

	Send(msg api.PlayerPosition) error

	// Get the underlying raw connection object for advanced usage if needed.
	Raw() interface{}
}
