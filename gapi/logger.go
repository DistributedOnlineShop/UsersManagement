package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context, req interface{}, in *grpc.UnaryServerInfo, handle grpc.UnaryHandler) (rsp interface{}, err error) {
	startTime := time.Now()
	duration := time.Since(startTime)
	result, err := handle(ctx, req)
	statusCode := codes.Unknown

	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Str("method", in.FullMethod).
		Int("status_code", int(statusCode)).
		Dur("duration", duration).
		Msg("received an grpc request")

	return result, err
}
