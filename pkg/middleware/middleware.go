package middleware

import (
	"context"

	"github.com/Emircaan/reverse-proxy/pkg/proto"
)

type Middleware interface {
	Handle(ctx context.Context, req *proto.Request) (*proto.Request, error)
}

type MiddlewareChain struct {
	middlewares []Middleware
}

func NewMiddlewareChain(mws ...Middleware) *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: mws,
	}
}

func (mc *MiddlewareChain) Apply(ctx context.Context, req *proto.Request) (*proto.Request, error) {
	var err error
	for _, mw := range mc.middlewares {
		req, err = mw.Handle(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return req, nil
}
