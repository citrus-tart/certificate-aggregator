package repository

import (
	"github.com/citrus-tart/certificate-aggregator/certificate"
)

type Repository struct {
	Certs map[string]certificate.Certificate
}

func (r Repository) GetById(id string) certificate.Certificate {
	return r.Certs[id]
}

func (r Repository) Save(c certificate.Certificate) certificate.Certificate {
	r.Certs[c.ID] = c
	return c
}

func New() Repository {
	var r Repository
	r.Certs = make(map[string]certificate.Certificate)
	return r
}
