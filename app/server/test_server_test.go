package server

import (
	"context"
	"testing"

	api_pb "backend/api"
)

func Test_TestServiceServer_ListTests(t *testing.T) {
	svr := NewTestServiceServer()

	ctx := context.Background()
	req := &api_pb.ListTestsRequest{}

	resp, err := svr.ListTests(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TestServiceServer_GetTest(t *testing.T) {
	svr := NewTestServiceServer()

	ctx := context.Background()
	req := &api_pb.GetTestRequest{}

	resp, err := svr.GetTest(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TestServiceServer_CreateTest(t *testing.T) {
	svr := NewTestServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateTestRequest{}

	resp, err := svr.CreateTest(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TestServiceServer_UpdateTest(t *testing.T) {
	svr := NewTestServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateTestRequest{}

	resp, err := svr.UpdateTest(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_TestServiceServer_DeleteTest(t *testing.T) {
	svr := NewTestServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteTestRequest{}

	resp, err := svr.DeleteTest(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
