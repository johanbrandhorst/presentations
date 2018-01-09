package backend

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/johanbrandhorst/demo/proto/server"
)

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
}

// Ensure struct implements interface
var _ server.BackendServer = (*Backend)(nil)

func (b Backend) GetUser(ctx context.Context, req *server.UserRequest) (*server.User, error) {
	if req.GetId() != "abcd" {
		return nil, status.Error(codes.NotFound, "user could not be found")
	}

	return &server.User{
		Id:   "abcd",
		Name: "Johan",
		Age:  27,
	}, nil
}
