package models

import "gopkg.in/mgo.v2/bson"

type Restaurant struct {
	ID      bson.ObjectId `json:"id, omitempty" bson:"_id"`
	Name    string        `json:"name, omitempty" bson:"name"`
	Address *Address      `json:"address, omitempty" bson:"address"`
}

type Address struct {
	City   string `json:"city, omitempty" bson:"city"`
	Street string `json:"street, omitempty" bson:"street"`
}
