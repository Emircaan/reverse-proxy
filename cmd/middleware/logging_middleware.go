package middleware

import (
	"context"
	"log"

	"github.com/Emircaan/reverse-proxy/pkg/proto"
)

type LogMiddleware struct{}

func (l *LogMiddleware) Handle(ctx context.Context, req *proto.Request) (*proto.Request, error) {

	log.Printf("Logging Middleware: İstek alındı: %s", req.Message)
	return req, nil
}
