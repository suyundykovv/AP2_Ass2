package nats

import (
	"encoding/json"
	"log"
	"statistic-service/internal/service"

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
