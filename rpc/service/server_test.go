package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func (s *ServerTestSuite) TestStatusCode() {
	reason := "testing"
	err := status.Errorf(codes.FailedPrecondition, "failed precondition %s", reason)

	stErr, ok := status.FromError(err)
	s.True(ok)
	s.Equal(codes.FailedPrecondition, stErr.Code())
	// "message: failed precondition testing with code: FailedPrecondition"
	logrus.Infof("message: %s with code: %v", stErr.Message(), stErr.Code())
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

type ServerTestSuite struct {
	suite.Suite
	clientConn *grpc.ClientConn
	ctx        context.Context
}

// The SetupSuite method will be run before any tests are run.
func (s *ServerTestSuite) SetupSuite() {
	s.clientConn = clientConnWithLocalServer()
	s.ctx = context.Background()
}

// The TearDownSuite method will be run after all tests have been run.
func (s *ServerTestSuite) TearDownSuite() {
}
