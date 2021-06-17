package server

import (
	"context"
	"database/sql"
	"github.com/volare-backend/monolith/infra/repository"
	"github.com/volare-backend/monolith/usecase"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	api_pb "github.com/volare-backend/monolith/api"
)

// SampleServiceServer is a composite interface of api_pb.SampleServiceServer and grapiserver.Server.
type SampleServiceServer interface {
	api_pb.SampleServiceServer
	grapiserver.Server
}

// NewSampleServiceServer creates a new SampleServiceServer instance.
func NewSampleServiceServer(db *sql.DB) SampleServiceServer {
	return &sampleServiceServerImpl{
		sampleUsecase: usecase.NewSampleUsecase(repository.NewSampleRepository(db)),
	}
}

type sampleServiceServerImpl struct {
	sampleUsecase *usecase.SampleUsecase
}

func (s *sampleServiceServerImpl) ListSamples(ctx context.Context, req *api_pb.ListSamplesRequest) (*api_pb.ListSamplesResponse, error) {
	list, err := s.sampleUsecase.GetList()
	if err != nil {
		return nil, err
	}

	var respSamples []*api_pb.Sample

	for _, l := range list {
		respSamples = append(respSamples, &api_pb.Sample{
			Hoge: l.Hoge,
			Name: l.Name,
			Age:  l.Age,
		})
	}

	resp := &api_pb.ListSamplesResponse{
		Samples: respSamples,
	}

	return resp, nil
}

func (s *sampleServiceServerImpl) GetSample(ctx context.Context, req *api_pb.GetSampleRequest) (*api_pb.Sample, error) {
	id, err := strconv.Atoi(req.SampleId)
	if err != nil {
		return nil, err
	}

	e, err := s.sampleUsecase.GetByID(int64(id))
	if err != nil {
		return nil, err
	}

	resp := &api_pb.Sample{
		Hoge: e.Hoge,
		Name: e.Name,
		Age:  e.Age,
	}
	return resp, nil
}

func (s *sampleServiceServerImpl) CreateSample(ctx context.Context, req *api_pb.CreateSampleRequest) (*api_pb.Sample, error) {
	created, err := s.sampleUsecase.Create(req.Sample.Hoge, req.Sample.Name, req.Sample.Age)
	if err != nil {
		return nil, err
	}
	resp := &api_pb.Sample{
		Hoge: created.Hoge,
		Name: created.Name,
		Age:  created.Age,
	}
	return resp, nil
}

func (s *sampleServiceServerImpl) UpdateSample(ctx context.Context, req *api_pb.UpdateSampleRequest) (*api_pb.Sample, error) {
	updated, err := s.sampleUsecase.Update(req.Sample.Hoge, req.Sample.Name, req.Sample.Age, int64(id))
	if err != nil {
		return nil, err
	}

	resp := &api_pb.Sample{
		Hoge: updated.Hoge,
		Name: updated.Name,
		Age:  updated.Age,
	}

	return resp, nil
}

func (s *sampleServiceServerImpl) DeleteSample(ctx context.Context, req *api_pb.DeleteSampleRequest) (*empty.Empty, error) {
	err = s.sampleUsecase.Delete(int64(id))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
