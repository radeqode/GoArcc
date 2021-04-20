// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	gql_ptypes_timestamp "github.com/ysugimoto/grpc-graphql-gateway/ptypes/timestamp"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__enum_SOURCE                      *graphql.Enum        // enum SOURCE in user-profile-service.proto
	gql__type_UserProfile                 *graphql.Object      // message UserProfile in user-profile-service.proto
	gql__type_UpdateUserProfileRequest    *graphql.Object      // message UpdateUserProfileRequest in user-profile-service.proto
	gql__type_GetUserProfileRequest       *graphql.Object      // message GetUserProfileRequest in user-profile-service.proto
	gql__type_GetUserProfileBySubRequest  *graphql.Object      // message GetUserProfileBySubRequest in user-profile-service.proto
	gql__type_DeleteUserProfileRequest    *graphql.Object      // message DeleteUserProfileRequest in user-profile-service.proto
	gql__type_CreateUserProfileRequest    *graphql.Object      // message CreateUserProfileRequest in user-profile-service.proto
	gql__input_UserProfile                *graphql.InputObject // message UserProfile in user-profile-service.proto
	gql__input_UpdateUserProfileRequest   *graphql.InputObject // message UpdateUserProfileRequest in user-profile-service.proto
	gql__input_GetUserProfileRequest      *graphql.InputObject // message GetUserProfileRequest in user-profile-service.proto
	gql__input_GetUserProfileBySubRequest *graphql.InputObject // message GetUserProfileBySubRequest in user-profile-service.proto
	gql__input_DeleteUserProfileRequest   *graphql.InputObject // message DeleteUserProfileRequest in user-profile-service.proto
	gql__input_CreateUserProfileRequest   *graphql.InputObject // message CreateUserProfileRequest in user-profile-service.proto
)

func Gql__enum_SOURCE() *graphql.Enum {
	if gql__enum_SOURCE == nil {
		gql__enum_SOURCE = graphql.NewEnum(graphql.EnumConfig{
			Name: "Pb_Enum_SOURCE",
			Values: graphql.EnumValueConfigMap{
				"UNKNOWN": &graphql.EnumValueConfig{
					Value: SOURCE(0),
				},
				"GITHUB": &graphql.EnumValueConfig{
					Value: SOURCE(1),
				},
				"GITLAB": &graphql.EnumValueConfig{
					Value: SOURCE(2),
				},
				"BITBUCKET": &graphql.EnumValueConfig{
					Value: SOURCE(3),
				},
				"EMAIL_PASSWORD": &graphql.EnumValueConfig{
					Value: SOURCE(4),
				},
			},
		})
	}
	return gql__enum_SOURCE
}

func Gql__type_UserProfile() *graphql.Object {
	if gql__type_UserProfile == nil {
		gql__type_UserProfile = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_UserProfile",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"sub": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"user_name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"phone_number": &graphql.Field{
					Type: graphql.String,
				},
				"external_source": &graphql.Field{
					Type: Gql__enum_SOURCE(),
				},
				"profile_pic_url": &graphql.Field{
					Type: graphql.String,
				},
				"token_valid_till": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
			},
		})
	}
	return gql__type_UserProfile
}

func Gql__type_UpdateUserProfileRequest() *graphql.Object {
	if gql__type_UpdateUserProfileRequest == nil {
		gql__type_UpdateUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_UpdateUserProfileRequest",
			Fields: graphql.Fields{
				"user_profile": &graphql.Field{
					Type: Gql__type_UserProfile(),
				},
			},
		})
	}
	return gql__type_UpdateUserProfileRequest
}

func Gql__type_GetUserProfileRequest() *graphql.Object {
	if gql__type_GetUserProfileRequest == nil {
		gql__type_GetUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetUserProfileRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetUserProfileRequest
}

func Gql__type_GetUserProfileBySubRequest() *graphql.Object {
	if gql__type_GetUserProfileBySubRequest == nil {
		gql__type_GetUserProfileBySubRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetUserProfileBySubRequest",
			Fields: graphql.Fields{
				"sub": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetUserProfileBySubRequest
}

func Gql__type_DeleteUserProfileRequest() *graphql.Object {
	if gql__type_DeleteUserProfileRequest == nil {
		gql__type_DeleteUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_DeleteUserProfileRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_DeleteUserProfileRequest
}

func Gql__type_CreateUserProfileRequest() *graphql.Object {
	if gql__type_CreateUserProfileRequest == nil {
		gql__type_CreateUserProfileRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_CreateUserProfileRequest",
			Fields: graphql.Fields{
				"user_profile": &graphql.Field{
					Type: Gql__type_UserProfile(),
				},
			},
		})
	}
	return gql__type_CreateUserProfileRequest
}

func Gql__input_UserProfile() *graphql.InputObject {
	if gql__input_UserProfile == nil {
		gql__input_UserProfile = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_UserProfile",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"sub": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"user_name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"email": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"phone_number": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"external_source": &graphql.InputObjectFieldConfig{
					Type: Gql__enum_SOURCE(),
				},
				"profile_pic_url": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"token_valid_till": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
			},
		})
	}
	return gql__input_UserProfile
}

func Gql__input_UpdateUserProfileRequest() *graphql.InputObject {
	if gql__input_UpdateUserProfileRequest == nil {
		gql__input_UpdateUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_UpdateUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_profile": &graphql.InputObjectFieldConfig{
					Type: Gql__input_UserProfile(),
				},
			},
		})
	}
	return gql__input_UpdateUserProfileRequest
}

func Gql__input_GetUserProfileRequest() *graphql.InputObject {
	if gql__input_GetUserProfileRequest == nil {
		gql__input_GetUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetUserProfileRequest
}

func Gql__input_GetUserProfileBySubRequest() *graphql.InputObject {
	if gql__input_GetUserProfileBySubRequest == nil {
		gql__input_GetUserProfileBySubRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetUserProfileBySubRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"sub": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetUserProfileBySubRequest
}

func Gql__input_DeleteUserProfileRequest() *graphql.InputObject {
	if gql__input_DeleteUserProfileRequest == nil {
		gql__input_DeleteUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_DeleteUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_DeleteUserProfileRequest
}

func Gql__input_CreateUserProfileRequest() *graphql.InputObject {
	if gql__input_CreateUserProfileRequest == nil {
		gql__input_CreateUserProfileRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_CreateUserProfileRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_profile": &graphql.InputObjectFieldConfig{
					Type: Gql__input_UserProfile(),
				},
			},
		})
	}
	return gql__input_CreateUserProfileRequest
}

// graphql__resolver_UserProfileService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_UserProfileService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_UserProfileService creates pointer of service struct
func new_graphql_resolver_UserProfileService(conn *grpc.ClientConn) *graphql__resolver_UserProfileService {
	return &graphql__resolver_UserProfileService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_UserProfileService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_UserProfileService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"user": &graphql.Field{
			Type: Gql__type_UserProfile(),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetUserProfileRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for user")
				}
				client := NewUserProfileServiceClient(conn)
				resp, err := client.GetUserProfile(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetUserProfile")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_UserProfileService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterUserProfileServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterUserProfileServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterUserProfileServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service UserProfileService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterUserProfileServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_UserProfileService(conn))
}
