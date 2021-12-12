package service

import (
	"context"
	pingv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/ping/v1"
	"github.com/sirupsen/logrus"
)

func (s *ServerTestSuite) TestPing() {
	cli := pingv1.NewPingServiceClient(s.clientConn)
	ctx := context.Background()
	req := &pingv1.PingRequest{
		From: "testing",
	}
	resp, err := cli.Ping(ctx, req)
	s.Nil(err)
	logrus.Infof("[client] got response: %+v", resp)
}
