// Package event implements a simple event bus for publishing and subscribing to events within the application.
package event

import "sync"

type EventType string

const (
	VotacaoAberta    EventType = "votacao_aberta"
	VotacaoFechada   EventType = "votacao_fechada"
	VotacaoCancelada EventType = "votacao_cancelada"
	VotoRegistrado   EventType = "voto_registrado"
)

type Event struct {
	Type    EventType
	Payload any
}

type Bus struct {
	mu          sync.RWMutex
	subscribers map[chan Event]struct{}
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[chan Event]struct{}),
	}
}

func (b *Bus) Subscribe() chan Event {
	ch := make(chan Event, 10)
	b.mu.Lock()
	b.subscribers[ch] = struct{}{}
	b.mu.Unlock()
	return ch
}

func (b *Bus) Unsubscribe(ch chan Event) {
	b.mu.Lock()
	delete(b.subscribers, ch)
	close(ch)
	b.mu.Unlock()
}

func (b *Bus) Publish(e Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for ch := range b.subscribers {
		select {
		case ch <- e:
		default:
			// cliente lento, descarta o evento para não bloquear
		}
	}
}
