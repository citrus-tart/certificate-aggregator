package processor

import (
	"encoding/json"

	"github.com/citrus-tart/certificate-aggregator/certificate"
	"github.com/citrus-tart/certificate-aggregator/events"
)

type Repository interface {
	GetById(string) certificate.Certificate
	Save(certificate.Certificate) certificate.Certificate
}

type processor struct {
	repo Repository
}

type certificateCreatedPayload struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Authority   string `json:"authority"`
}

type certificateUpdatedPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Authority   string `json:"authority"`
}

func (p processor) ProcessEvent(ev events.Event) {
	var c certificate.Certificate

	switch ev.Name {
	case "certificate:created":
		// unmarshal payload in to struct
		var payload certificateCreatedPayload
		json.Unmarshal([]byte(ev.Payload), &payload)

		// update object
		c.ID = payload.ID
		c.Name = payload.Name
		c.Description = payload.Description
		c.Authority = payload.Authority
		c.Created = ev.Timestamp

		// store
		p.repo.Save(c)

	case "certificate:updated":
		// get from repo
		c = p.repo.GetById(ev.EntityID)

		// unmarshal payload in to struct
		var payload certificateUpdatedPayload
		json.Unmarshal([]byte(ev.Payload), &payload)

		// update object
		c.ID = ev.EntityID
		c.Name = payload.Name
		c.Description = payload.Description
		c.Authority = payload.Authority

		// store
		p.repo.Save(c)
	}
}

func New(r Repository) processor {
	var p processor
	p.repo = r
	return p
}
