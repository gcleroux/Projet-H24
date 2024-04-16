package connections

type Connection interface {
	Dial(addr string) error

	// Close the connection.
	Close() error

	Read() ([]byte, error)

	Write(msg []byte) error

	// Get the underlying raw connection object for advanced usage if needed.
	Raw() interface{}
}
