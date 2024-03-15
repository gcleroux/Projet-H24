package networking

import "github.com/gcleroux/Projet-H24/internal/game/characters"

// The networkHandler sends the player position to the server
// and receives updates from the server about the opponents positions.
type NetworkHandler interface {
	Open() error
	Close() error

	SendPlayerPosition(x, y float64) error
	ReadPeerPosition() error
	Peers() []characters.Peer
}
