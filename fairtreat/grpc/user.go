package grpc

import (
	"context"
	"fmt"
	"math/rand"

	"fairtreat.suwageeks.org/fairtreat/model"

	pb "fairtreat.suwageeks.org/fairtreat/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	// 入力チェック
	if req.Name == "" {
		return &pb.AddUserResponse{
			Status: false,
			Guest: nil,
		}, nil
	}
	
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	var result model.Bill
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Println("Error: id => ObjId")
		fmt.Println(err)
	}
	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}}

	// Bill存在チェック
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error: ObjectDecode")
		fmt.Println(err)
		return &pb.AddUserResponse{
			Status: false,
			Guest: nil,
		}, nil
	}

	var userId int32 = rand.Int31() // userIdの生成
	guest := model.User{
		Name: req.Name,
		Id: userId,
	}
	
	_, err = coll.UpdateOne(context.TODO(), filter, bson.D{{
		Key: "$push",
		Value: bson.D{{
			Key: "Guests",
			Value: guest,
		}},
	}})
	if err != nil {
		fmt.Println(err)
	}

	return &pb.AddUserResponse{
		Status: (err == nil),
		Guest: &pb.User{
			Name: guest.Name,
			Id: guest.Id,
		},
	}, nil
}

func (s *Server) GetUsersList(ctx context.Context, req *pb.GetUsersListRequest) (*pb.GetUsersListResponse, error) {
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	var result model.Bill
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Println("Error: id => ObjId")
		fmt.Println(err)
	}
	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}}

	// Bill存在チェック
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error: ObjectDecode")
		fmt.Println(err)
		return &pb.GetUsersListResponse{
			Count: 0,
			Users: nil,
		}, nil
	}

	// ユーザーリストを作成
	var users []*pb.User
	for _, v := range result.Guests {
		users = append(users, &pb.User{
			Id: v.Id,
			Name: v.Name,
		})
	}

	// 現在の参加者数を取得
	var count int32 = int32(len(users))

	return &pb.GetUsersListResponse{
		Count: count,
		Users: users,
	}, nil
}
