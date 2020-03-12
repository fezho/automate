// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/tokens.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/tokens.proto", fileDescriptor_210e42bd7205e452)
}

var fileDescriptor_210e42bd7205e452 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xb1, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0xe5, 0x50, 0xac, 0xe8, 0x52, 0x42, 0x7b, 0x55, 0x25, 0xcb, 0x62, 0x88, 0x2c, 0x01,
	0x25, 0x6a, 0x6c, 0x91, 0x6e, 0x01, 0x09, 0x02, 0x48, 0x48, 0x08, 0x51, 0x54, 0x60, 0xa1, 0xaa,
	0xc4, 0xd5, 0xfe, 0x70, 0x5c, 0x6c, 0xdf, 0xd5, 0x77, 0x6e, 0xa9, 0xaa, 0x02, 0x0a, 0x82, 0x21,
	0x62, 0xa1, 0x4c, 0x48, 0x8c, 0x2c, 0x3c, 0x40, 0x5e, 0x80, 0x8d, 0x11, 0xb1, 0xf0, 0x00, 0x4c,
	0x6c, 0xbc, 0x01, 0xf2, 0xb9, 0x49, 0x6c, 0x92, 0x50, 0x77, 0x8a, 0xfc, 0xf9, 0xf7, 0xf9, 0xbe,
	0xff, 0x3f, 0xff, 0xbb, 0x43, 0x2b, 0x36, 0x0d, 0x18, 0x0d, 0x21, 0x14, 0xdc, 0x22, 0xb1, 0xa0,
	0x01, 0x11, 0xd0, 0x70, 0x89, 0x80, 0x5d, 0xb2, 0x67, 0x11, 0xe6, 0x59, 0x1e, 0x09, 0xac, 0x9d,
	0xa6, 0x25, 0xe8, 0x33, 0x08, 0xb9, 0xc9, 0x22, 0x2a, 0x28, 0xd6, 0xec, 0x0e, 0x3c, 0x35, 0x07,
	0xb8, 0x49, 0x98, 0x67, 0x7a, 0x24, 0x30, 0x77, 0x9a, 0xfa, 0x39, 0x97, 0x52, 0xd7, 0x07, 0xd9,
	0x49, 0xc2, 0x90, 0x0a, 0x22, 0x3c, 0x3a, 0xe8, 0xd3, 0x97, 0xe5, 0x8f, 0xdd, 0x70, 0x21, 0x6c,
	0xf0, 0x5d, 0xe2, 0xba, 0x10, 0x59, 0x94, 0x49, 0x62, 0x02, 0x7d, 0xa5, 0xe0, 0x68, 0x11, 0x6c,
	0xc7, 0xc0, 0x45, 0x6e, 0x44, 0xfd, 0x6a, 0xe1, 0x66, 0xce, 0x68, 0xc8, 0x21, 0xdf, 0x7d, 0x7d,
	0x62, 0x77, 0xc4, 0x6c, 0x2b, 0xa3, 0x80, 0x51, 0xdf, 0xb3, 0xf7, 0xa6, 0x48, 0x3d, 0xc9, 0x17,
	0x92, 0x51, 0xc6, 0xbe, 0xd0, 0xfc, 0x32, 0x8b, 0xd4, 0x87, 0x72, 0x28, 0xfc, 0xbb, 0x84, 0x2a,
	0x37, 0x23, 0x20, 0x02, 0x64, 0x01, 0x2f, 0x99, 0xd3, 0xfe, 0x00, 0x33, 0x83, 0xad, 0xc1, 0xb6,
	0x7e, 0xa9, 0x20, 0xc9, 0x99, 0xf1, 0xa6, 0x74, 0xd8, 0xfe, 0xac, 0x20, 0x35, 0x35, 0x62, 0xeb,
	0xa3, 0x82, 0xaa, 0xcf, 0x1b, 0x36, 0x75, 0xa0, 0xc1, 0x49, 0xc0, 0x7c, 0xe0, 0xf8, 0xad, 0xd2,
	0x7c, 0xad, 0xa0, 0x57, 0x4a, 0xfd, 0x05, 0xaa, 0xa2, 0x19, 0x9f, 0x84, 0x2e, 0x56, 0xf5, 0x99,
	0x3b, 0x0f, 0x56, 0xef, 0x21, 0x1f, 0xa9, 0x9c, 0xc6, 0x91, 0x0d, 0x78, 0x53, 0x7f, 0xb2, 0x6f,
	0x84, 0x24, 0x00, 0xa3, 0x55, 0x33, 0xe4, 0xb7, 0x6a, 0x97, 0x8d, 0xe5, 0x9a, 0xe1, 0x39, 0xc3,
	0x42, 0x43, 0x16, 0x88, 0x2d, 0xbc, 0x9d, 0x84, 0x12, 0x51, 0x0c, 0xcb, 0x35, 0x83, 0x45, 0x74,
	0x0b, 0x6c, 0xc1, 0x8d, 0x56, 0x6d, 0xdd, 0x00, 0xc2, 0x45, 0x23, 0x02, 0xd7, 0xa3, 0x61, 0xc2,
	0xee, 0xc2, 0xe8, 0x71, 0xe3, 0xa0, 0xdb, 0xd7, 0xce, 0xa0, 0x0a, 0x89, 0x45, 0xa7, 0x95, 0x0e,
	0xdb, 0xed, 0x6b, 0x65, 0xac, 0xda, 0x52, 0x50, 0xaf, 0xaf, 0xcd, 0x22, 0xe4, 0x91, 0xe0, 0xe8,
	0x5d, 0xaf, 0xaf, 0x2d, 0xe0, 0xf9, 0xd1, 0x73, 0x2b, 0xc5, 0xba, 0x3f, 0x7e, 0x7d, 0x28, 0x2d,
	0x18, 0xd5, 0x7c, 0xbe, 0x5b, 0x4a, 0x1d, 0x7f, 0x53, 0x50, 0xf9, 0x36, 0x88, 0xd4, 0xe9, 0xf3,
	0xd3, 0xfd, 0x1b, 0x30, 0x89, 0xcd, 0x17, 0x8a, 0x60, 0x9c, 0x19, 0xe1, 0x61, 0xbb, 0x3c, 0x70,
	0xb8, 0xdb, 0xd7, 0x30, 0x9a, 0xcb, 0xa8, 0x68, 0xed, 0x7b, 0x4e, 0xa2, 0xed, 0x34, 0x3e, 0xe5,
	0x82, 0xe8, 0xf5, 0xb5, 0x79, 0x74, 0x36, 0x33, 0x77, 0xf2, 0xb6, 0xd7, 0xd7, 0xe6, 0x70, 0x35,
	0x53, 0x74, 0x41, 0x48, 0x25, 0x8b, 0x78, 0x21, 0xaf, 0xc4, 0x4a, 0x70, 0xfc, 0xa7, 0x84, 0x2a,
	0x8f, 0x98, 0x53, 0x24, 0x38, 0x19, 0xec, 0x98, 0xe0, 0xe4, 0x48, 0xce, 0x8c, 0x77, 0xa5, 0xc3,
	0xf6, 0xa7, 0x51, 0x70, 0xde, 0x8f, 0x07, 0xe7, 0x65, 0xf3, 0x00, 0xed, 0xd7, 0xf7, 0xc6, 0x52,
	0xe3, 0x0e, 0x53, 0xb3, 0xa1, 0xaf, 0x8f, 0x52, 0x13, 0xcb, 0x05, 0x9c, 0x5a, 0x9a, 0x1e, 0x59,
	0x3d, 0x61, 0x5e, 0x38, 0x8d, 0x45, 0x27, 0x17, 0x98, 0xc9, 0x56, 0x97, 0xb1, 0x9a, 0x2e, 0x36,
	0xcd, 0xed, 0x7c, 0x74, 0x52, 0x56, 0x1a, 0xae, 0xe9, 0x93, 0x0c, 0x4f, 0xf2, 0xf3, 0x53, 0x41,
	0x95, 0x5b, 0xe0, 0x43, 0x01, 0xcf, 0x33, 0xd8, 0x31, 0x9e, 0xe7, 0x48, 0xce, 0x0c, 0x51, 0x24,
	0x48, 0x65, 0xac, 0x3a, 0xb2, 0xaf, 0x98, 0xba, 0x94, 0x4d, 0xe3, 0x54, 0x9f, 0x18, 0xa7, 0xaf,
	0x0a, 0x42, 0x77, 0x3d, 0x2e, 0x8e, 0x8e, 0xa5, 0x8b, 0xd3, 0xe7, 0x1d, 0x51, 0x89, 0xb0, 0xa5,
	0x62, 0x20, 0x67, 0x06, 0xc9, 0xeb, 0x1a, 0xdb, 0xe6, 0x2a, 0x9e, 0x89, 0x80, 0x38, 0x13, 0x36,
	0xf9, 0x3c, 0xce, 0x0a, 0xf4, 0x3d, 0x9e, 0x6e, 0x8c, 0x39, 0xfc, 0xcf, 0x16, 0xc7, 0xdf, 0x15,
	0xb4, 0xb8, 0x06, 0x1c, 0x44, 0xdb, 0xf7, 0xe5, 0xca, 0xf7, 0x8f, 0x82, 0x84, 0x9b, 0xd3, 0xc7,
	0x9c, 0xd8, 0x90, 0x48, 0x5b, 0x39, 0x71, 0x0f, 0x67, 0xc6, 0xaa, 0x54, 0xd9, 0xf1, 0x1c, 0x07,
	0xc2, 0x29, 0x87, 0xd9, 0x30, 0x96, 0xff, 0x3f, 0xcc, 0x52, 0xec, 0x46, 0xfb, 0xf1, 0x35, 0xd7,
	0x13, 0x9d, 0x78, 0xd3, 0xb4, 0x69, 0x60, 0x25, 0x13, 0x0d, 0x2f, 0x1d, 0xab, 0xd8, 0x45, 0xb8,
	0xa9, 0xca, 0x5b, 0x67, 0xe5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x8e, 0x5f, 0x91, 0x11,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokensClient is the client API for Tokens service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokensClient interface {
	//
	//Creates a token
	//
	//Creates a token.
	//Active defaults to true when not specified.
	//Value is auto-generated when not specified.
	//
	//Note that this creates *non-admin* tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*``).
	//
	//You cannot create admin tokens via the REST API.
	//Admin tokens can only be created by specifying the `--admin` flag to this chef-automate sub-command:
	//```
	//chef-automate iam token create <your-token-name> --admin`
	//```
	//
	//Authorization Action:
	//```
	//iam:tokens:create
	//```
	CreateToken(ctx context.Context, in *request.CreateTokenReq, opts ...grpc.CallOption) (*response.CreateTokenResp, error)
	//
	//Gets a token
	//
	//Returns the details for a token.
	//
	//Authorization Action:
	//```
	//iam:tokens:get
	//```
	GetToken(ctx context.Context, in *request.GetTokenReq, opts ...grpc.CallOption) (*response.GetTokenResp, error)
	//
	//Updates a token
	//
	//This operation overwrites all fields excepting ID, timestamps, and value,
	//including those omitted from the request, so be sure to specify all properties.
	//Properties that you do not include are reset to empty values.
	//
	//Authorization Action:
	//```
	//iam:tokens:update
	//```
	UpdateToken(ctx context.Context, in *request.UpdateTokenReq, opts ...grpc.CallOption) (*response.UpdateTokenResp, error)
	//
	//Deletes a token
	//
	//Deletes a token and remove it from any policies.
	//
	//Authorization Action:
	//```
	//iam:tokens:delete
	//```
	DeleteToken(ctx context.Context, in *request.DeleteTokenReq, opts ...grpc.CallOption) (*response.DeleteTokenResp, error)
	//
	//Lists all tokens
	//
	//Lists all tokens, both admin and non-admin.
	//
	//Authorization Action:
	//```
	//iam:tokens:list
	//```
	ListTokens(ctx context.Context, in *request.ListTokensReq, opts ...grpc.CallOption) (*response.ListTokensResp, error)
	// Expose on GRPC API only so we don't expose this to the enduser.
	// Just want to be able to trigger this via automate-cli.
	ResetAllTokenProjects(ctx context.Context, in *request.ResetAllTokenProjectsReq, opts ...grpc.CallOption) (*response.ResetAllTokenProjectsResp, error)
}

