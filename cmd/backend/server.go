package main

import (
	"context"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Emircaan/reverse-proxy/pkg/proto"
	"google.golang.org/grpc"
)

type BackendServer struct {
	proto.UnimplementedProxyServiceServer
	port string
}

func (s *BackendServer) ForwardRequest(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Printf("Backend %s İstek Alindi:", s.port)
	processedMessage := strings.ToUpper(req.Message)
	log.Println("Backend: Gelen Mesaj İslendi: ", processedMessage)
	return &proto.Response{Result: processedMessage}, nil
}

func main() {
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "50052"
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Backend %s sunucusu başlatılamadı: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterProxyServiceServer(grpcServer, &BackendServer{port: port})

	log.Println("Backend server çalışıyor: localhost:50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Backend %s sunucusu çalıştırılamadı: %v", port, err)
	}
}
