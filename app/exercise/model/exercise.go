package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Exercise struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name"`
	Tags      []string      `json:"tags"`
	Question  string        `json:"question"`
	Variables []*Variable   `json:"variables"`
	Hints []*Hint           `json:"hints"`
	Answer *Answer          `json:"answer"`
}
