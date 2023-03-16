// Config

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: fairtreat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FairTreat_CreateBill_FullMethodName        = "/fairtreat.FairTreat/CreateBill"
	FairTreat_GetBill_FullMethodName           = "/fairtreat.FairTreat/GetBill"
	FairTreat_ConnectBill_FullMethodName       = "/fairtreat.FairTreat/ConnectBill"
	FairTreat_ConfirmBill_FullMethodName       = "/fairtreat.FairTreat/ConfirmBill"
	FairTreat_GetConfirmBill_FullMethodName    = "/fairtreat.FairTreat/GetConfirmBill"
	FairTreat_AddUser_FullMethodName           = "/fairtreat.FairTreat/AddUser"
	FairTreat_GetUsersList_FullMethodName      = "/fairtreat.FairTreat/GetUsersList"
	FairTreat_GetItemsList_FullMethodName      = "/fairtreat.FairTreat/GetItemsList"
	FairTreat_SetOwners_FullMethodName         = "/fairtreat.FairTreat/SetOwners"
	FairTreat_GetItemOwnersList_FullMethodName = "/fairtreat.FairTreat/GetItemOwnersList"
)

// FairTreatClient is the client API for FairTreat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FairTreatClient interface {
	CreateBill(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*CreateBillResponse, error)
	GetBill(ctx context.Context, in *GetBillRequest, opts ...grpc.CallOption) (*GetBillResponse, error)
	ConnectBill(ctx context.Context, in *ConnectBillRequest, opts ...grpc.CallOption) (FairTreat_ConnectBillClient, error)
	ConfirmBill(ctx context.Context, in *ConfirmBillRequest, opts ...grpc.CallOption) (*ConfirmBillResponse, error)
	GetConfirmBill(ctx context.Context, in *GetConfirmBillRequest, opts ...grpc.CallOption) (*GetConfirmBillResponse, error)
	// ############
	// Methods for each room
	// ############
	// Users
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	GetUsersList(ctx context.Context, in *GetUsersListRequest, opts ...grpc.CallOption) (*GetUsersListResponse, error)
	// Items
	GetItemsList(ctx context.Context, in *GetItemsListRequest, opts ...grpc.CallOption) (*GetItemsListResponse, error)
	// ItemOwners
	// // 特定商品の支払う人のみを取得したい時に使う
	SetOwners(ctx context.Context, in *SetItemOwnerRequest, opts ...grpc.CallOption) (*SetItemOwnerResponse, error)
	GetItemOwnersList(ctx context.Context, in *GetItemOwnersRequest, opts ...grpc.CallOption) (*GetItemOwnersResponse, error)
}

type fairTreatClient struct {
	cc grpc.ClientConnInterface
}

func NewFairTreatClient(cc grpc.ClientConnInterface) FairTreatClient {
	return &fairTreatClient{cc}
}

