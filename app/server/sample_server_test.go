package server

import (
	"context"
	"testing"

	api_pb "monolith/api"
)

func Test_SampleServiceServer_ListSamples(t *testing.T) {
	svr := NewSampleServiceServer()

	ctx := context.Background()
	req := &api_pb.ListSamplesRequest{}

	resp, err := svr.ListSamples(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SampleServiceServer_GetSample(t *testing.T) {
	svr := NewSampleServiceServer()

	ctx := context.Background()
	req := &api_pb.GetSampleRequest{}

	resp, err := svr.GetSample(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SampleServiceServer_CreateSample(t *testing.T) {
	svr := NewSampleServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateSampleRequest{}

	resp, err := svr.CreateSample(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SampleServiceServer_UpdateSample(t *testing.T) {
	svr := NewSampleServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateSampleRequest{}

	resp, err := svr.UpdateSample(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_SampleServiceServer_DeleteSample(t *testing.T) {
	svr := NewSampleServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteSampleRequest{}

	resp, err := svr.DeleteSample(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
