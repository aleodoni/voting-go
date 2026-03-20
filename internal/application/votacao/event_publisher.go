package votacao

import (
	"github.com/aleodoni/go-ddd/domain"
	"github.com/aleodoni/voting-go/internal/platform/event"
)

// publishEvents publica no barramento todos os eventos retornados por PullEvents(),
// usando o mapper fornecido para converter cada domain event em um event.Event.
// Eventos não mapeados (mapper retorna ok=false) são silenciosamente ignorados.
func publishEvents(
	bus *event.Bus,
	events []domain.DomainEvent,
	mapper func(domain.DomainEvent) (event.Event, bool),
) {
	for _, e := range events {
		if mapped, ok := mapper(e); ok {
			bus.Publish(mapped)
		}
	}
}
