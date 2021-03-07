package middleware

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddLogging returns grpc.Server config option that turn on logging.
func AddInterceptors(logger *zap.Logger, opts []grpc.ServerOption) []grpc.ServerOption {
	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	//grpc recovery options
	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(grpcPanicsRecovery),
	}
	// Add unary interceptor
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(logger),
		//Adding prothesis monitoring
		grpc_prometheus.UnaryServerInterceptor,
		//validate the incoming request - inbound in proto file
		//If request is not correct the error will be sent to client
		grpc_validator.UnaryServerInterceptor(),
		//turns grpc panics into unknown error
		grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
	))

	// Add stream interceptor (added as an example here)
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.StreamServerInterceptor(logger),
		grpc_prometheus.StreamServerInterceptor,
		//validate the incoming request - inbound in proto file
		//If request is not correct the error will be sent to client
		grpc_validator.StreamServerInterceptor(),
		grpc_recovery.StreamServerInterceptor(recoveryOptions...),
	))

	return opts
}

//Unknown error will be sent to user if any panics will triggered
func grpcPanicsRecovery(in interface{}) error {
	return status.Errorf(codes.Unknown, "panic triggered: %v", in)
}
