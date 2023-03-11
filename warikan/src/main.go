package main

import (
	"time"
	"fmt"
	"log"
	"flag"
	"net"
	"context"
	"google.golang.org/grpc"
	pb "warikan.suwageeks.org/warikan/pkg/pb"
)

// commandLine variable
var (
	port = flag.Int("port", 50000, "The gRPC Server Port")
)

// wrap
type server struct {
	pb.UnimplementedWarikanServer
}

// grpc function
//// Bill
func (s *server) CreateBill(ctx context.Context, req *pb.CreateBillRequest) (*pb.CreateBillResponse, error) {
	return &pb.CreateBillResponse{
		Bill: &pb.Bill{
			Id: "3c3c0c3b-6430-49b3-59b2-9a1138e75f03",
			Password: "Vr9TBcAw",
			Host: &pb.User{
				Id: 123456789,
				Name: "TestUser",
			},
			Guests: nil,
		},
		OwnerPassword: "Az2Ka478",
	}, nil
}
func (s *server) ConnectBill (req *pb.ConnectBillRequest, stream pb.Warikan_ConnectBillServer) error {
	var id int32 = 123456789
	for i := 0; i < 5; i++ {
		res := &pb.ConnectBillResponse{
			Type: pb.BILL_CHANGE_TYPE_BILL,
			Id: id,
		}

		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
		id += 1
	}


	return nil
}
func (s *server) ConfirmBill (ctx context.Context, req *pb.ConfirmBillRequest) (*pb.ConfirmBillResponse, error) {
	return &pb.ConfirmBillResponse{
		Status: true,
	}, nil
}

//// User
func (s *server) AddUser (ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	return &pb.AddUserResponse{
		Status: true,
	}, nil
}
func (s *server) RemoveUser (ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	return &pb.RemoveUserResponse{
		Status: true,
	}, nil
}
func (s *server) GetUsersList (ctx context.Context, req *pb.GetUsersListRequest) (*pb.GetUsersListResponse, error) {
	users := [] *pb.User{
		{
			Id: 123456789,
			Name: "TestUser",
		},
		{
			Id: 987654321,
			Name: "TestUser2",
		},
	}

	return &pb.GetUsersListResponse{
		Count: 2,
		Users: users,
	}, nil
}

//// Item
func (s *server) SetItems (ctx context.Context, req *pb.SetItemsRequest) (*pb.SetItemsResponse, error) {
	return &pb.SetItemsResponse{
		Status: true,
	}, nil
}
func (s *server) AddItem (ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	return &pb.AddItemResponse{
		Status: true,
	}, nil
}
func (s *server) RemoveItem (ctx context.Context, req *pb.RemoveItemRequest) (*pb.RemoveItemResponse, error) {
	return &pb.RemoveItemResponse{
		Status: true,
	}, nil
}
func (s *server) GetItemsList (ctx context.Context, req *pb.GetItemsListRequest) (*pb.GetItemsListResponse, error) {
	items := [] *pb.Item{
		{
			Id: 123456789,
			Name: "item1",
			Amount: 1,
			Price: 2000,
			Owners: [] *pb.User{
				{
					Id: 123456789,
					Name: "TestUser",
				},
			},
		},
		{
			Id: 987654321,
			Name: "item2",
			Amount: 3,
			Price: 4500,
			Owners: [] *pb.User{
				{
					Id: 123456789,
					Name: "TestUser",
				},
				{
					Id: 987654321,
					Name: "TestUser2",
				},
			},
		},
	}

	return &pb.GetItemsListResponse{
		Count: 2,
		Items: items,
	}, nil
}
func (s *server) GetItem (ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	return &pb.GetItemResponse{
		Item: &pb.Item{
			Id: 123456789,
			Name: "item1",
			Amount: 1,
			Price: 2000,
			Owners: [] *pb.User{
				{
					Id: 123456789,
					Name: "TestUser",
				},
			},
		},
	}, nil
}


//// Owner
func (s *server) SetItemOwners (ctx context.Context, req *pb.SetItemOwnersRequest) (*pb.SetItemOwnersResponse, error) {
	return &pb.SetItemOwnersResponse{
		Status: true,
	}, nil
}
func (s *server) AddItemOwner (ctx context.Context, req *pb.AddItemOwnerRequest) (*pb.AddItemOwnerResponse, error) {
	return &pb.AddItemOwnerResponse{
		Status: true,
	}, nil
}
func (s *server) RemoveItemOwner (ctx context.Context, req *pb.RemoveItemOwnerRequest) (*pb.RemoveItemOwnerResponse, error) {
	return &pb.RemoveItemOwnerResponse{
		Status: true,
	}, nil
}
func (s *server) GetItemOwners (ctx context.Context, req *pb.GetItemOwnersRequest) (*pb.GetItemOwnersResponse, error) {
	users := [] *pb.User{
		{
			Id: 123456789,
			Name: "TestUser",
		},
		{
			Id: 987654321,
			Name: "TestUser2",
		},
	}

	return &pb.GetItemOwnersResponse{
		Count: 2,
		Owners: users,
	}, nil
}


// main
func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to Listen:\n\t %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWarikanServer(s, &server{})
	log.Printf("Start Server Listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve:\n\t%v", err)
	}
}
