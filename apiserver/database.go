package main

import (
	"errors"
	"github.com/jmpq/cloud10x/apiserver/db"
)

type Database interface {
	Init(path string, dbName string) error
	Close() error

	Insert(col string, docs ...interface{}) error
	Update(col string, query interface{}, docs interface{}) error
	Upsert(col string, query interface{}, docs interface{}) error
	One(col string, query interface{}, result interface{}) error
	All(col string, query interface{}, result interface{}) error
	Remove(col string, query interface{}) error
}

func newDBDriver(t string) (Database, error) {
	if t == "mongo" {
		return db.NewMongoDriver()
	}

	return nil, errors.New("Dirver not found")
}
