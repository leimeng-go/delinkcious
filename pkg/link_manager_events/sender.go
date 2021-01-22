package link_manager_events

import (
	"log"

	"github.com/nats-io/go-nats"

	om "github.com/pingguodeli573365/delinkcious/pkg/object_model"
)

type eventSender struct {
	hostname string
	nats     *nats.EncodedConn
}

func (s *eventSender) OnLinkAdded(username string, link *om.Link) {
	event := Event{om.LinkAdded, username, link}
	log.Printf("[link manager events]OnLinkAdded(), sending to subject: %s, event: %v\n", event)
	err := s.nats.Publish(subject, event)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *eventSender) OnLinkUpdated(username string, link *om.Link) {
	err := s.nats.Publish(subject, Event{om.LinkUpdated, username, link})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *eventSender) OnLinkDeleted(username string, url string) {
	// Ignore link delete events
}

func NewEventSender(url string) (om.LinkManagerEvents, error) {
	ec, err := connect(url)
	if err != nil {
		return nil, err
	}
	return &eventSender{hostname: url, nats: ec}, nil
}
