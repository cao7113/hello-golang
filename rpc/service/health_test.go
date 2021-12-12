package service

import (
	healthv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/health/v1"
	"github.com/sirupsen/logrus"
)

func (s *ServerTestSuite) TestHealthCheck() {
	cli := healthv1.NewHealthClient(s.clientConn)
	req := &healthv1.HealthCheckRequest{
		Service: "",
	}
	resp, err := cli.Check(s.ctx, req)
	s.Nil(err)
	logrus.Infof("[client] got response: %+v", resp)
}
