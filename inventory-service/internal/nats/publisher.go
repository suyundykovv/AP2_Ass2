// inventory-service/internal/nats/publisher.go
package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	eventspb "github.com/suyundykovv/protos/gen/go/events/v1"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) *Publisher {
	return &Publisher{conn: conn}
}

func (p *Publisher) PublishInventoryCreated(event *eventspb.InventoryEvent) error {
	return p.publishEvent("inventory.created", event)
}

func (p *Publisher) PublishInventoryUpdated(event *eventspb.InventoryEvent) error {
	return p.publishEvent("inventory.updated", event)
}

func (p *Publisher) PublishInventoryDeleted(event *eventspb.InventoryEvent) error {
	return p.publishEvent("inventory.deleted", event)
}

func (p *Publisher) publishEvent(subject string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	if err := p.conn.Publish(subject, data); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	log.Printf("Published inventory event to %s: %+v", subject, event)
	return nil
}
func Connect(url string, maxAttempts int, delay time.Duration) (*nats.Conn, error) {
	var (
		conn *nats.Conn
		err  error
	)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		conn, err = nats.Connect(url)
		if err == nil {
			return conn, nil
		}

		if attempt < maxAttempts {
			log.Printf("Failed to connect to NATS (attempt %d/%d): %v. Retrying in %v...",
				attempt, maxAttempts, err, delay)
			time.Sleep(delay)
		}
	}

	return nil, fmt.Errorf("after %d attempts, last error: %w", maxAttempts, err)
}
