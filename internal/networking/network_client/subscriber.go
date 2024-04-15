package network_client

import (
	"log"

	"github.com/gcleroux/Projet-H24/api/v1"
)

type Subscriber[T api.PlayerPosition | api.PeerPosition] struct {
	Chan chan T
}

func NewSubscriber[T api.PlayerPosition | api.PeerPosition](size int) Subscriber[T] {
	return Subscriber[T]{
		Chan: make(chan T, size),
	}
}

func (s *Subscriber[T]) Listen(callback func(T) error) {
	for msg := range s.Chan {
		if err := callback(msg); err != nil {
			log.Println(err)
		}
	}
}
