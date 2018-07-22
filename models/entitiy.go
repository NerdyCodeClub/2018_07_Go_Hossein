package models

import "gopkg.in/mgo.v2/bson"

// Entity represents a company ...
type Entity struct {
	ID                  bson.ObjectId `bson:"_id" json:"id"`
	Name                string        `bson:"name" json:"name"`
	IncorporationNumber string        `bson:"incorporation_number" json:"incorporation_number"`
}