func (c *fairTreatClient) CreateBill(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*CreateBillResponse, error) {
	out := new(CreateBillResponse)
	err := c.cc.Invoke(ctx, FairTreat_CreateBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) GetBill(ctx context.Context, in *GetBillRequest, opts ...grpc.CallOption) (*GetBillResponse, error) {
	out := new(GetBillResponse)
	err := c.cc.Invoke(ctx, FairTreat_GetBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) ConnectBill(ctx context.Context, in *ConnectBillRequest, opts ...grpc.CallOption) (FairTreat_ConnectBillClient, error) {
	stream, err := c.cc.NewStream(ctx, &FairTreat_ServiceDesc.Streams[0], FairTreat_ConnectBill_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fairTreatConnectBillClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FairTreat_ConnectBillClient interface {
	Recv() (*ConnectBillResponse, error)
	grpc.ClientStream
}

type fairTreatConnectBillClient struct {
	grpc.ClientStream
}

func (x *fairTreatConnectBillClient) Recv() (*ConnectBillResponse, error) {
	m := new(ConnectBillResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fairTreatClient) ConfirmBill(ctx context.Context, in *ConfirmBillRequest, opts ...grpc.CallOption) (*ConfirmBillResponse, error) {
	out := new(ConfirmBillResponse)
	err := c.cc.Invoke(ctx, FairTreat_ConfirmBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) GetConfirmBill(ctx context.Context, in *GetConfirmBillRequest, opts ...grpc.CallOption) (*GetConfirmBillResponse, error) {
	out := new(GetConfirmBillResponse)
	err := c.cc.Invoke(ctx, FairTreat_GetConfirmBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, FairTreat_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) GetUsersList(ctx context.Context, in *GetUsersListRequest, opts ...grpc.CallOption) (*GetUsersListResponse, error) {
	out := new(GetUsersListResponse)
	err := c.cc.Invoke(ctx, FairTreat_GetUsersList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) GetItemsList(ctx context.Context, in *GetItemsListRequest, opts ...grpc.CallOption) (*GetItemsListResponse, error) {
	out := new(GetItemsListResponse)
	err := c.cc.Invoke(ctx, FairTreat_GetItemsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) SetOwners(ctx context.Context, in *SetItemOwnerRequest, opts ...grpc.CallOption) (*SetItemOwnerResponse, error) {
	out := new(SetItemOwnerResponse)
	err := c.cc.Invoke(ctx, FairTreat_SetOwners_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fairTreatClient) GetItemOwnersList(ctx context.Context, in *GetItemOwnersRequest, opts ...grpc.CallOption) (*GetItemOwnersResponse, error) {
	out := new(GetItemOwnersResponse)
	err := c.cc.Invoke(ctx, FairTreat_GetItemOwnersList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FairTreatServer is the server API for FairTreat service.
// All implementations must embed UnimplementedFairTreatServer
// for forward compatibility
type FairTreatServer interface {
	CreateBill(context.Context, *CreateBillRequest) (*CreateBillResponse, error)
	GetBill(context.Context, *GetBillRequest) (*GetBillResponse, error)
	ConnectBill(*ConnectBillRequest, FairTreat_ConnectBillServer) error
	ConfirmBill(context.Context, *ConfirmBillRequest) (*ConfirmBillResponse, error)
	GetConfirmBill(context.Context, *GetConfirmBillRequest) (*GetConfirmBillResponse, error)
	// ############
	// Methods for each room
	// ############
	// Users
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	GetUsersList(context.Context, *GetUsersListRequest) (*GetUsersListResponse, error)
	// Items
	GetItemsList(context.Context, *GetItemsListRequest) (*GetItemsListResponse, error)
	// ItemOwners
	// // 特定商品の支払う人のみを取得したい時に使う
	SetOwners(context.Context, *SetItemOwnerRequest) (*SetItemOwnerResponse, error)
	GetItemOwnersList(context.Context, *GetItemOwnersRequest) (*GetItemOwnersResponse, error)
	mustEmbedUnimplementedFairTreatServer()
}

// UnimplementedFairTreatServer must be embedded to have forward compatible implementations.
type UnimplementedFairTreatServer struct {
}

func (UnimplementedFairTreatServer) CreateBill(context.Context, *CreateBillRequest) (*CreateBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBill not implemented")
}
func (UnimplementedFairTreatServer) GetBill(context.Context, *GetBillRequest) (*GetBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBill not implemented")
}
func (UnimplementedFairTreatServer) ConnectBill(*ConnectBillRequest, FairTreat_ConnectBillServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectBill not implemented")
}
func (UnimplementedFairTreatServer) ConfirmBill(context.Context, *ConfirmBillRequest) (*ConfirmBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmBill not implemented")
}
func (UnimplementedFairTreatServer) GetConfirmBill(context.Context, *GetConfirmBillRequest) (*GetConfirmBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfirmBill not implemented")
}
func (UnimplementedFairTreatServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedFairTreatServer) GetUsersList(context.Context, *GetUsersListRequest) (*GetUsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersList not implemented")
}
func (UnimplementedFairTreatServer) GetItemsList(context.Context, *GetItemsListRequest) (*GetItemsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemsList not implemented")
}
func (UnimplementedFairTreatServer) SetOwners(context.Context, *SetItemOwnerRequest) (*SetItemOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOwners not implemented")
}
func (UnimplementedFairTreatServer) GetItemOwnersList(context.Context, *GetItemOwnersRequest) (*GetItemOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemOwnersList not implemented")
}
func (UnimplementedFairTreatServer) mustEmbedUnimplementedFairTreatServer() {}

// UnsafeFairTreatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FairTreatServer will
// result in compilation errors.
type UnsafeFairTreatServer interface {
	mustEmbedUnimplementedFairTreatServer()
}

func RegisterFairTreatServer(s grpc.ServiceRegistrar, srv FairTreatServer) {
	s.RegisterService(&FairTreat_ServiceDesc, srv)
}

func _FairTreat_CreateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).CreateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_CreateBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).CreateBill(ctx, req.(*CreateBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_GetBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).GetBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_GetBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).GetBill(ctx, req.(*GetBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_ConnectBill_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectBillRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FairTreatServer).ConnectBill(m, &fairTreatConnectBillServer{stream})
}

type FairTreat_ConnectBillServer interface {
	Send(*ConnectBillResponse) error
	grpc.ServerStream
}

type fairTreatConnectBillServer struct {
	grpc.ServerStream
}

func (x *fairTreatConnectBillServer) Send(m *ConnectBillResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FairTreat_ConfirmBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).ConfirmBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_ConfirmBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).ConfirmBill(ctx, req.(*ConfirmBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_GetConfirmBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfirmBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).GetConfirmBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_GetConfirmBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).GetConfirmBill(ctx, req.(*GetConfirmBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_GetUsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).GetUsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_GetUsersList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).GetUsersList(ctx, req.(*GetUsersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_GetItemsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).GetItemsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_GetItemsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).GetItemsList(ctx, req.(*GetItemsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_SetOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetItemOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).SetOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_SetOwners_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).SetOwners(ctx, req.(*SetItemOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FairTreat_GetItemOwnersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FairTreatServer).GetItemOwnersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FairTreat_GetItemOwnersList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FairTreatServer).GetItemOwnersList(ctx, req.(*GetItemOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FairTreat_ServiceDesc is the grpc.ServiceDesc for FairTreat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FairTreat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fairtreat.FairTreat",
	HandlerType: (*FairTreatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBill",
			Handler:    _FairTreat_CreateBill_Handler,
		},
		{
			MethodName: "GetBill",
			Handler:    _FairTreat_GetBill_Handler,
		},
		{
			MethodName: "ConfirmBill",
			Handler:    _FairTreat_ConfirmBill_Handler,
		},
		{
			MethodName: "GetConfirmBill",
			Handler:    _FairTreat_GetConfirmBill_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _FairTreat_AddUser_Handler,
		},
		{
			MethodName: "GetUsersList",
			Handler:    _FairTreat_GetUsersList_Handler,
		},
		{
			MethodName: "GetItemsList",
			Handler:    _FairTreat_GetItemsList_Handler,
		},
		{
			MethodName: "SetOwners",
			Handler:    _FairTreat_SetOwners_Handler,
		},
		{
			MethodName: "GetItemOwnersList",
			Handler:    _FairTreat_GetItemOwnersList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectBill",
			Handler:       _FairTreat_ConnectBill_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fairtreat.proto",
}