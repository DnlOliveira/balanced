package pkg

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Server   string
	Database string
}

var db *mgo.Database
var c *mgo.Collection

func (m *Mongo) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	c = db.C("users")
	fmt.Println("mongo connected")

	// db := session.DB("book-inventory").C("inventory")
	// err = db.Insert(&Book{"Masters", "Daniel"})
	// if err != nil {
	// 	panic(err)
	// }
}

func (m *Mongo) Find(key string, p *string) User {
	u := User{}
	err := c.Find(bson.M{key: p}).One(&u)
	if err != nil {
		panic(err)
	}
	return u
}
