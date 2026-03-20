package event

import "sync"

// EventType representa o tipo de um evento publicado no bus.
type EventType string

const (
	VotacaoAberta    EventType = "votacao_aberta"
	VotacaoFechada   EventType = "votacao_fechada"
	VotacaoCancelada EventType = "votacao_cancelada"
	VotoRegistrado   EventType = "voto_registrado"
)

// Event representa um evento publicado no bus.
type Event struct {
	Type    EventType
	Payload any
}

// Subscriber representa um usuário conectado ao bus de eventos.
type Subscriber struct {
	UserID   string
	Username string
	IsAdmin  bool
	ch       chan Event
}

// Bus é um event bus simples para publicação e assinatura de eventos dentro da aplicação.
//
// Comportamento:
//   - múltiplos subscribers podem se inscrever e receber eventos simultaneamente
//   - eventos publicados para clientes lentos são descartados para não bloquear o bus
type Bus struct {
	mu          sync.RWMutex
	subscribers map[chan Event]*Subscriber
}

// NewBus cria uma nova instância de [Bus].
func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[chan Event]*Subscriber),
	}
}

// Subscribe registra um novo subscriber no bus e retorna o canal de eventos.
func (b *Bus) Subscribe(userID, username string, isAdmin bool) chan Event {
	ch := make(chan Event, 10)
	b.mu.Lock()
	b.subscribers[ch] = &Subscriber{
		UserID:   userID,
		Username: username,
		IsAdmin:  isAdmin,
		ch:       ch,
	}
	b.mu.Unlock()
	return ch
}

// Unsubscribe remove o subscriber do bus e fecha o canal associado.
func (b *Bus) Unsubscribe(ch chan Event) {
	b.mu.Lock()
	delete(b.subscribers, ch)
	close(ch)
	b.mu.Unlock()
}

// Publish envia o evento para todos os subscribers registrados.
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

// ConnectedUsers retorna a lista de todos os subscribers atualmente conectados.
func (b *Bus) ConnectedUsers() []*Subscriber {
	b.mu.RLock()
	defer b.mu.RUnlock()
	users := make([]*Subscriber, 0, len(b.subscribers))
	for _, s := range b.subscribers {
		users = append(users, s)
	}
	return users
}

// ConnectedAdmins retorna a lista de subscribers conectados que possuem perfil de administrador.
func (b *Bus) ConnectedAdmins() []*Subscriber {
	b.mu.RLock()
	defer b.mu.RUnlock()
	var admins []*Subscriber
	for _, s := range b.subscribers {
		if s.IsAdmin {
			admins = append(admins, s)
		}
	}
	return admins
}

// ConnectedCount retorna o número total de subscribers atualmente conectados.
func (b *Bus) ConnectedCount() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.subscribers)
}
