package main

import (
	"fmt"
	"log"
	"flag"
	"net"
	"context"

	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	
	pb "fairtreat.suwageeks.org/fairtreat/pb"
)

// commandLine variable
var (
	port = flag.Int("port", 50000, "The gRPC Server Port")
)

// wrap
type server struct {
	pb.UnimplementedFairTreatServer
	db *mongo.Client
}

// db URI
var (
	URI = "mongodb://root:root@db:27017"
)
// main
func main() {
	// サーバーの構造体を生成
	server := server{}

	// 構造体ないにデータベースのクライアントを登録
	server.db, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	// 切断する処理をdeferとして登録
	defer func() {
		if err := server.db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	} ()

	if err := server.db.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("Successfully connected and pinged.")

	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to Listen:\n\t %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFairTreatServer(s, &server)
	log.Printf("Start Server Listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve:\n\t%v", err)
	}
}
