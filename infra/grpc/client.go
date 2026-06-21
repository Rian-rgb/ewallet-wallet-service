package grpc

import (
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"time"
)

type Config struct {
	Target string
}

type ClientConnWrapper struct {
	Conn *grpc.ClientConn
}

func NewClientConn(cfg Config) (*ClientConnWrapper, error) {
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	}

	conn, err := grpc.NewClient(
		cfg.Target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		logger.Error("failed to dial target: ", cfg.Target, "with error: ", err)
		return nil, err
	}

	return &ClientConnWrapper{Conn: conn}, nil
}

func (w *ClientConnWrapper) Close() {
	if w.Conn != nil {
		err := w.Conn.Close()
		if err != nil {
			logger.Error("failed closing grpc connection: ", err)
		}
	}
}
