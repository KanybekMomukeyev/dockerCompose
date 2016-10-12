package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `form:"title" json:"title"`
	Description string        `form:"description" json:"description"`
	Price       float64       `form:"price" json:"price"`
}