package model

import	"go.mongodb.org/mongo-driver/bson/primitive"

type Bill struct {
	ID	primitive.ObjectID	`bson:"_id"`
	Status 	bool
	Host	User
	Guests	[]User
	Items	[]Item
}

type Item struct {
	Id			int32
	Name		string
	Price		int32
	Owners		[]User
}

type User struct {
	Name	string
	Id		int32
}