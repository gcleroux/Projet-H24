package network_client

import "github.com/gcleroux/Projet-H24/api/v1"

type Publisher[T api.PlayerPosition | api.PeerPosition] struct {
	Subs []Subscriber[T]
}

func NewPublisher[T api.PlayerPosition | api.PeerPosition]() Publisher[T] {
	return Publisher[T]{
		Subs: make([]Subscriber[T], 0),
	}
}

func (p *Publisher[T]) AddSubscriber(sub Subscriber[T]) {
	p.Subs = append(p.Subs, sub)
}

func (p *Publisher[T]) Broadcast(msg T) {
	for _, sub := range p.Subs {
		sub.Chan <- msg
	}
}
