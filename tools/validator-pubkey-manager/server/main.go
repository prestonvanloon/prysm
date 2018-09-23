package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	pb "github.com/prysmaticlabs/prysm/proto/validator-pubkey-manager/v1"
)

var (
	port        = flag.Int("port", 10000, "The server port")
	metricsPort = flag.Int("metrics-port", 10001, "Prometheus metrics port")
)

func init() {
	grpc_prometheus.EnableHandlingTimeHistogram()
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		)))

	pb.RegisterPubkeyManagerServer(grpcServer, newServer())

	grpc_prometheus.Register(grpcServer)
	http.Handle("/metrics", prometheus.Handler())

	log.Printf("Running metrics on port %d", *metricsPort)
	go http.ListenAndServe(fmt.Sprintf(":%d", *metricsPort), nil)

	go grpcServer.Serve(lis)

	log.Printf("gRPC server is running on port %d", *port)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	<-sigc

	log.Println("Shutting down gRPC server")
	grpcServer.GracefulStop()
}
