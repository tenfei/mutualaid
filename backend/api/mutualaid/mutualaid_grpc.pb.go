// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/mutualaid/mutualaid.proto

package mutualaid

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 公众号oauth
	WxOAuth2(ctx context.Context, in *WxOAuth2Req, opts ...grpc.CallOption) (*WxOAuth2Resp, error)
	// 小程序wx.Login
	WxLogin(ctx context.Context, in *WxLoginReq, opts ...grpc.CallOption) (*WxLoginResp, error)
	// 小程序激活用户, 是WxLogin的超集
	ActiveUser(ctx context.Context, in *ActiveUserReq, opts ...grpc.CallOption) (*ActiveUserResp, error)
	// 小程序取用户手机号
	WxPhoneNumber(ctx context.Context, in *WxPhoneNumberReq, opts ...grpc.CallOption) (*WxPhoneNumberResp, error)
	// 查看自己的信息
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	// 辅助前端进行参数签名
	JSAPISign(ctx context.Context, in *JSAPISignReq, opts ...grpc.CallOption) (*JSAPISignResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) WxOAuth2(ctx context.Context, in *WxOAuth2Req, opts ...grpc.CallOption) (*WxOAuth2Resp, error) {
	out := new(WxOAuth2Resp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserService/WxOAuth2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) WxLogin(ctx context.Context, in *WxLoginReq, opts ...grpc.CallOption) (*WxLoginResp, error) {
	out := new(WxLoginResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserService/WxLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActiveUser(ctx context.Context, in *ActiveUserReq, opts ...grpc.CallOption) (*ActiveUserResp, error) {
	out := new(ActiveUserResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserService/ActiveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) WxPhoneNumber(ctx context.Context, in *WxPhoneNumberReq, opts ...grpc.CallOption) (*WxPhoneNumberResp, error) {
	out := new(WxPhoneNumberResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserService/WxPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) JSAPISign(ctx context.Context, in *JSAPISignReq, opts ...grpc.CallOption) (*JSAPISignResp, error) {
	out := new(JSAPISignResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserService/JSAPISign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 公众号oauth
	WxOAuth2(context.Context, *WxOAuth2Req) (*WxOAuth2Resp, error)
	// 小程序wx.Login
	WxLogin(context.Context, *WxLoginReq) (*WxLoginResp, error)
	// 小程序激活用户, 是WxLogin的超集
	ActiveUser(context.Context, *ActiveUserReq) (*ActiveUserResp, error)
	// 小程序取用户手机号
	WxPhoneNumber(context.Context, *WxPhoneNumberReq) (*WxPhoneNumberResp, error)
	// 查看自己的信息
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	// 辅助前端进行参数签名
	JSAPISign(context.Context, *JSAPISignReq) (*JSAPISignResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) WxOAuth2(context.Context, *WxOAuth2Req) (*WxOAuth2Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxOAuth2 not implemented")
}
func (UnimplementedUserServiceServer) WxLogin(context.Context, *WxLoginReq) (*WxLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxLogin not implemented")
}
func (UnimplementedUserServiceServer) ActiveUser(context.Context, *ActiveUserReq) (*ActiveUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveUser not implemented")
}
func (UnimplementedUserServiceServer) WxPhoneNumber(context.Context, *WxPhoneNumberReq) (*WxPhoneNumberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPhoneNumber not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) JSAPISign(context.Context, *JSAPISignReq) (*JSAPISignResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JSAPISign not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_WxOAuth2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxOAuth2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).WxOAuth2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserService/WxOAuth2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).WxOAuth2(ctx, req.(*WxOAuth2Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_WxLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).WxLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserService/WxLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).WxLogin(ctx, req.(*WxLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActiveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActiveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserService/ActiveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActiveUser(ctx, req.(*ActiveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_WxPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxPhoneNumberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).WxPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserService/WxPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).WxPhoneNumber(ctx, req.(*WxPhoneNumberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_JSAPISign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JSAPISignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).JSAPISign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserService/JSAPISign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).JSAPISign(ctx, req.(*JSAPISignReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mutualaid.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WxOAuth2",
			Handler:    _UserService_WxOAuth2_Handler,
		},
		{
			MethodName: "WxLogin",
			Handler:    _UserService_WxLogin_Handler,
		},
		{
			MethodName: "ActiveUser",
			Handler:    _UserService_ActiveUser_Handler,
		},
		{
			MethodName: "WxPhoneNumber",
			Handler:    _UserService_WxPhoneNumber_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "JSAPISign",
			Handler:    _UserService_JSAPISign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mutualaid/mutualaid.proto",
}

// MutualAidQueryClient is the client API for MutualAidQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutualAidQueryClient interface {
	//*
	//  发现接口，需要过滤，只展示等待审核的
	//  发现接口
	Discovery(ctx context.Context, in *DiscoveryReq, opts ...grpc.CallOption) (*DiscoveryResp, error)
}

type mutualAidQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewMutualAidQueryClient(cc grpc.ClientConnInterface) MutualAidQueryClient {
	return &mutualAidQueryClient{cc}
}

func (c *mutualAidQueryClient) Discovery(ctx context.Context, in *DiscoveryReq, opts ...grpc.CallOption) (*DiscoveryResp, error) {
	out := new(DiscoveryResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.MutualAidQuery/Discovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutualAidQueryServer is the server API for MutualAidQuery service.
// All implementations must embed UnimplementedMutualAidQueryServer
// for forward compatibility
type MutualAidQueryServer interface {
	//*
	//  发现接口，需要过滤，只展示等待审核的
	//  发现接口
	Discovery(context.Context, *DiscoveryReq) (*DiscoveryResp, error)
	mustEmbedUnimplementedMutualAidQueryServer()
}

// UnimplementedMutualAidQueryServer must be embedded to have forward compatible implementations.
type UnimplementedMutualAidQueryServer struct {
}

func (UnimplementedMutualAidQueryServer) Discovery(context.Context, *DiscoveryReq) (*DiscoveryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discovery not implemented")
}
func (UnimplementedMutualAidQueryServer) mustEmbedUnimplementedMutualAidQueryServer() {}

// UnsafeMutualAidQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutualAidQueryServer will
// result in compilation errors.
type UnsafeMutualAidQueryServer interface {
	mustEmbedUnimplementedMutualAidQueryServer()
}

func RegisterMutualAidQueryServer(s grpc.ServiceRegistrar, srv MutualAidQueryServer) {
	s.RegisterService(&MutualAidQuery_ServiceDesc, srv)
}

func _MutualAidQuery_Discovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualAidQueryServer).Discovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.MutualAidQuery/Discovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualAidQueryServer).Discovery(ctx, req.(*DiscoveryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MutualAidQuery_ServiceDesc is the grpc.ServiceDesc for MutualAidQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MutualAidQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mutualaid.MutualAidQuery",
	HandlerType: (*MutualAidQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Discovery",
			Handler:    _MutualAidQuery_Discovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mutualaid/mutualaid.proto",
}

// ExamineAidClient is the client API for ExamineAid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExamineAidClient interface {
	//*
	//  审核用户登录
	ExamineLogin(ctx context.Context, in *ExamineLoginReq, opts ...grpc.CallOption) (*ExamineLoginResp, error)
	//*
	//  获取求助列表
	GetExamineList(ctx context.Context, in *GetExamineListReq, opts ...grpc.CallOption) (*ExamineListResp, error)
	//*
	//  审核消息
	ExamineAid(ctx context.Context, in *ExamineAidReq, opts ...grpc.CallOption) (*ExamineAidResp, error)
	// 拉黑用户（包含屏蔽基消息）
	BlockUser(ctx context.Context, in *BlockUserReq, opts ...grpc.CallOption) (*BlockUserResp, error)
	// 获取屏蔽用户列表
	GetBlockUserList(ctx context.Context, in *GetBlockUserListReq, opts ...grpc.CallOption) (*GetBlockUserListResp, error)
	// 恢复用户（包含恢复其消息）
	PassUser(ctx context.Context, in *PassUserReq, opts ...grpc.CallOption) (*PassUserResp, error)
}

type examineAidClient struct {
	cc grpc.ClientConnInterface
}

func NewExamineAidClient(cc grpc.ClientConnInterface) ExamineAidClient {
	return &examineAidClient{cc}
}

func (c *examineAidClient) ExamineLogin(ctx context.Context, in *ExamineLoginReq, opts ...grpc.CallOption) (*ExamineLoginResp, error) {
	out := new(ExamineLoginResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.ExamineAid/ExamineLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examineAidClient) GetExamineList(ctx context.Context, in *GetExamineListReq, opts ...grpc.CallOption) (*ExamineListResp, error) {
	out := new(ExamineListResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.ExamineAid/GetExamineList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examineAidClient) ExamineAid(ctx context.Context, in *ExamineAidReq, opts ...grpc.CallOption) (*ExamineAidResp, error) {
	out := new(ExamineAidResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.ExamineAid/ExamineAid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examineAidClient) BlockUser(ctx context.Context, in *BlockUserReq, opts ...grpc.CallOption) (*BlockUserResp, error) {
	out := new(BlockUserResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.ExamineAid/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examineAidClient) GetBlockUserList(ctx context.Context, in *GetBlockUserListReq, opts ...grpc.CallOption) (*GetBlockUserListResp, error) {
	out := new(GetBlockUserListResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.ExamineAid/GetBlockUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examineAidClient) PassUser(ctx context.Context, in *PassUserReq, opts ...grpc.CallOption) (*PassUserResp, error) {
	out := new(PassUserResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.ExamineAid/PassUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExamineAidServer is the server API for ExamineAid service.
// All implementations must embed UnimplementedExamineAidServer
// for forward compatibility
type ExamineAidServer interface {
	//*
	//  审核用户登录
	ExamineLogin(context.Context, *ExamineLoginReq) (*ExamineLoginResp, error)
	//*
	//  获取求助列表
	GetExamineList(context.Context, *GetExamineListReq) (*ExamineListResp, error)
	//*
	//  审核消息
	ExamineAid(context.Context, *ExamineAidReq) (*ExamineAidResp, error)
	// 拉黑用户（包含屏蔽基消息）
	BlockUser(context.Context, *BlockUserReq) (*BlockUserResp, error)
	// 获取屏蔽用户列表
	GetBlockUserList(context.Context, *GetBlockUserListReq) (*GetBlockUserListResp, error)
	// 恢复用户（包含恢复其消息）
	PassUser(context.Context, *PassUserReq) (*PassUserResp, error)
	mustEmbedUnimplementedExamineAidServer()
}

// UnimplementedExamineAidServer must be embedded to have forward compatible implementations.
type UnimplementedExamineAidServer struct {
}

func (UnimplementedExamineAidServer) ExamineLogin(context.Context, *ExamineLoginReq) (*ExamineLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExamineLogin not implemented")
}
func (UnimplementedExamineAidServer) GetExamineList(context.Context, *GetExamineListReq) (*ExamineListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExamineList not implemented")
}
func (UnimplementedExamineAidServer) ExamineAid(context.Context, *ExamineAidReq) (*ExamineAidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExamineAid not implemented")
}
func (UnimplementedExamineAidServer) BlockUser(context.Context, *BlockUserReq) (*BlockUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedExamineAidServer) GetBlockUserList(context.Context, *GetBlockUserListReq) (*GetBlockUserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockUserList not implemented")
}
func (UnimplementedExamineAidServer) PassUser(context.Context, *PassUserReq) (*PassUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassUser not implemented")
}
func (UnimplementedExamineAidServer) mustEmbedUnimplementedExamineAidServer() {}

// UnsafeExamineAidServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExamineAidServer will
// result in compilation errors.
type UnsafeExamineAidServer interface {
	mustEmbedUnimplementedExamineAidServer()
}

func RegisterExamineAidServer(s grpc.ServiceRegistrar, srv ExamineAidServer) {
	s.RegisterService(&ExamineAid_ServiceDesc, srv)
}

func _ExamineAid_ExamineLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExamineLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamineAidServer).ExamineLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.ExamineAid/ExamineLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamineAidServer).ExamineLogin(ctx, req.(*ExamineLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamineAid_GetExamineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExamineListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamineAidServer).GetExamineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.ExamineAid/GetExamineList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamineAidServer).GetExamineList(ctx, req.(*GetExamineListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamineAid_ExamineAid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExamineAidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamineAidServer).ExamineAid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.ExamineAid/ExamineAid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamineAidServer).ExamineAid(ctx, req.(*ExamineAidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamineAid_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamineAidServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.ExamineAid/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamineAidServer).BlockUser(ctx, req.(*BlockUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamineAid_GetBlockUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamineAidServer).GetBlockUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.ExamineAid/GetBlockUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamineAidServer).GetBlockUserList(ctx, req.(*GetBlockUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamineAid_PassUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamineAidServer).PassUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.ExamineAid/PassUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamineAidServer).PassUser(ctx, req.(*PassUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ExamineAid_ServiceDesc is the grpc.ServiceDesc for ExamineAid service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExamineAid_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mutualaid.ExamineAid",
	HandlerType: (*ExamineAidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExamineLogin",
			Handler:    _ExamineAid_ExamineLogin_Handler,
		},
		{
			MethodName: "GetExamineList",
			Handler:    _ExamineAid_GetExamineList_Handler,
		},
		{
			MethodName: "ExamineAid",
			Handler:    _ExamineAid_ExamineAid_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _ExamineAid_BlockUser_Handler,
		},
		{
			MethodName: "GetBlockUserList",
			Handler:    _ExamineAid_GetBlockUserList_Handler,
		},
		{
			MethodName: "PassUser",
			Handler:    _ExamineAid_PassUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mutualaid/mutualaid.proto",
}

// UserAidQueryClient is the client API for UserAidQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAidQueryClient interface {
	//*
	//  我的帮助查询接口,这里可以展示自己全部的消息
	ListAidOffered(ctx context.Context, in *ListAidOfferedReq, opts ...grpc.CallOption) (*ListAidOfferedResp, error)
	//*
	//  我的求助查询接口,这里可以展示自己全部的消息
	//  我的求助查询接
	ListAidNeeds(ctx context.Context, in *ListAidNeedsReq, opts ...grpc.CallOption) (*ListAidNeedsResp, error)
	//*
	//  求助详情查询接口
	GetAidDetail(ctx context.Context, in *GetAidDetailReq, opts ...grpc.CallOption) (*GetAidDetailResp, error)
}

type userAidQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAidQueryClient(cc grpc.ClientConnInterface) UserAidQueryClient {
	return &userAidQueryClient{cc}
}

func (c *userAidQueryClient) ListAidOffered(ctx context.Context, in *ListAidOfferedReq, opts ...grpc.CallOption) (*ListAidOfferedResp, error) {
	out := new(ListAidOfferedResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidQuery/ListAidOffered", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAidQueryClient) ListAidNeeds(ctx context.Context, in *ListAidNeedsReq, opts ...grpc.CallOption) (*ListAidNeedsResp, error) {
	out := new(ListAidNeedsResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidQuery/ListAidNeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAidQueryClient) GetAidDetail(ctx context.Context, in *GetAidDetailReq, opts ...grpc.CallOption) (*GetAidDetailResp, error) {
	out := new(GetAidDetailResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidQuery/GetAidDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAidQueryServer is the server API for UserAidQuery service.
// All implementations must embed UnimplementedUserAidQueryServer
// for forward compatibility
type UserAidQueryServer interface {
	//*
	//  我的帮助查询接口,这里可以展示自己全部的消息
	ListAidOffered(context.Context, *ListAidOfferedReq) (*ListAidOfferedResp, error)
	//*
	//  我的求助查询接口,这里可以展示自己全部的消息
	//  我的求助查询接
	ListAidNeeds(context.Context, *ListAidNeedsReq) (*ListAidNeedsResp, error)
	//*
	//  求助详情查询接口
	GetAidDetail(context.Context, *GetAidDetailReq) (*GetAidDetailResp, error)
	mustEmbedUnimplementedUserAidQueryServer()
}

// UnimplementedUserAidQueryServer must be embedded to have forward compatible implementations.
type UnimplementedUserAidQueryServer struct {
}

func (UnimplementedUserAidQueryServer) ListAidOffered(context.Context, *ListAidOfferedReq) (*ListAidOfferedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAidOffered not implemented")
}
func (UnimplementedUserAidQueryServer) ListAidNeeds(context.Context, *ListAidNeedsReq) (*ListAidNeedsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAidNeeds not implemented")
}
func (UnimplementedUserAidQueryServer) GetAidDetail(context.Context, *GetAidDetailReq) (*GetAidDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAidDetail not implemented")
}
func (UnimplementedUserAidQueryServer) mustEmbedUnimplementedUserAidQueryServer() {}

// UnsafeUserAidQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAidQueryServer will
// result in compilation errors.
type UnsafeUserAidQueryServer interface {
	mustEmbedUnimplementedUserAidQueryServer()
}

func RegisterUserAidQueryServer(s grpc.ServiceRegistrar, srv UserAidQueryServer) {
	s.RegisterService(&UserAidQuery_ServiceDesc, srv)
}

func _UserAidQuery_ListAidOffered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAidOfferedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidQueryServer).ListAidOffered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidQuery/ListAidOffered",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidQueryServer).ListAidOffered(ctx, req.(*ListAidOfferedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAidQuery_ListAidNeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAidNeedsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidQueryServer).ListAidNeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidQuery/ListAidNeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidQueryServer).ListAidNeeds(ctx, req.(*ListAidNeedsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAidQuery_GetAidDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAidDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidQueryServer).GetAidDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidQuery/GetAidDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidQueryServer).GetAidDetail(ctx, req.(*GetAidDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAidQuery_ServiceDesc is the grpc.ServiceDesc for UserAidQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAidQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mutualaid.UserAidQuery",
	HandlerType: (*UserAidQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAidOffered",
			Handler:    _UserAidQuery_ListAidOffered_Handler,
		},
		{
			MethodName: "ListAidNeeds",
			Handler:    _UserAidQuery_ListAidNeeds_Handler,
		},
		{
			MethodName: "GetAidDetail",
			Handler:    _UserAidQuery_GetAidDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mutualaid/mutualaid.proto",
}

// UserAidManagerClient is the client API for UserAidManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAidManagerClient interface {
	//*
	//  新建求助
	CreateAid(ctx context.Context, in *CreateAidReq, opts ...grpc.CallOption) (*CreateAidResp, error)
	//*
	//  我的求助查询接口
	CancelAid(ctx context.Context, in *CancelAidReq, opts ...grpc.CallOption) (*CancelAidResp, error)
	//*
	//  完成救助接口
	FinishAid(ctx context.Context, in *FinishAidReq, opts ...grpc.CallOption) (*FinishAidResp, error)
	//*
	//  留言接口
	CreateAidMessage(ctx context.Context, in *CreateAidMessageReq, opts ...grpc.CallOption) (*CreateAidMessageResp, error)
}

type userAidManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAidManagerClient(cc grpc.ClientConnInterface) UserAidManagerClient {
	return &userAidManagerClient{cc}
}

func (c *userAidManagerClient) CreateAid(ctx context.Context, in *CreateAidReq, opts ...grpc.CallOption) (*CreateAidResp, error) {
	out := new(CreateAidResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidManager/CreateAid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAidManagerClient) CancelAid(ctx context.Context, in *CancelAidReq, opts ...grpc.CallOption) (*CancelAidResp, error) {
	out := new(CancelAidResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidManager/CancelAid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAidManagerClient) FinishAid(ctx context.Context, in *FinishAidReq, opts ...grpc.CallOption) (*FinishAidResp, error) {
	out := new(FinishAidResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidManager/FinishAid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAidManagerClient) CreateAidMessage(ctx context.Context, in *CreateAidMessageReq, opts ...grpc.CallOption) (*CreateAidMessageResp, error) {
	out := new(CreateAidMessageResp)
	err := c.cc.Invoke(ctx, "/api.mutualaid.UserAidManager/CreateAidMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAidManagerServer is the server API for UserAidManager service.
// All implementations must embed UnimplementedUserAidManagerServer
// for forward compatibility
type UserAidManagerServer interface {
	//*
	//  新建求助
	CreateAid(context.Context, *CreateAidReq) (*CreateAidResp, error)
	//*
	//  我的求助查询接口
	CancelAid(context.Context, *CancelAidReq) (*CancelAidResp, error)
	//*
	//  完成救助接口
	FinishAid(context.Context, *FinishAidReq) (*FinishAidResp, error)
	//*
	//  留言接口
	CreateAidMessage(context.Context, *CreateAidMessageReq) (*CreateAidMessageResp, error)
	mustEmbedUnimplementedUserAidManagerServer()
}

// UnimplementedUserAidManagerServer must be embedded to have forward compatible implementations.
type UnimplementedUserAidManagerServer struct {
}

func (UnimplementedUserAidManagerServer) CreateAid(context.Context, *CreateAidReq) (*CreateAidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAid not implemented")
}
func (UnimplementedUserAidManagerServer) CancelAid(context.Context, *CancelAidReq) (*CancelAidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAid not implemented")
}
func (UnimplementedUserAidManagerServer) FinishAid(context.Context, *FinishAidReq) (*FinishAidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishAid not implemented")
}
func (UnimplementedUserAidManagerServer) CreateAidMessage(context.Context, *CreateAidMessageReq) (*CreateAidMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAidMessage not implemented")
}
func (UnimplementedUserAidManagerServer) mustEmbedUnimplementedUserAidManagerServer() {}

// UnsafeUserAidManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAidManagerServer will
// result in compilation errors.
type UnsafeUserAidManagerServer interface {
	mustEmbedUnimplementedUserAidManagerServer()
}

func RegisterUserAidManagerServer(s grpc.ServiceRegistrar, srv UserAidManagerServer) {
	s.RegisterService(&UserAidManager_ServiceDesc, srv)
}

func _UserAidManager_CreateAid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidManagerServer).CreateAid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidManager/CreateAid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidManagerServer).CreateAid(ctx, req.(*CreateAidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAidManager_CancelAid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidManagerServer).CancelAid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidManager/CancelAid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidManagerServer).CancelAid(ctx, req.(*CancelAidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAidManager_FinishAid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishAidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidManagerServer).FinishAid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidManager/FinishAid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidManagerServer).FinishAid(ctx, req.(*FinishAidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAidManager_CreateAidMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAidMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAidManagerServer).CreateAidMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mutualaid.UserAidManager/CreateAidMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAidManagerServer).CreateAidMessage(ctx, req.(*CreateAidMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAidManager_ServiceDesc is the grpc.ServiceDesc for UserAidManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAidManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mutualaid.UserAidManager",
	HandlerType: (*UserAidManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAid",
			Handler:    _UserAidManager_CreateAid_Handler,
		},
		{
			MethodName: "CancelAid",
			Handler:    _UserAidManager_CancelAid_Handler,
		},
		{
			MethodName: "FinishAid",
			Handler:    _UserAidManager_FinishAid_Handler,
		},
		{
			MethodName: "CreateAidMessage",
			Handler:    _UserAidManager_CreateAidMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mutualaid/mutualaid.proto",
}