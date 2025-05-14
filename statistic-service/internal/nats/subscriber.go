package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"statistic-service/internal/service"
	"time"

	"github.com/nats-io/nats.go"
	pb "github.com/suyundykovv/protos/gen/go/events/v1"
)

type Subscriber struct {
	service service.StatisticService
	conn    *nats.Conn
}

func NewSubscriber(svc service.StatisticService, conn *nats.Conn) *Subscriber {
	return &Subscriber{service: svc, conn: conn}
}

// ConnectWithRetry attempts to connect to NATS with retries
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
func (s *Subscriber) Subscribe() error {
	if _, err := s.conn.Subscribe("orders.*", func(msg *nats.Msg) {
		var event pb.OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Error unmarshaling order event: %v", err)
			return
		}

		switch event.EventType {
		case "created":
			if err := s.service.ProcessOrderCreated(&event); err != nil {
				log.Printf("Error processing order created event: %v", err)
			}
		case "updated":
			if err := s.service.ProcessOrderUpdated(&event); err != nil {
				log.Printf("Error processing order updated event: %v", err)
			}
		case "deleted":
			if err := s.service.ProcessOrderDeleted(&event); err != nil {
				log.Printf("Error processing order deleted event: %v", err)
			}
		}
	}); err != nil {
		return err
	}

	// Subscribe to inventory events
	if _, err := s.conn.Subscribe("inventory.*", func(msg *nats.Msg) {
		var event pb.InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Error unmarshaling inventory event: %v", err)
			return
		}

		if err := s.service.ProcessInventoryEvent(&event); err != nil {
			log.Printf("Error processing inventory event: %v", err)
		}
	}); err != nil {
		return err
	}

	return nil
}
