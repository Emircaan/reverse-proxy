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
	backendClient proto.ProxyServiceClient
}

func (s *Server) ForwardRequest(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println("Proxy: İstek Alındı:" + request.Message)

	resp, err := s.backendClient.ForwardRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("backend hatası: %v", err)
	}

	log.Println("Proxy: Backend'den cevap alındı:", resp.Result)
	return resp, nil

}

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Backend'e bağlanırken hata olustu %v", err)
	}
	defer conn.Close()
	backendClient := proto.NewProxyServiceClient(conn)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Proxy sunucu başlatılamadı: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterProxyServiceServer(grpcServer, &Server{
		backendClient: backendClient,
	})
	log.Println("Proxy server çalışıyor: localhost:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Proxy sunucusu çalıştırılamadı: %v", err)
	}
}
