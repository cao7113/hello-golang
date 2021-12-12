package service

import (
	"context"
	healthv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/health/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HealthServer struct {
	healthv1.UnimplementedHealthServer
}

func (h HealthServer) Check(ctx context.Context, req *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	resp := &healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_UNKNOWN,
	}
	switch req.Service {
	default:
		resp.Status = healthv1.HealthCheckResponse_SERVING
	}
	return resp, nil
}

func (h HealthServer) Watch(req *healthv1.HealthCheckRequest, svr healthv1.Health_WatchServer) error {
	resp := &healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_UNKNOWN,
	}
	switch req.Service {
	default:
		resp.Status = healthv1.HealthCheckResponse_SERVING
	}
	err := svr.Send(resp)
	if err != nil {
		logrus.Errorf("Send() error: %s", err.Error())
		return status.Errorf(codes.Internal, "Send() error: %s", err.Error())
	}
	return nil
}
