package db

import "github.com/uptrace/bun"

type DB struct {
	conn *bun.DB
}

func NewDB(conn *bun.DB) *DB {
	return &DB{conn: conn}
}
