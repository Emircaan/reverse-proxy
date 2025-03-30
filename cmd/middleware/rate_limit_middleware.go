package middleware

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Emircaan/reverse-proxy/pkg/proto"
	"google.golang.org/grpc/peer"
)

var ErrTooManyRequest = errors.New("Kısa süre içerisinde çok fazla istek gönderdiniz , lütfen biraz sonra tekrar deneyiniz")

func getIpAdress(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("ip bilgisi alınamadı")
	}
	ip := p.Addr.String()
	return ip, nil
}

type RateLimitMiddleware struct {
	mu       sync.Mutex
	requests map[string]int
	limit    int
	window   time.Duration
}

func (rl *RateLimitMiddleware) resetCounts() {
	ticker := time.NewTicker(rl.window)
	for range ticker.C {
		rl.mu.Lock()
		rl.requests = make(map[string]int)
		rl.mu.Unlock()
	}
}
func NewRateLimitMiddleware(limit int, window time.Duration) *RateLimitMiddleware {
	rl := &RateLimitMiddleware{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
	}
	go rl.resetCounts()
	return rl
}

func (rl *RateLimitMiddleware) Handle(ctx context.Context, req *proto.Request) (*proto.Request, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	clientID, err := getIpAdress(ctx)
	if err != nil {
		return nil, err
	}

	if rl.requests[clientID] >= rl.limit {
		return nil, ErrTooManyRequest
	}

	rl.requests[clientID]++
	return req, nil
}
