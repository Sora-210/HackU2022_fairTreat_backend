package grpc

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Server) GetBill(ctx context.Context, req *pb.GetBillRequest) (*pb.GetBillResponse, error) {
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	var result pb.Bill
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Println("Error: id => ObjId")
		fmt.Println(err)
	}
	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}}

	// Bill取得
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error: ObjectDecode")
		fmt.Println(err)
		return &pb.GetBillResponse{
			Bill: nil,
		}, nil
	}

	return &pb.GetBillResponse{
		Bill: &result,
	}, nil
}

func (s *Server) ConfirmBill(ctx context.Context, req *pb.ConfirmBillRequest) (*pb.ConfirmBillResponse, error) {
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	var result pb.Bill
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Println("Error: id => ObjId")
		fmt.Println(err)
	}
	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}, {
		Key: "status",
		Value: true,
	}}

	// Bill取得
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error: ObjectDecode")
		fmt.Println(err)
		return &pb.ConfirmBillResponse{
			Status: false,
		}, nil
	}

	// 処理用配列を作成
	comfirmPrices := map[int32]*model.PayPrice{}
	comfirmPrices[result.Host.Id] = &model.PayPrice{
		User: model.User{
			Id: result.Host.Id,
			Name: result.Host.Name,
		},
		Price: 0,
	}
	for _, v := range result.Guests {
		comfirmPrices[v.Id] = &model.PayPrice{
			User: model.User{
				Id: v.Id,
				Name: v.Name,
			},
			Price: 0,
		}
	}

	// 金額計算処理
	for _, v := range result.Items {
		fmt.Println(v)
		var price int32 = v.Price
		var count int32 = int32(len(v.Owners))
		var priceOfPeople int32 = price / count
		var priceOfMod int32 = price - (priceOfPeople * count)

		flag := true
		for _, w := range v.Owners {
			if _, ok := comfirmPrices[w.Id]; !ok {
				return &pb.ConfirmBillResponse{
					Status: false,
				}, nil
			}
			if flag {
				comfirmPrices[w.Id].Price += priceOfMod
				flag = false
			}
			comfirmPrices[w.Id].Price += priceOfPeople
		}
	}

	// 確定データ変換
	var comfirmBill []*pb.PayPrice
	for _, v := range comfirmPrices {
		comfirmBill = append(comfirmBill, &pb.PayPrice{
			User: &pb.User{
				Id: v.User.Id,
				Name: v.User.Name,
			},
			Price: v.Price,
		})
	}

	// データ登録
	_, err = coll.UpdateOne(context.TODO(), filter, bson.D{{
		Key: "$set",
		Value: bson.D{{
			Key: "comfirm",
			Value: comfirmBill,
		}},
	}, {
		Key: "$set",
		Value: bson.D{{
			Key: "status",
			Value: false,
		}},
	}})
	if err != nil {
		fmt.Println(err)
	}

	return &pb.ConfirmBillResponse{
		Status: (err == nil),
	}, nil
}

func (s *Server) GetConfirmBill(ctx context.Context, req *pb.GetConfirmBillRequest) (*pb.GetConfirmBillResponse, error) {
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")
	var result model.ComfirmBill
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		fmt.Println("Error: id => ObjId")
		fmt.Println(err)
	}
	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}, {
		Key: "status",
		Value: false,
	}}

	// ComfirmBill取得
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error: ObjectDecode")
		fmt.Println(err)
		return &pb.GetConfirmBillResponse{
			Count: 0,
			PayPrices: nil,
		}, nil
	}

	// 個数を取得
	var count int32 = int32(len(result.Comfirm))

	return &pb.GetConfirmBillResponse{
		Count: count,
		PayPrices: result.Comfirm,
	}, nil
}