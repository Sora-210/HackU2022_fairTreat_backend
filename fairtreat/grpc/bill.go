package grpc

import (
	"context"
	"log"
	"math/rand"
	"regexp"
	"strconv"

	pb "fairtreat.suwageeks.org/fairtreat/pb"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fairtreat.suwageeks.org/fairtreat/model"
)


func (s *Server) CreateBill(ctx context.Context, req *pb.CreateBillRequest) (*pb.CreateBillResponse, error) {
	// 入力チェック
	if req.HostName == "" {
		return &pb.CreateBillResponse{
			BillId: "",
			Host: nil,
		}, nil
	}

	// コレクションを取得
	coll := s.DB.Database("fairtreat").Collection("Bill")

	// ホスト情報を生成
	hostName := req.HostName
	var hostId int32 = rand.Int31()

	// アイテム情報の整形
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

	// 明細初期データの生成
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

	// 生成データを挿入
	_, err := coll.InsertOne(context.Background(), bill)
	if err != nil {
		log.Printf("[CreateBill] Failed to insert initial data.\n%s\n", err)
	}

	// ホストをゲストリストに追加
	filter := bson.D{{
		Key: "_id",
		Value: bill.ID,
	}, {
		Key: "status",
		Value: true,
	}}
	_, err = coll.UpdateOne(context.TODO(), filter, bson.D{{
		Key: "$push",
		Value: bson.D{{
			Key: "Guests",
			Value: bill.Host,
		}},
	}})

	// 明細生成完了, レスポンス
	log.Printf("[CreateBill] Create Bill '%s'.\n", bill.ID)
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
		log.Printf("[GetBill] Error: id to ObjId.\n%s\n", err)
		return &pb.GetBillResponse{
			Bill: nil,
		}, nil
	}
	//明細が未確定のもの
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
		log.Printf("[GetBill] Error: ObjectDecode.\n%s\n", err)
		return &pb.GetBillResponse{
			Bill: nil,
		}, nil
	}

	// データを返す
	return &pb.GetBillResponse{
		Bill: &result,
	}, nil
}


func (s *Server) ConfirmBill(ctx context.Context, req *pb.ConfirmBillRequest) (*pb.ConfirmBillResponse, error) {
	// 入力チェック
	if req.Id == "" {
		return &pb.ConfirmBillResponse{
			Status: false,
		}, nil
	}
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")

	// Billを取得
	var result pb.Bill
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("[ConfirmBill] Error: id to ObjId\n%sn", err)
	}
	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}, {
		Key: "status",
		Value: true,
	}}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("[ConfirmBill] Error: ObjectDecode.\n%s\n", err)
		return &pb.ConfirmBillResponse{
			Status: false,
		}, nil
	}

	// 処理用配列を作成
	confirmPrices := map[int32]*model.PayPrice{}
	for _, v := range result.Guests {
		confirmPrices[v.Id] = &model.PayPrice{
			User: model.User{
				Id: v.Id,
				Name: v.Name,
			},
			Price: 0,
		}
	}

	// 金額計算処理
	for _, v := range result.Items {

		var count int32 = int32(len(v.Owners))
		// 除算対策
		if count == 0 {
			v.Owners = append(v.Owners, result.Host)
			count = 1
		}

		var price int32 = v.Price
		var priceOfPeople int32 = price / count
		var priceOfMod int32 = price - (priceOfPeople * count)
		flag := true
		for _, w := range v.Owners {
			if _, ok := confirmPrices[w.Id]; !ok {
				return &pb.ConfirmBillResponse{
					Status: false,
				}, nil
			}
			if flag {
				confirmPrices[w.Id].Price += priceOfMod
				flag = false
			}
			confirmPrices[w.Id].Price += priceOfPeople
		}
	}

	// 確定データ変換
	var confirmBill []*pb.PayPrice
	for _, v := range confirmPrices {
		confirmBill = append(confirmBill, &pb.PayPrice{
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
			Key: "confirm",
			Value: confirmBill,
		}},
	}, {
		Key: "$set",
		Value: bson.D{{
			Key: "status",
			Value: false,
		}},
	}})
	if err != nil {
		log.Printf("[ConfirmBill] Error: Update Document.\n%s\n", err)
	}

	// データを送信
	return &pb.ConfirmBillResponse{
		Status: (err == nil),
	}, nil
}


func (s *Server) GetConfirmBill(ctx context.Context, req *pb.GetConfirmBillRequest) (*pb.GetConfirmBillResponse, error) {
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")

	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("[GetConfirmBill] Error: ObjectDecode.\n%s\n", err)
	}

	filter := bson.D{{
		Key: "_id",
		Value: objId,
	}, {
		Key: "status",
		Value: false,
	}}

	// confirmBill取得
	var result model.ConfirmBill
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("[GetConfirmBill] Error: ObjectDecode.\n%s\n", err)
		return &pb.GetConfirmBillResponse{
			Count: 0,
			PayPrices: nil,
		}, nil
	}

	// 個数を取得
	var count int32 = int32(len(result.Confirm))

	return &pb.GetConfirmBillResponse{
		Count: count,
		PayPrices: result.Confirm,
	}, nil
}


func (s *Server) ConnectBill(req *pb.ConnectBillRequest, stream pb.FairTreat_ConnectBillServer) error {
	// コレクション取得
	coll := s.DB.Database("fairtreat").Collection("Bill")

	// Idを取得
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Printf("[ConnectBill] Error: id to ObjId\n%sn", err)
	}

	// streamを作成
	changeStream, err := coll.Watch(context.TODO(), mongo.Pipeline{bson.D{{
		Key: "$match",
		Value: bson.D{{
			Key: "documentKey",
			Value: bson.D{{
				Key: "_id",
				Value: objId,
			}},
		}},
	}}})
	if err != nil {
		panic(err)
	}
	defer changeStream.Close(context.TODO())

	// 変更を検知
	for changeStream.Next(context.TODO()) {
		var result model.ChangeStream
		var resType pb.BILL_CHANGE_TYPE
		var id int32

		changeStream.Decode(&result)
		for k, _ := range result.UpdateDescription.UpdatedFields {
			// 正規表現の生成
			guestRe := regexp.MustCompile(`^(Guest)`)
			ownersRe := regexp.MustCompile(`items.\d.owners`)
			comfirmRe := regexp.MustCompile(`^(comfirm)`)
			
			// AddUser
			if (guestRe.MatchString(k)) {
				resType = pb.BILL_CHANGE_TYPE_GUEST
				id = 0
			}
			// SetOwners
			if (ownersRe.MatchString(k)) {
				resType = pb.BILL_CHANGE_TYPE_ITEM

				itemId := k[6:]
				findIdRe := regexp.MustCompile(`.owners`)
				end := int(findIdRe.FindStringIndex(itemId)[0])
				itemId = itemId[:end]
				i, _ := strconv.Atoi(itemId)
				id = int32(i)
			}
			// ComfirmBill
			if (comfirmRe.MatchString(k)) {
				resType = pb.BILL_CHANGE_TYPE_CONFIRM
				id = 0
			}
		}

		res := &pb.ConnectBillResponse{
			Type: resType,
			Id: id,
		}
		stream.Send(res)
	}
	return nil
}
