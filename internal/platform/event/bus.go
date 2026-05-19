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

type Subscriber struct {
	UserID   string
	Username string
	IsAdmin  bool
	ch       chan Event
}

type Bus struct {
	mu          sync.RWMutex
	subscribers map[chan Event]*Subscriber
}

func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[chan Event]*Subscriber),
	}
}

func (b *Bus) Subscribe(userID, username string, isAdmin bool) chan Event {
	ch := make(chan Event, 10)

	b.mu.Lock()
	// Remove conexão anterior do mesmo usuário se existir
	for existingCh, sub := range b.subscribers {
		if sub.UserID == userID {
			delete(b.subscribers, existingCh)
			close(existingCh)
			break
		}
	}

	b.subscribers[ch] = &Subscriber{
		UserID:   userID,
		Username: username,
		IsAdmin:  isAdmin,
		ch:       ch,
	}
	b.mu.Unlock()

	return ch
}

func (b *Bus) Unsubscribe(ch chan Event) {
	b.mu.Lock()
	_, exists := b.subscribers[ch]
	if exists {
		delete(b.subscribers, ch)
		close(ch)
	}
	b.mu.Unlock()
}
func (b *Bus) ConnectedUsers() []*Subscriber {
	b.mu.RLock()
	defer b.mu.RUnlock()

	users := make([]*Subscriber, 0, len(b.subscribers))
	for _, s := range b.subscribers {
		users = append(users, s)
	}

	return users
}

func (b *Bus) Publish(e Event) {
	b.mu.RLock()

	subs := make([]chan Event, 0, len(b.subscribers))
	for ch := range b.subscribers {
		subs = append(subs, ch)
	}

	b.mu.RUnlock()

	for _, ch := range subs {
		select {
		case ch <- e:
		default:
		}
	}
}
