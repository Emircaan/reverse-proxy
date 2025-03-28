package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/Emircaan/reverse-proxy/pkg/proto"
	"google.golang.org/grpc"
)

type BackendServer struct {
	proto.UnimplementedProxyServiceServer
}

func (s *BackendServer) ForwardRequest(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Println("Backend: İstek Alindi:", req.Message)
	processedMessage := strings.ToUpper(req.Message)
	log.Println("Backend: Gelen Mesaj İslendi: ", processedMessage)
	return &proto.Response{Result: processedMessage}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Backend sunucusu başlatılamadı: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterProxyServiceServer(grpcServer, &BackendServer{})

	log.Println("Backend server çalışıyor: localhost:50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Backend sunucusu çalıştırılamadı: %v", err)
	}
}
