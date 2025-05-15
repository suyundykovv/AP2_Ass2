// order-service/internal/nats/publisher.go
package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	pb "github.com/suyundykovv/protos/gen/go/events/v1"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) *Publisher {
	return &Publisher{conn: conn}
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
func (p *Publisher) PublishOrderCreated(event *pb.OrderEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.conn.Publish("orders.created", data)
}

func (p *Publisher) PublishOrderUpdated(event *pb.OrderEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.conn.Publish("orders.updated", data)
}
