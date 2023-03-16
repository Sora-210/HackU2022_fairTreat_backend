package grpc

import (
	"context"
	"fmt"
	"log"

	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"fairtreat.suwageeks.org/fairtreat/model"
)

func (s *Server) CreateBill(ctx context.Context, req *pb.CreateBillRequest) (*pb.CreateBillResponse, error) {
	// コレクションを取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	hostName := req.HostName
	var hostId int32 = 0

	// 初期データの挿入
	bill := model.Bill{
		primitive.NewObjectID(),
		true,
		model.User{
			hostName,
			hostId,
		},
		nil,
		nil,
	}
	_, err := coll.InsertOne(context.Background(), bill)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to insert initial data.")
	}

	// オブジェクト発行完了
	
	return &pb.CreateBillResponse{
		BillId: bill.ID.Hex(),
		Host: &pb.User{
			Id: bill.Host.Id,
			Name: bill.Host.Name,
		},
	}, nil
}
