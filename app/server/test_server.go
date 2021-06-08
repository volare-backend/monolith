package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "backend/api"
)

// TestServiceServer is a composite interface of api_pb.TestServiceServer and grapiserver.Server.
type TestServiceServer interface {
	api_pb.TestServiceServer
	grapiserver.Server
}

// NewTestServiceServer creates a new TestServiceServer instance.
func NewTestServiceServer() TestServiceServer {
	return &testServiceServerImpl{}
}

type testServiceServerImpl struct {
}

func (s *testServiceServerImpl) ListTests(ctx context.Context, req *api_pb.ListTestsRequest) (*api_pb.ListTestsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *testServiceServerImpl) GetTest(ctx context.Context, req *api_pb.GetTestRequest) (*api_pb.Test, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *testServiceServerImpl) CreateTest(ctx context.Context, req *api_pb.CreateTestRequest) (*api_pb.Test, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *testServiceServerImpl) UpdateTest(ctx context.Context, req *api_pb.UpdateTestRequest) (*api_pb.Test, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *testServiceServerImpl) DeleteTest(ctx context.Context, req *api_pb.DeleteTestRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
