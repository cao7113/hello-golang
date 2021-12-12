package service

import (
	"context"
	"fmt"
	hellov1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/hello/v1"
	"github.com/sirupsen/logrus"
)

type HelloServer struct {
	hellov1.UnimplementedHelloServiceServer
}

func (s HelloServer) Hello(ctx context.Context, req *hellov1.HelloRequest) (*hellov1.HelloResponse, error) {
	logrus.Infof("[server] handling hello-request with %+v", req)
	resp := &hellov1.HelloResponse{
		Message: fmt.Sprintf("Welcome %s", req.From),
	}
	return resp, nil
}
