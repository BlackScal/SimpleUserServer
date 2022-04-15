package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"userserver/api/proto/rpc"
	pbuserv1 "userserver/api/proto/user/v1"
	"userserver/internal/pkg/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pbuserv1.UnimplementedUserServiceServer
	ip         string
	port       int
	grpcServer *grpc.Server
}

func NewServer(ip string, port int) *Server {
	return &Server{ip: ip, port: port}
}

func (s *Server) SetRPCServer(grpcServer *grpc.Server) {
	s.grpcServer = grpcServer
}

func (s *Server) Init() error {
	if s.grpcServer == nil {
		return errors.New("grpc server is empty")
	}
	return nil
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.ip, s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(lis)
}

func getRid(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Warn(log.LogFields{}, "receive a request without rid")
		return ""
	}
	return md["rid"][0]
}
func (s *Server) Login(
	ctx context.Context,
	in *pbuserv1.LoginRequest) (*pbuserv1.LoginReply, error) {

	var reply pbuserv1.LoginReply
	if in.Username == "" || in.Passwd == "" {
		st := status.New(codes.Code(rpc.ERR_INVALID_PARAMS), "Username or Passwd Error.")
		return nil, st.Err()
	}

	reply.Userid = "0001"
	reply.Token = "XXX"
	reply.Auth = 0666
	return &reply, nil
}

func (s *Server) Logout(
	ctx context.Context,
	in *pbuserv1.LogoutRequest) (*pbuserv1.LogoutReply, error) {

	return &pbuserv1.LogoutReply{}, nil
}

func (s *Server) GetUser(
	ctx context.Context,
	in *pbuserv1.GetUserRequest) (*pbuserv1.GetUserReply, error) {
	st := status.New(codes.Code(rpc.ERR_IMPLEMENTED), rpc.ERR_IMPLEMENTED.String())
	return nil, st.Err()
}
