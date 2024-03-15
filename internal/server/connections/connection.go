package connections

import (
	"context"
	"net/http"
)

type Connection interface {
	// Open the connection
	Open(ctx context.Context, w http.ResponseWriter, r *http.Request) error

	// Close the connection.
	Close() error

	// Get the underlying raw connection object for advanced usage if needed.
	Raw() interface{}
}
