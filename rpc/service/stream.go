package service

import (
	"errors"
	"fmt"
	streamv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/stream/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"sync"
	"time"
)

type StreamServer struct {
	streamv1.UnimplementedStreamServiceServer
}

func (s StreamServer) Hi(req *streamv1.HiRequest, svr streamv1.StreamService_HiServer) error {
	logrus.Infof("[server] handling request with: %+v", req)
	for i := int32(0); i < req.MsgCount; i++ {
		hr := &streamv1.HiResponse{
			Message: fmt.Sprintf("index=%d response for %s request", i, req.From),
			Index:   i,
		}
		err := svr.Send(hr)
		if err != nil {
			return err
		}
		logrus.Infof("[server] sent %d msg: %+v", i, hr)
	}
	return nil
}

func (s StreamServer) ClientStream(cs streamv1.StreamService_ClientStreamServer) error {
	cnt := 0
	for {
		csr, err := cs.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				logrus.Errorf("Recv() unexecpted error: %s", err.Error())
				return status.Errorf(codes.Internal, "Recv() error: %s", err.Error())
			}
			break
		}
		logrus.Infof("[server] got %d from idx=%d request: %+v", cnt, csr.Index, csr)
		cnt++
	}
	return nil
}

func (s StreamServer) BiStream(svr streamv1.StreamService_BiStreamServer) error {
	// 应该设置超时时间，防止过长不活跃连接 长期占用系统资源

	// first put a WELCOME message
	wMsg := fmt.Sprintf("Welcome %s", time.Now().Format(time.RFC3339))
	_ = svr.Send(&streamv1.BiStreamResponse{Message: wMsg})

	cnt := 0
	for {
		logrus.Infof("[server] waiting count: %d", cnt)
		bsr, err := svr.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				logrus.Errorf("[server] Recv() unexecpted error: %s", err.Error())
			}
			break
		}
		logrus.Infof("[server] got %d from idx=%d request: %+v", cnt, bsr.Index, bsr)
		cIdx := bsr.Index
		if cIdx%2 == 0 {
			magic := cIdx * cIdx
			bResp := &streamv1.BiStreamResponse{
				Message: fmt.Sprintf("index=%d hit magic: %d", cIdx, magic),
				Index:   cIdx,
			}
			err = svr.Send(bResp)
			if err != nil {
				logrus.Errorf("[server] Send() unexecpted error: %s", err.Error())
			}
		}
		cnt++
	}
	logrus.Infof("[server] response stream end")
	return nil
}

func (s StreamServer) BiStream2(svr streamv1.StreamService_BiStreamServer) error {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// pull messages
	go func() {
		cnt := 0
		for {
			csr, err := svr.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					logrus.Errorf("[server] Recv() unexecpted error: %s", err.Error())
				}
				break
			}
			logrus.Infof("[server] got %d from idx=%d request: %+v", cnt, csr.Index, csr)
			cnt++
		}
		wg.Done()
	}()

	// server stream, push messages
	go func() {
		for i := int32(0); i < 3; i++ {
			hr := &streamv1.BiStreamResponse{
				Message: fmt.Sprintf("index=%d response message", i),
				Index:   i,
			}
			err := svr.Send(hr)
			if err != nil {
				logrus.Errorf("[server] Send() unexpected error: %s", err.Error())
				break
			}
			logrus.Infof("[server] sent %d msg: %+v", i, hr)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
