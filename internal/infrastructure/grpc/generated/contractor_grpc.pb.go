// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: internal/infrastructure/grpc/contractor.proto

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Contractor_CreateJobProposal_FullMethodName = "/grpc.Contractor/createJobProposal"
)

// ContractorClient is the client API for Contractor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractorClient interface {
	CreateJobProposal(ctx context.Context, in *ProposalRequest, opts ...grpc.CallOption) (*LawyersResponse, error)
}

type contractorClient struct {
	cc grpc.ClientConnInterface
}

func NewContractorClient(cc grpc.ClientConnInterface) ContractorClient {
	return &contractorClient{cc}
}

func (c *contractorClient) CreateJobProposal(ctx context.Context, in *ProposalRequest, opts ...grpc.CallOption) (*LawyersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LawyersResponse)
	err := c.cc.Invoke(ctx, Contractor_CreateJobProposal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractorServer is the server API for Contractor service.
// All implementations must embed UnimplementedContractorServer
// for forward compatibility.
type ContractorServer interface {
	CreateJobProposal(context.Context, *ProposalRequest) (*LawyersResponse, error)
	mustEmbedUnimplementedContractorServer()
}

// UnimplementedContractorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedContractorServer struct{}

func (UnimplementedContractorServer) CreateJobProposal(context.Context, *ProposalRequest) (*LawyersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJobProposal not implemented")
}
func (UnimplementedContractorServer) mustEmbedUnimplementedContractorServer() {}
func (UnimplementedContractorServer) testEmbeddedByValue()                    {}

// UnsafeContractorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractorServer will
// result in compilation errors.
type UnsafeContractorServer interface {
	mustEmbedUnimplementedContractorServer()
}

func RegisterContractorServer(s grpc.ServiceRegistrar, srv ContractorServer) {
	// If the following call pancis, it indicates UnimplementedContractorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Contractor_ServiceDesc, srv)
}

func _Contractor_CreateJobProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractorServer).CreateJobProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contractor_CreateJobProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractorServer).CreateJobProposal(ctx, req.(*ProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Contractor_ServiceDesc is the grpc.ServiceDesc for Contractor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contractor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Contractor",
	HandlerType: (*ContractorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createJobProposal",
			Handler:    _Contractor_CreateJobProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/infrastructure/grpc/contractor.proto",
}