package model

import	(
	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bill struct {
	ID	primitive.ObjectID	`bson:"_id"`
	Status 	bool			`bson:"status"`
	Host	User			`bson:"host"`
	Guests	[]User			`bson:"guests"`
	Items	[]Item			`bson:"items"`
}

type Item struct {
	Id			int32	`bson:"id"`
	Name		string	`bson:"name"`
	Price		int32	`bson:"price"`
	Owners		[]User	`bson:"owners"`
}

type User struct {
	Name	string	`bson:"name"`
	Id		int32	`bson:"id"`
}

type Owners struct {
	Owners	[]User	`bson:"owners"`
}

type PayPrice struct {
	User	User
	Price	int32
}

type ComfirmBill struct {
	ID			primitive.ObjectID	`bson:"_id"`
	Comfirm		[]*pb.PayPrice			`bson:"comfirm"`
}