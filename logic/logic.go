package logic

import "node/db"

type Status interface {
	Server() ServerStatusData
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
