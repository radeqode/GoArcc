// Code generated by proroc-gen-graphql, DO NOT EDIT.
package pb

import (
	"context"

	types "alfred/protos/types"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"

	gql_ptypes_timestamp "github.com/ysugimoto/grpc-graphql-gateway/ptypes/timestamp"
)

var (
	gql__type_Repository               *graphql.Object      // message Repository in git-service.proto
	gql__type_ListRepositoryResponse   *graphql.Object      // message ListRepositoryResponse in git-service.proto
	gql__type_ListRepositoryRequest    *graphql.Object      // message ListRepositoryRequest in git-service.proto
	gql__type_GetRepositoryRequest     *graphql.Object      // message GetRepositoryRequest in git-service.proto
	gql__type_CloneRepositoryResponse  *graphql.Object      // message CloneRepositoryResponse in git-service.proto
	gql__type_CloneRepositoryRequest   *graphql.Object      // message CloneRepositoryRequest in git-service.proto
	gql__input_Repository              *graphql.InputObject // message Repository in git-service.proto
	gql__input_ListRepositoryResponse  *graphql.InputObject // message ListRepositoryResponse in git-service.proto
	gql__input_ListRepositoryRequest   *graphql.InputObject // message ListRepositoryRequest in git-service.proto
	gql__input_GetRepositoryRequest    *graphql.InputObject // message GetRepositoryRequest in git-service.proto
	gql__input_CloneRepositoryResponse *graphql.InputObject // message CloneRepositoryResponse in git-service.proto
	gql__input_CloneRepositoryRequest  *graphql.InputObject // message CloneRepositoryRequest in git-service.proto
)

func Gql__type_Repository() *graphql.Object {
	if gql__type_Repository == nil {
		gql__type_Repository = graphql.NewObject(graphql.ObjectConfig{
			Name:        "Pb_Type_Repository",
			Description: `todo: rename git url to generic name`,
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"node_id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"full_name": &graphql.Field{
					Type: graphql.String,
				},
				"default_branch": &graphql.Field{
					Type: graphql.String,
				},
				"master_branch": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
				"pushed_at": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
				"updated_at": &graphql.Field{
					Type: gql_ptypes_timestamp.Gql__type_Timestamp(),
				},
				"clone_url": &graphql.Field{
					Type: graphql.String,
				},
				"git_url": &graphql.Field{
					Type: graphql.String,
				},
				"size": &graphql.Field{
					Type: graphql.Int,
				},
				"private": &graphql.Field{
					Type: graphql.Boolean,
				},
				"branches": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__type_Repository
}

func Gql__type_ListRepositoryResponse() *graphql.Object {
	if gql__type_ListRepositoryResponse == nil {
		gql__type_ListRepositoryResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_ListRepositoryResponse",
			Fields: graphql.Fields{
				"repositories": &graphql.Field{
					Type: graphql.NewList(Gql__type_Repository()),
				},
			},
		})
	}
	return gql__type_ListRepositoryResponse
}

func Gql__type_ListRepositoryRequest() *graphql.Object {
	if gql__type_ListRepositoryRequest == nil {
		gql__type_ListRepositoryRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_ListRepositoryRequest",
			Fields: graphql.Fields{
				"provider": &graphql.Field{
					Type: types.Gql__enum_GitProviders(),
				},
				"user_id": &graphql.Field{
					Type: graphql.String,
				},
				"account_id": &graphql.Field{
					Type: graphql.String,
				},
				"label": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_ListRepositoryRequest
}

func Gql__type_GetRepositoryRequest() *graphql.Object {
	if gql__type_GetRepositoryRequest == nil {
		gql__type_GetRepositoryRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_GetRepositoryRequest",
			Fields: graphql.Fields{
				"provider": &graphql.Field{
					Type: types.Gql__enum_GitProviders(),
				},
				"repo_name": &graphql.Field{
					Type: graphql.String,
				},
				"account_id": &graphql.Field{
					Type: graphql.String,
				},
				"owner_name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetRepositoryRequest
}

func Gql__type_CloneRepositoryResponse() *graphql.Object {
	if gql__type_CloneRepositoryResponse == nil {
		gql__type_CloneRepositoryResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_CloneRepositoryResponse",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"logs": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_CloneRepositoryResponse
}

func Gql__type_CloneRepositoryRequest() *graphql.Object {
	if gql__type_CloneRepositoryRequest == nil {
		gql__type_CloneRepositoryRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Pb_Type_CloneRepositoryRequest",
			Fields: graphql.Fields{
				"repo": &graphql.Field{
					Type: Gql__type_Repository(),
				},
			},
		})
	}
	return gql__type_CloneRepositoryRequest
}

func Gql__input_Repository() *graphql.InputObject {
	if gql__input_Repository == nil {
		gql__input_Repository = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_Repository",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"node_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"full_name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"default_branch": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"master_branch": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"created_at": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
				"pushed_at": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
				"updated_at": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamp.Gql__input_Timestamp(),
				},
				"clone_url": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"git_url": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"size": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"private": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"branches": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
		})
	}
	return gql__input_Repository
}

