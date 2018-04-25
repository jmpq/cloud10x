package db

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	session *mgo.Session
	Db      *mgo.Database
}

func NewMongoDriver() (*Mongo, error) {
	return &Mongo{}, nil
}

func (db *Mongo) Init(path string, dbName string) error {
	session, err := mgo.Dial(path)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	db.session = session
	db.Db = session.DB(dbName)

	return nil
}

func (db *Mongo) Close() error {
	db.session.Close()

	return nil
}

func (db *Mongo) Insert(col string, docs ...interface{}) error {
	c := db.Db.C(col)
	err := c.Insert(docs...)

	return err
}

func (db *Mongo) Update(col string, query interface{}, docs interface{}) error {
	c := db.Db.C(col)
	err := c.Update(query, docs)

	return err
}

func (db *Mongo) Upsert(col string, query interface{}, docs interface{}) error {
	c := db.Db.C(col)
	_, err := c.Upsert(query, docs)

	return err
}

func (db *Mongo) One(col string, query interface{}, result interface{}) error {
	c := db.Db.C(col)

	err := c.Find(query).One(result)
	return err
}

func (db *Mongo) All(col string, query interface{}, result interface{}) error {
	c := db.Db.C(col)

	err := c.Find(query).All(result)
	return err
}

func (db *Mongo) Remove(col string, query interface{}) error {
	return db.Db.C(col).Remove(query)
}

func (db *Mongo) Count(col string, query interface{}) (int, error) {
	c := db.Db.C(col)
	return c.Find(query).Count()
}
