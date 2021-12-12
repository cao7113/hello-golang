package service

import (
	"context"
	hellov1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/hello/v1"
	"github.com/sirupsen/logrus"
)

func (s *ServerTestSuite) TestHello() {
	cli := hellov1.NewHelloServiceClient(s.clientConn)
	ctx := context.Background()
	req := &hellov1.HelloRequest{
		From: "testing",
	}
	resp, err := cli.Hello(ctx, req)
	s.Nil(err)
	logrus.Infof("[client] got response: %+v", resp)
}
