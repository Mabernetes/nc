package logic

import (
	"github.com/docker/docker/api/types"
	"node/db"
)

type Status interface {
	Server() ServerStatusData
	Deployment(deployment, pod string) ([]types.Container, error)
}

type Logic struct {
	db     *db.DB
	Status Status
}

func NewLogic(db *db.DB) *Logic {
	return &Logic{
		db:     db,
		Status: NewStatusLogic(),
	}
}
