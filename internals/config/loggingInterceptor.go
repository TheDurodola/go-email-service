package config

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	
	resp, err := handler(ctx, req)

	
	st, _ := status.FromError(err)

	
	slog.Info("gRPC Request",
		"method", info.FullMethod,
		"duration", time.Since(start),
		"code", st.Code().String(),
		"error", err,
	)

	return resp, err
}