func Gql__input_ListRepositoryResponse() *graphql.InputObject {
	if gql__input_ListRepositoryResponse == nil {
		gql__input_ListRepositoryResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_ListRepositoryResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"repositories": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_Repository()),
				},
			},
		})
	}
	return gql__input_ListRepositoryResponse
}

func Gql__input_ListRepositoryRequest() *graphql.InputObject {
	if gql__input_ListRepositoryRequest == nil {
		gql__input_ListRepositoryRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_ListRepositoryRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"provider": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"user_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"account_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"label": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_ListRepositoryRequest
}

func Gql__input_GetRepositoryRequest() *graphql.InputObject {
	if gql__input_GetRepositoryRequest == nil {
		gql__input_GetRepositoryRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_GetRepositoryRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"provider": &graphql.InputObjectFieldConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"repo_name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"account_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"owner_name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetRepositoryRequest
}

func Gql__input_CloneRepositoryResponse() *graphql.InputObject {
	if gql__input_CloneRepositoryResponse == nil {
		gql__input_CloneRepositoryResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_CloneRepositoryResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"logs": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_CloneRepositoryResponse
}

func Gql__input_CloneRepositoryRequest() *graphql.InputObject {
	if gql__input_CloneRepositoryRequest == nil {
		gql__input_CloneRepositoryRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Pb_Input_CloneRepositoryRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"repo": &graphql.InputObjectFieldConfig{
					Type: Gql__input_Repository(),
				},
			},
		})
	}
	return gql__input_CloneRepositoryRequest
}

// graphql__resolver_GitService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_GitService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_GitService creates pointer of service struct
func new_graphql_resolver_GitService(conn *grpc.ClientConn) *graphql__resolver_GitService {
	return &graphql__resolver_GitService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_GitService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_GitService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"repositories": &graphql.Field{
			Type: Gql__type_ListRepositoryResponse(),
			Args: graphql.FieldConfigArgument{
				"provider": &graphql.ArgumentConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"account_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"label": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ListRepositoryRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for repositories")
				}
				client := NewGitServiceClient(conn)
				resp, err := client.ListRepository(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC listRepository")
				}
				return resp, nil
			},
		},
		"repository": &graphql.Field{
			Type: Gql__type_Repository(),
			Args: graphql.FieldConfigArgument{
				"provider": &graphql.ArgumentConfig{
					Type: types.Gql__enum_GitProviders(),
				},
				"repo_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"account_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"owner_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetRepositoryRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for repository")
				}
				client := NewGitServiceClient(conn)
				resp, err := client.GetRepository(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetRepository")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_GitService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"cloneRepository": &graphql.Field{
			Type: Gql__type_CloneRepositoryResponse(),
			Args: graphql.FieldConfigArgument{
				"repo": &graphql.ArgumentConfig{
					Type: Gql__input_Repository(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req CloneRepositoryRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for cloneRepository")
				}
				client := NewGitServiceClient(conn)
				resp, err := client.CloneRepository(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC CloneRepository")
				}
				return resp, nil
			},
		},
	}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterGitServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterGitServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterGitServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service GitService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterGitServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_GitService(conn))
}
