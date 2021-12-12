package service

import (
	"context"
	tryv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/try/v1"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func (s *TryTestSuite) TestTry() {
	cli := tryv1.NewTryServiceClient(clientConnWithLocalServer())
	req := &tryv1.TryRequest{
		Name:  "Geek",
		Score: uint32(99),
		Gender: &tryv1.TryRequest_Male{
			Male: true,
		},
	}
	resp, err := cli.Try(s.ctx, req)
	s.Nil(err)
	checkRPCErr(err)
	logrus.Infof("[client] got response: %+v", resp)
}

func (s *TryTestSuite) TestPortUsed() {
	used := checkLocalPortUsed(59999, true)
	s.False(used)
	_ = clientConnWithLocalServer()
	used = checkLocalPortUsed(defaultLocalRCPPort, true)
	s.True(used)
}

func (s *TryTestSuite) TestFibonacci() {
	cli := tryv1.NewTryServiceClient(clientConnWithLocalServer())
	req := &tryv1.FibonacciRequest{
		N:    40,
		From: "testing",
	}
	resp, err := cli.Fibonacci(s.ctx, req)
	s.Nil(err)
	checkRPCErr(err)
	logrus.Infof("[client] got response: %s", resp)
	s.EqualValues(102334155, resp.Result)
}

//func (s *TryTestSuite) TestFatal() {
//	cli := tryv1.NewTryServiceClient(clientConnWithRemoteSvr())
//	req := &tryv1.FatalRequest{From: "testing"}
//	resp, err := cli.Fatal(s.ctx, req)
//	s.NotNil(err)
//	checkRPCErr(err)
//	logrus.Infof("[client] got response: %+v", resp)
//}

func (s *TryTestSuite) TestSlow() {
	cli := tryv1.NewTryServiceClient(clientConnWithLocalServer())
	req := &tryv1.SlowRequest{
		Guid: time.Now().Format(time.RFC3339Nano),
		NMs:  1_000,
	}
	resp, err := cli.Slow(s.ctx, req)
	s.Nil(err)
	checkRPCErr(err)
	logrus.Infof("[client] got %s", resp)
}

//func (s *TryTestSuite) TestGracefulStop(){
//	cli := tryv1.NewTryServiceClient(clientConnWithRemoteSvr())
//	cnt := 0
//	for {
//		req := &tryv1.SlowRequest{
//			Guid: fmt.Sprintf("%d--%s", cnt, time.Now().Format(time.RFC3339Nano)),
//			NMs: 5_000,
//		}
//		resp, err := cli.Slow(s.ctx, req)
//		s.Nil(err)
//		checkRPCErr(err)
//		logrus.Infof("[client] got %s", resp)
//		if err != nil {
//			break
//		}
//		cnt ++
//	}
//	logrus.Info("graceful-stop over")
//}

func (s *TryTestSuite) TestTimeout() {
	cli := tryv1.NewTryServiceClient(clientConnWithLocalServer())
	s.Run("sever timeout", func() {
		req := &tryv1.TimeoutRequest{
			TimeoutInMs: 200,
		}
		cd := time.Duration(req.TimeoutInMs) * time.Millisecond
		reqCt, cancel := context.WithTimeout(s.ctx, cd)
		defer func() {
			logrus.Infof("[client] client cancel request")
			cancel()
		}()
		resp, err := cli.Timeout(reqCt, req)
		checkRPCErr(err)
		logrus.Infof("[client] got response: %+v", resp)
		time.Sleep(1000 * time.Millisecond)
	})

	s.Run("client early cancel", func() {
		req := &tryv1.TimeoutRequest{
			TimeoutInMs: 400,
		}
		cd := time.Duration(req.TimeoutInMs) * time.Millisecond
		reqCt, cancel := context.WithTimeout(s.ctx, cd)
		go func() {
			time.Sleep(200 * time.Millisecond)
			logrus.Infof("[client] client cancel request")
			cancel()
		}()
		resp, err := cli.Timeout(reqCt, req)
		checkRPCErr(err)
		logrus.Infof("[client] got response: %+v", resp)
		time.Sleep(2000 * time.Millisecond) // should greater than above timeout
	})
}

func (s *TryTestSuite) TestDetailError() {
	cli := tryv1.NewTryServiceClient(clientConnWithLocalServer())
	req := &tryv1.DetailErrorRequest{
		Code: 1234,
		From: "testing",
	}
	r, err := cli.DetailError(s.ctx, req)
	s.Nil(r)
	s.NotNil(err)
	se, ok := status.FromError(err)
	s.True(ok)
	s.EqualValues(codes.FailedPrecondition, se.Code())
	for _, d := range se.Details() {
		switch info := d.(type) {
		case *tryv1.Error:
			s.EqualValues(req.Code, d.(*tryv1.Error).Code)
			logrus.Infof("hit mock error: %+v", info)
		default:
			logrus.Fatalf("Unexpected type: %s", info)
		}
	}
}

func TestTryTestSuite(t *testing.T) {
	suite.Run(t, new(TryTestSuite))
}

type TryTestSuite struct {
	suite.Suite
	ctx context.Context
}

// The SetupSuite method will be run before any tests are run.
func (s *TryTestSuite) SetupSuite() {
	s.ctx = context.Background()
}

// The TearDownSuite method will be run after all tests have been run.
func (s *TryTestSuite) TearDownSuite() {
}
