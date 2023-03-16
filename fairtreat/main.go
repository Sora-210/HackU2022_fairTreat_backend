package main

import (
	"fmt"
	"log"
	"flag"
	"net"
	"context"

	"google.golang.org/grpc"
	pb "fairtreat.suwageeks.org/fairtreat/pb"
	gr "fairtreat.suwageeks.org/fairtreat/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// commandLine variable
var (
	port = flag.Int("port", 50000, "The gRPC Server Port")
)

// db URI
var (
	URI = "mongodb://root:root@db-primary:27017"
)
// main
func main() {
	// サーバーの構造体を生成
	server := gr.Server{}

	// 構造体ないにデータベースのクライアントを登録
	server.DB, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	// 切断する処理をdeferとして登録
	defer func() {
		if err := server.DB.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	} ()

	if err := server.DB.Ping(context.TODO(), readpref.Primary()); err != nil {
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
