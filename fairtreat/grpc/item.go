package grpc

import (
	"context"
	"fmt"

	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fairtreat.suwageeks.org/fairtreat/model"
)

func (s *Server) GetItemsList(ctx context.Context, req *pb.GetItemsListRequest) (*pb.GetItemsListResponse, error) {
	// 入力チェック
	if (req.Id == "") {
		return &pb.GetItemsListResponse{
			Count: 0,
			Items: nil,
		}, nil
	}
	
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	// 明細の存在チェック
	var result model.Bill
	objID, _ := primitive.ObjectIDFromHex(req.Id)
	err := coll.FindOne(context.TODO(), bson.D{{
		Key: "_id", 
		Value: objID,
	}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return &pb.GetItemsListResponse{
			Count: 0,
			Items: nil,
		}, nil
	}

	// 商品リストを作成
	var items []*pb.Item
	for _, v := range result.Items {
		var owners []*pb.User
		for _, w := range v.Owners {
			owners = append(owners, &pb.User{
				Id: w.Id,
				Name: w.Name,
			})
		}

		items = append(items, &pb.Item{
			Id: v.Id,
			Name: v.Name,
			Price: v.Price,
			Owners: owners,
		})
	}

	// 商品の個数を取得
	var count int32 = int32(len(items))

	return &pb.GetItemsListResponse{
		Count: count,
		Items: items,
	}, nil
}
