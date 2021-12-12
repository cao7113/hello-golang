package service

import (
	"context"
	"errors"
	streamv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/stream/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

func (s *ServerTestSuite) TestHiStream() {
	cli := streamv1.NewStreamServiceClient(s.clientConn)
	ctx := context.Background()
	req := &streamv1.HiRequest{
		From:     "testing",
		MsgCount: 5,
	}
	logrus.Infof("[client] requesting with %+v", req)
	rs, err := cli.Hi(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logrus.Fatalf("error: code: %d message: %s", st.Code(), st.Message())
		} else {
			logrus.Fatalf("Hi error: %v", err)
		}
	}

	if rs != nil {
		cnt := 0
		for {
			hr, err := rs.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					logrus.Infof("[client] got %d messages and stream end", cnt)
				} else {
					logrus.Errorf("Recv() unexpected error: %v", err)
				}
				break
			}
			if hr != nil {
				logrus.Infof("[client] got idx=%d response: %+v", hr.Index, hr)
			}
			cnt++
			//time.Sleep(10 * time.Millisecond)
		}
		s.EqualValues(req.MsgCount, cnt)
	}
}

func (s *ServerTestSuite) TestClientStream() {
	ctx := context.Background()
	cli := streamv1.NewStreamServiceClient(s.clientConn)
	cs, err := cli.ClientStream(ctx)
	if err != nil {
		logrus.Fatalf("ClientStream fatal: %s", err.Error())
	}
	cnt := 3
	for i := 0; i < cnt; i++ {
		req := &streamv1.ClientStreamRequest{
			From:  "testing",
			Index: int32(i),
		}
		err = cs.Send(req)
		if err != nil {
			logrus.Fatalf("Send() error: %s", err.Error())
		}
		logrus.Infof("[client] sent msg: %+v", req)
	}
	time.Sleep(2 * time.Second)
}

func (s *ServerTestSuite) TestBiStream() {
	ctx := context.Background()
	cli := streamv1.NewStreamServiceClient(s.clientConn)
	bs, err := cli.BiStream(ctx)
	if err != nil {
		logrus.Fatalf("ClientStream fatal: %s", err.Error())
	}

	// push messages
	go func() {
		cnt := 5
		for i := 0; i < cnt; i++ {
			req := &streamv1.BiStreamRequest{
				From:  "testing",
				Index: int32(i),
			}
			err = bs.Send(req)
			if err != nil {
				logrus.Fatalf("Send() error: %s", err.Error())
			}
			logrus.Infof("[client] sent msg: %+v", req)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// get messages
	go func() {
		for {
			bsResp, err := bs.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					logrus.Errorf("[client] Recv() error: %s", err.Error())
				}
				break
			}
			logrus.Infof("[client] got response: %+v", bsResp)
		}
	}()

	time.Sleep(3 * time.Second)
	//select {}
}
