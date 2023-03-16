package grpc

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"fairtreat.suwageeks.org/fairtreat/model"
)

func (s *Server) CreateBill(ctx context.Context, req *pb.CreateBillRequest) (*pb.CreateBillResponse, error) {
	// コレクションを取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	hostName := req.HostName
	var hostId int32 = rand.Int31()

	// アイテム情報の整理
	items := []model.Item{}
	for _, v := range req.Items {
		items = append(items, model.Item{
			Id: v.Id,
			Name: v.Name,
			Price: v.Price,
			Owners: []model.User{{
				Id: hostId,
				Name: hostName,
			}},
		})
	}

	// 初期データの挿入
	bill := model.Bill{
		ID: primitive.NewObjectID(),
		Status: true,
		Host: model.User{
			Name: hostName,
			Id: hostId,
		},
		Guests: nil,
		Items: items,
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
