package service

import (
	"context"
	healthv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/health/v1"
	hellov1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/hello/v1"
	pingv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/ping/v1"
	streamv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/stream/v1"
	tryv1 "github.com/cao7113/hellogolang/proto/gosdk/grpc/try/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

func StartRPCServer(port int, host string) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	address := strings.Join([]string{host, strconv.Itoa(port)}, ":")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("failed to listen address: %s with error: %v", address, err)
	}
	defer lis.Close()

	var opts []grpc.ServerOption
	opts = setupMiddlewares(opts)
	s := grpc.NewServer(opts...)
	hellov1.RegisterHelloServiceServer(s, &HelloServer{})
	pingv1.RegisterPingServiceServer(s, &PingServer{})
	tryv1.RegisterTryServiceServer(s, &TryServer{})
	streamv1.RegisterStreamServiceServer(s, &StreamServer{})
	healthv1.RegisterHealthServer(s, &HealthServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// graceful-stop https://gist.github.com/embano1/e0bf49d24f1cdd07cffad93097c04f0a
	wg := sync.WaitGroup{}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		sig := <-sigCh
		logrus.Infof("got signal %v, attempting graceful shutdown", sig)
		cancel()
		s.GracefulStop()
		wg.Done()
	}()

	logrus.Infof("[gRPC] running server at %s", address)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("Serve() failed: %v", err)
	}
	wg.Wait()
	logrus.Infof("[gRPC] clean shutdown")
}

func setupMiddlewares(opts []grpc.ServerOption) []grpc.ServerOption {
	streamInterceptors := []grpc.StreamServerInterceptor{
		grpc_ctxtags.StreamServerInterceptor(),
		grpctrace.StreamServerInterceptor(),
	}
	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)))

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		//grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(panicFunc)),
		grpc_ctxtags.UnaryServerInterceptor(),
		grpctrace.UnaryServerInterceptor(),
	}
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)))

	return opts
}

//var panicFunc = func(ctx context.Context, p interface{}) (err error) {
//	wrappedError := fmt.Errorf("panic occurs %v", p)
//
//	//txn := newrelic.FromContext(ctx)
//	//if txn != nil {
//	//	txn.NoticeError(wrappedError)
//	//}
//
//	logrus.Error(wrappedError)
//	return status.Error(codes.Internal, "Internal Error")
//}
