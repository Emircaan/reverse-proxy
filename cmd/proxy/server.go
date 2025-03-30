package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Emircaan/reverse-proxy/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	proto.UnimplementedProxyServiceServer
	backendClients []proto.ProxyServiceClient
	currentIndex   int
}

func (s *Server) ForwardRequest(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("Proxy: İstek Alındı:" + request.Message)
	backendIndex := s.currentIndex
	backendClient := s.backendClients[backendIndex]
	s.currentIndex = (s.currentIndex + 1) % len(s.backendClients)
	resp, err := backendClient.ForwardRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("backend'den cevap gelmedi %v", err)
	}
	log.Println("Proxy: Backend'den cevap alındı:", resp.Result)
	return resp, nil

}

func main() {
	backendAdress := []string{
		"localhost:50052",
		"localhost:50053",
		"localhost:50054",
	}
	var backendClients []proto.ProxyServiceClient
	for _, address := range backendAdress {
		conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Backend sunucusuna bağlanılamadı: %v", err)
		}
		backendClients = append(backendClients, proto.NewProxyServiceClient(conn))
	}
	proxyServer := &Server{
		backendClients: backendClients,
	}
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Proxy sunucusu başlatılamadı: %v", err)
	}
	proto.RegisterProxyServiceServer(grpcServer, proxyServer)
	log.Println("Proxy server başlatıldı...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server başlatılamadı: %v", err)
	}
}
