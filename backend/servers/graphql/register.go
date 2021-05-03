package graphql

import (
	"alfred/logger"
	accPb "alfred/modules/AccountService/v1/pb"
	architectureSuggesterPb "alfred/modules/ArchitectureSuggesterService/pb"
	gitPb "alfred/modules/GitService/v1/pb"
	userProfilePb "alfred/modules/UserProfileService/v1/pb"
	vcsPb "alfred/modules/VCSConnectionService/v1/pb"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterGraphqlModules Todo : Whenever any new modules will be in alfred : it must be registered in below method
/*
  RegisterGraphqlModules: Mapping the services with the single graphql endpoint
*/
func RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := userProfilePb.RegisterUserProfileServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}

	if err := gitPb.RegisterGitServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := architectureSuggesterPb.RegisterArchitectureSuggesterServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := vcsPb.RegisterVCSConnectionServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	if err := accPb.RegisterAccountServiceGraphqlHandler(mux, conn); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
		return err
	}
	return nil
}
