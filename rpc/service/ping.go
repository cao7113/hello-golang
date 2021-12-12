package service

import (
	"context"
	"fmt"
	pingv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/ping/v1"
	"github.com/sirupsen/logrus"
	"time"
)

type PingServer struct {
	pingv1.UnimplementedPingServiceServer
}

func (s PingServer) Ping(ctx context.Context, req *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	logrus.Infof("handling ping-request with %+v", req)
	msg := fmt.Sprintf("pong for request from %s", req.From)
	resp := &pingv1.PingResponse{
		Message:   msg,
		Timestamp: time.Now().Unix(),
	}
	return resp, nil
}
