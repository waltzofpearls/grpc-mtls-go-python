package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/waltzofpearls/grpc-mtls-go-python/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	s := newServer()
	s.run()
}

type server struct {
	tlsCert []byte
	tlsKey  []byte
	rootCA  []byte
	address string
}

func newServer() *server {
	return &server{
		tlsCert: []byte(os.Getenv("TLS_SERVER_CERT")),
		tlsKey:  []byte(os.Getenv("TLS_SERVER_KEY")),
		rootCA:  []byte(os.Getenv("TLS_ROOT_CA")),
		address: os.Getenv("GRPC_SERVER_ADDRESS"),
	}
}

func (s *server) run() error {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("could not listen to %s %v", s.address, err)
	}

	serverOption, err := s.tlsServerOption()
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(serverOption)
	api.RegisterMetricsServer(grpcServer, s)

	log.Println("Starting gRPC server", s.address)
	return grpcServer.Serve(listen)
}

func (s *server) tlsServerOption() (grpc.ServerOption, error) {
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(s.rootCA) {
		return nil, errors.New("failed to append root CA cert")
	}
	certificate, err := tls.X509KeyPair(s.tlsCert, s.tlsKey)
	if err != nil {
		return nil, fmt.Errorf("failed load server TLS key and cert: %s", err)
	}
	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	}

	return grpc.Creds(credentials.NewTLS(tlsConfig)), nil
}

func (s *server) Query(ctx context.Context, req *api.QueryMetricsRequest) (*api.QueryMetricsResponse, error) {
	timestamp, _ := ptypes.TimestampProto(time.Now())
	return &api.QueryMetricsResponse{
		Metrics: []*api.Metric{
			&api.Metric{
				Name: "steps",
				Labels: map[string]string{
					"some":    "label",
					"another": "one",
				},
				Values: []*api.Metric_SamplePair{
					&api.Metric_SamplePair{
						Time:  timestamp,
						Value: 5500,
					},
				},
			},
		},
	}, nil
}
