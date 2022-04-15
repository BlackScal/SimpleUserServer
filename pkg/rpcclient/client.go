package rpcclient

import (
	pbuserv1 "userserver/api/proto/user/v1"

	"google.golang.org/grpc"
)

func NewUserV1Client(conn grpc.ClientConnInterface) pbuserv1.UserServiceClient {
	return pbuserv1.NewUserServiceClient(conn)
}
