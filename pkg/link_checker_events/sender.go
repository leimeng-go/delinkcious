package link_checker_events

import (
	"log"

	"github.com/nats-io/go-nats"
	om "github.com/pingguodeli573365/delinkcious/pkg/object_model"
)

type eventSender struct {
	hostname string
	nats     *nats.EncodedConn
}

func (s *eventSender) OnLinkChecked(username string, url string, status om.LinkStatus) {
	err := s.nats.Publish(subject, Event{username, url, status})
	if err != nil {
		log.Fatal(err)
	}
}

func NewEventSender(url string) (om.LinkCheckerEvents, error) {
	ec, err := connect(url)
	if err != nil {
		return nil, err
	}
	return &eventSender{hostname: url, nats: ec}, nil
}
