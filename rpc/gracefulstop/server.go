package main

import (
	"context"
	"fmt"
	greeterv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/greeter/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type greeterServer struct {
	intCh <-chan int
	greeterv1.UnimplementedGreeterServiceServer
}

func generate(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-time.Tick(time.Second):
				rand.Seed(time.Now().UnixNano())
				n := rand.Int()
				// log.Printf("generated %d", n)
				ch <- n
			case <-ctx.Done():
				logrus.Infof("context hit done")
				close(ch)
				return
			}
		}
	}()
	return ch
}

func (g *greeterServer) SayHelloStream(req *greeterv1.HelloRequest, stream greeterv1.GreeterService_SayHelloStreamServer) error {
	for n := range g.intCh {
		resp := greeterv1.HelloReply{
			Message: strconv.Itoa(n),
		}

		if err := stream.Send(&resp); err != nil {
			if status.Code(err) == codes.Canceled {
				logrus.Infoln("stream closed (context cancelled)")
				return nil
			}
			logrus.Infof("could not send over stream: %v", err)
			return err
		}
		logrus.Infof("sent %d", n)
	}
	return nil
}

func (g *greeterServer) SayHello(context.Context, *empty.Empty) (*greeterv1.HelloReply, error) {
	resp := greeterv1.HelloReply{
		Message: "this WORKED!",
	}
	return &resp, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := generate(ctx)
	gSrv := greeterServer{
		intCh: ch,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	grpc := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(grpc, &gSrv)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		logrus.Infof("[server] got signal %v, attempting graceful stop", s)
		cancel()
		grpc.GracefulStop()
		// grpc.Stop() // leads to error while receiving stream response: rpc error: code = Unavailable desc = transport is closing
		wg.Done()
	}()

	logrus.Infoln("starting grpc server")
	err = grpc.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
	wg.Wait()
	logrus.Infoln("clean shutdown")
}
