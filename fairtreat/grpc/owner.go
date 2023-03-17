package grpc

import (
	"context"
	"log"

	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"fairtreat.suwageeks.org/fairtreat/model"
)

func (s *Server) SetOwners(ctx context.Context, req *pb.SetItemOwnerRequest) (*pb.SetItemOwnerResponse, error) {
	// 入力チェック
	if (req.Id == "") {
		return &pb.SetItemOwnerResponse{
			Status: false,
		}, nil
	}

	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	// 明細の存在チェック
	objID, _ := primitive.ObjectIDFromHex(req.Id)
	itemId := req.ItemId
	
	cur, _ := coll.Aggregate(ctx, mongo.Pipeline{
		bson.D{{ // 対象のBillに絞る
			Key: "$match",
			Value: bson.D{{
				Key: "_id",
				Value: objID,
			}, {
				Key: "status",
				Value: true,
			}},
		}},
		bson.D{{ // 配列を展開
			Key: "$unwind",
			Value: "$items",
		}},
		bson.D{{ // itemIdで絞る
			Key: "$match",
			Value: bson.D{{
				Key: "items.id",
				Value: itemId,
			}},
		}},
		bson.D{{ // itemフィールドのみ取り出す
			Key: "$project",
			Value: bson.D{{
				Key: "items.owners",
				Value: 1,
			}, {
				Key: "_id",
				Value: 0,
			}},
		}},
	})
	if cur.RemainingBatchLength() == 0 {
		return &pb.SetItemOwnerResponse{
			Status: false,
		}, nil
	}

	// owners配列を生成
	owners := []model.User{}
	for _, v := range req.Owners {
		owners = append(owners, model.User{
			Id: v.Id,
			Name: v.Name,
		})
	}

	// データを上書き
	filter := bson.D{{
		Key: "_id",
		Value: objID,
	}, {
		Key: "items.id",
		Value: itemId,
	}}

	_, err := coll.UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"items.$.owners": owners,
		},
	})
	if err != nil {
		log.Println(err)
	}

	return &pb.SetItemOwnerResponse{
		Status: (err == nil),
	}, nil
}


func (s *Server) GetItemOwnersList(ctx context.Context, req *pb.GetItemOwnersRequest) (*pb.GetItemOwnersResponse, error) {
	// 入力チェック
	if (req.Id == "") {
		return &pb.GetItemOwnersResponse{
			Count: 0,
			Owners: nil,
		}, nil
	}

	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	// 明細の存在チェック
	objID, _ := primitive.ObjectIDFromHex(req.Id)
	itemId := req.ItemId
	
	cur, _ := coll.Aggregate(ctx, mongo.Pipeline{
		bson.D{{ // 対象のBillに絞る
			Key: "$match",
			Value: bson.D{{
				Key: "_id",
				Value: objID,
			}},
		}},
		bson.D{{ // 配列を展開
			Key: "$unwind",
			Value: "$items",
		}},
		bson.D{{ // itemIdで絞る
			Key: "$match",
			Value: bson.D{{
				Key: "items.id",
				Value: itemId,
			}},
		}},
		bson.D{{ // ownerフィールドのみ取り出す
			Key: "$project",
			Value: bson.D{{
				Key: "owners",
				Value: "$items.owners",
			}, {
				Key: "_id",
				Value: 0,
			}},
		}},
	})
	if cur.RemainingBatchLength() == 0 {
		return &pb.GetItemOwnersResponse{
			Count: 0,
			Owners: nil,
		}, nil
	}

	// 取得したドキュメントを型にデコード
	var result model.Owners
	for cur.Next(context.TODO()) {
		if err := cur.Decode(&result); err != nil {
			log.Println("[GetItemOweners] Error: Decode Object")
		}
	}
	cur.Close(context.TODO())

	// ownerリストを作成
	var owners []*pb.User
	for _, v := range result.Owners {
		owners = append(owners, &pb.User{
			Id: v.Id,
			Name: v.Name,
		})
	}

	// ownerの人数を取得
	var count int32 = int32(len(owners))

	return &pb.GetItemOwnersResponse{
		Count: count,
		Owners: owners,
	}, nil
}