type tokensClient struct {
	cc grpc.ClientConnInterface
}

func NewTokensClient(cc grpc.ClientConnInterface) TokensClient {
	return &tokensClient{cc}
}

func (c *tokensClient) CreateToken(ctx context.Context, in *request.CreateTokenReq, opts ...grpc.CallOption) (*response.CreateTokenResp, error) {
	out := new(response.CreateTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) GetToken(ctx context.Context, in *request.GetTokenReq, opts ...grpc.CallOption) (*response.GetTokenResp, error) {
	out := new(response.GetTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) UpdateToken(ctx context.Context, in *request.UpdateTokenReq, opts ...grpc.CallOption) (*response.UpdateTokenResp, error) {
	out := new(response.UpdateTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) DeleteToken(ctx context.Context, in *request.DeleteTokenReq, opts ...grpc.CallOption) (*response.DeleteTokenResp, error) {
	out := new(response.DeleteTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) ListTokens(ctx context.Context, in *request.ListTokensReq, opts ...grpc.CallOption) (*response.ListTokensResp, error) {
	out := new(response.ListTokensResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/ListTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) ResetAllTokenProjects(ctx context.Context, in *request.ResetAllTokenProjectsReq, opts ...grpc.CallOption) (*response.ResetAllTokenProjectsResp, error) {
	out := new(response.ResetAllTokenProjectsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/ResetAllTokenProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokensServer is the server API for Tokens service.
type TokensServer interface {
	//
	//Creates a token
	//
	//Creates a token.
	//Active defaults to true when not specified.
	//Value is auto-generated when not specified.
	//
	//Note that this creates *non-admin* tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*``).
	//
	//You cannot create admin tokens via the REST API.
	//Admin tokens can only be created by specifying the `--admin` flag to this chef-automate sub-command:
	//```
	//chef-automate iam token create <your-token-name> --admin`
	//```
	//
	//Authorization Action:
	//```
	//iam:tokens:create
	//```
	CreateToken(context.Context, *request.CreateTokenReq) (*response.CreateTokenResp, error)
	//
	//Gets a token
	//
	//Returns the details for a token.
	//
	//Authorization Action:
	//```
	//iam:tokens:get
	//```
	GetToken(context.Context, *request.GetTokenReq) (*response.GetTokenResp, error)
	//
	//Updates a token
	//
	//This operation overwrites all fields excepting ID, timestamps, and value,
	//including those omitted from the request, so be sure to specify all properties.
	//Properties that you do not include are reset to empty values.
	//
	//Authorization Action:
	//```
	//iam:tokens:update
	//```
	UpdateToken(context.Context, *request.UpdateTokenReq) (*response.UpdateTokenResp, error)
	//
	//Deletes a token
	//
	//Deletes a token and remove it from any policies.
	//
	//Authorization Action:
	//```
	//iam:tokens:delete
	//```
	DeleteToken(context.Context, *request.DeleteTokenReq) (*response.DeleteTokenResp, error)
	//
	//Lists all tokens
	//
	//Lists all tokens, both admin and non-admin.
	//
	//Authorization Action:
	//```
	//iam:tokens:list
	//```
	ListTokens(context.Context, *request.ListTokensReq) (*response.ListTokensResp, error)
	// Expose on GRPC API only so we don't expose this to the enduser.
	// Just want to be able to trigger this via automate-cli.
	ResetAllTokenProjects(context.Context, *request.ResetAllTokenProjectsReq) (*response.ResetAllTokenProjectsResp, error)
}

// UnimplementedTokensServer can be embedded to have forward compatible implementations.
type UnimplementedTokensServer struct {
}

func (*UnimplementedTokensServer) CreateToken(ctx context.Context, req *request.CreateTokenReq) (*response.CreateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokensServer) GetToken(ctx context.Context, req *request.GetTokenReq) (*response.GetTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedTokensServer) UpdateToken(ctx context.Context, req *request.UpdateTokenReq) (*response.UpdateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedTokensServer) DeleteToken(ctx context.Context, req *request.DeleteTokenReq) (*response.DeleteTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedTokensServer) ListTokens(ctx context.Context, req *request.ListTokensReq) (*response.ListTokensResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}
func (*UnimplementedTokensServer) ResetAllTokenProjects(ctx context.Context, req *request.ResetAllTokenProjectsReq) (*response.ResetAllTokenProjectsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetAllTokenProjects not implemented")
}

func RegisterTokensServer(s *grpc.Server, srv TokensServer) {
	s.RegisterService(&_Tokens_serviceDesc, srv)
}

func _Tokens_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).CreateToken(ctx, req.(*request.CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).GetToken(ctx, req.(*request.GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).UpdateToken(ctx, req.(*request.UpdateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).DeleteToken(ctx, req.(*request.DeleteTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/ListTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).ListTokens(ctx, req.(*request.ListTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_ResetAllTokenProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ResetAllTokenProjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).ResetAllTokenProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/ResetAllTokenProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).ResetAllTokenProjects(ctx, req.(*request.ResetAllTokenProjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokens_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Tokens",
	HandlerType: (*TokensServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Tokens_CreateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Tokens_GetToken_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _Tokens_UpdateToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Tokens_DeleteToken_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _Tokens_ListTokens_Handler,
		},
		{
			MethodName: "ResetAllTokenProjects",
			Handler:    _Tokens_ResetAllTokenProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2/tokens.proto",
}
