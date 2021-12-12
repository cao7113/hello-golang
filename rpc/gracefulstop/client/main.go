package main

import (
	"context"
	greeterv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/greeter/v1"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	pCtx := context.Background()
	ctx, cancel := context.WithTimeout(pCtx, 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx,
		"localhost:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		logrus.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()
	cli := greeterv1.NewGreeterServiceClient(conn)
	req := &greeterv1.HelloRequest{}
	stream, err := cli.SayHelloStream(ctx, req)
	if err != nil {
		logrus.Fatalf("could not create streaming client: %v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case s := <-sigCh:
			logrus.Infof("[client] got signal %v, attempting graceful shutdown", s)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		resp, err := cli.SayHello(ctx, &empty.Empty{})
		if err != nil {
			if status.Code(err) == codes.Canceled {
				logrus.Infoln("SayHello context cancelled")
				return
			}
			logrus.Fatalf("could not perform regular rpc request: %v", err)
		}
		logrus.Infof("received SayHello response: %+v", resp)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			r, err := stream.Recv()
			if err != nil {
				if err == io.EOF || status.Code(err) == codes.Canceled {
					logrus.Infoln("stream closed (context cancelled)")
					cancel()
					return
				}
				logrus.Fatalf("error while receiving stream response: %v", err)
			}
			logrus.Infof("received value: %+v", r)
		}
	}()
	wg.Wait()
}
