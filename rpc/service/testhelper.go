package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
	"strings"
	"time"
)

var defaultLocalRCPPort = 50051

func clientConnWithLocalServer() *grpc.ClientConn {
	return dialClientConnOn(defaultLocalRCPPort, true)
}

//func clientConnWithRemoteSvr() *grpc.ClientConn {
//	port := defaultLocalRCPPort
//	if p, ok := os.LookupEnv("PORT"); ok {
//		var err error
//		port, err = strconv.Atoi(p)
//		if err != nil {
//			logrus.Fatalf("invalid rpc port %v error: %s", p, err.Error())
//		}
//	}
//	return dialClientConnOn(port, false)
//}

func dialClientConnOn(port int, startServer bool) *grpc.ClientConn {
	if startServer {
		if !checkLocalPortUsed(port, false) {
			asyncStartRPC(port)
		}
	}
	addr := strings.Join([]string{"", strconv.Itoa(port)}, ":")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server addr: %s error: %v", addr, err)
	}
	return conn
}

func checkRPCErr(err error) {
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logrus.Errorf("status error code: %v message: %s", st.Code(), st.Message())
		}
	}
}

func checkLocalPortUsed(port int, verbose bool) bool {
	// health-check? todo  https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	addr := fmt.Sprintf("localhost:%d", port)
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		if verbose {
			logrus.Warnf("dial-test error: %s", err.Error())
		}
		return false
	}
	if conn != nil {
		logrus.Warnf("%s already used", addr)
		conn.Close()
		return true
	}
	return false
}

func asyncStartRPC(port int) {
	go StartRPCServer(port, "")
	time.Sleep(50 * time.Millisecond)
}
