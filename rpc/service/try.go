package service

import (
	"context"
	"fmt"
	tryv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/try/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type TryServer struct {
	tryv1.UnimplementedTryServiceServer
}

func (s TryServer) Try(ctx context.Context, req *tryv1.TryRequest) (*tryv1.TryResponse, error) {
	logrus.Infof("[server] handling request with: %+v", req)
	msg := req.Name
	switch req.Gender.(type) {
	case *tryv1.TryRequest_Male:
		msg += " Male"
	case *tryv1.TryRequest_Female:
		msg += " Female"
	}
	msg += fmt.Sprintf(" Score: %d", req.Score)
	resp := &tryv1.TryResponse{Message: msg}
	return resp, nil
}

func (s TryServer) Fibonacci(ctx context.Context, req *tryv1.FibonacciRequest) (*tryv1.FibonacciResponse, error) {
	logrus.Infof("[server] handle requesting with %+v", req)
	t0 := time.Now()
	result := fibN(req.N)
	du := time.Since(t0).Milliseconds()
	rp := &tryv1.FibonacciResponse{
		Result:  result,
		TakenMs: du,
	}
	return rp, nil
}

func (s TryServer) Slow(ctx context.Context, req *tryv1.SlowRequest) (*tryv1.SlowResponse, error) {
	logrus.Infof("[server] handle requesting guid=%s with %+v", req.Guid, req)
	time.Sleep(time.Duration(req.NMs) * time.Millisecond)
	resp := &tryv1.SlowResponse{
		Msg: fmt.Sprintf("reply guid=%s, after %d ms", req.Guid, req.NMs),
	}
	return resp, nil
}

func (s TryServer) Timeout(ctx context.Context, req *tryv1.TimeoutRequest) (*tryv1.TimeoutResponse, error) {
	logrus.Infof("[server] handling request with %+v", req)
	t0 := time.Now()
	i := 0
	for { // 检测超时
		if err := ctx.Err(); err != nil {
			logrus.Warnf("[server] hit error: %s after %d loop", err.Error(), i)
			break
		}
		logrus.Infof("[server] run %d loop", i)
		time.Sleep(100 * time.Millisecond)
		i++
	}
	resp := &tryv1.TimeoutResponse{
		Msg: fmt.Sprintf("reply after %d ms with client %d ms", time.Since(t0).Milliseconds(), req.TimeoutInMs),
	}
	return resp, nil
}

func (s TryServer) DetailError(ctx context.Context, req *tryv1.DetailErrorRequest) (*tryv1.DetailErrorResponse, error) {
	logrus.Infof("[server] handling request: %+v", req)
	st := status.New(codes.FailedPrecondition, "failed to satisfy pre-conditions")
	ds, err := st.WithDetails(
		&tryv1.Error{
			Code:    req.Code,
			Message: fmt.Sprintf("mock msg for code: %d", req.Code),
		},
	)
	if err != nil {
		return nil, st.Err()
	}
	return nil, ds.Err()
}

func (s TryServer) Fatal(ctx context.Context, req *tryv1.FatalRequest) (*tryv1.FatalResponse, error) {
	logrus.Fatalf("[server] handling fatal-request: %+v", req)
	return &tryv1.FatalResponse{Msg: "should not got this"}, nil
}

func fibN(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return fibN(n-1) + fibN(n-2)
}
