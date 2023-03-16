package grpc

import (
	"go.mongodb.org/mongo-driver/mongo"
	
	pb "fairtreat.suwageeks.org/fairtreat/pb"
)

// wrap
type Server struct {
	pb.UnimplementedFairTreatServer
	DB *mongo.Client
}
