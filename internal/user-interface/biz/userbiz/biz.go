package userbiz

import (
	"context"
	pbuserv1 "userserver/api/proto/user/v1"
	"userserver/internal/pkg/log"
	"userserver/internal/user-interface/data/tokendat"
	"userserver/pkg/rpcclient"

	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type User struct {
	UserId   string
	UserName *string
	Desc     *string
}

type UserBizInterface interface {
	Init() error
	Login(c context.Context, username string, passwd string) (string, string, error)
	Logout(c context.Context, userid string, token string) error
	GetUserInfo(c context.Context, userid string) (*User, error)
	SetUserInfo(c context.Context, userinfo *User) error
	AddUserInfo(c context.Context, userinfo *User) (string, error)
}

type UserBiz struct {
	grpcConn grpc.ClientConnInterface
	tokendat tokendat.TokenDatInterface
}

func NewUserBiz(
	grpcConn grpc.ClientConnInterface,
	tokendat tokendat.TokenDatInterface) *UserBiz {
	return &UserBiz{
		grpcConn: grpcConn,
		tokendat: tokendat}
}

func (u *UserBiz) Init() error {
	if u.grpcConn == nil {
		return errors.New("grpc conn is nil.")
	}
	if u.tokendat == nil {
		return errors.New("tokendat is nil")
	}
	if err := u.tokendat.Init(); err != nil {
		return errors.WithMessage(err, "tokendat init failed")
	}

	return nil
}

func (u *UserBiz) Login(c context.Context, username string, passwd string) (string, string, error) {
	rpcRequest := pbuserv1.LoginRequest{}
	rpcRequest.Username = username
	rpcRequest.Passwd = passwd

	md := metadata.Pairs("rid", c.Value("rid").(string))
	ctx := metadata.NewOutgoingContext(c, md)

	grpcclient := rpcclient.NewUserV1Client(u.grpcConn)
	rpcReply, err := grpcclient.Login(ctx, &rpcRequest)
	if err != nil {
		return "", "", err
	}

	userid, token, auth := rpcReply.Userid, rpcReply.Token, rpcReply.Auth
	if err := u.tokendat.SetToken(token, userid, auth); err != nil {
		log.Warn(log.LogFields{}, err.Error())
	}
	return userid, token, nil

}

func (u *UserBiz) Logout(c context.Context, userid string, token string) error {
	cacheuserid, _, err := u.tokendat.GetToken(token)
	if err == nil && cacheuserid != userid { //TODO
		return errors.New("Invalid Operation.")
	}

	rpcRequest := pbuserv1.LogoutRequest{}
	rpcRequest.Userid = userid
	rpcRequest.Token = token

	md := metadata.Pairs("rid", c.Value("rid").(string))
	ctx := metadata.NewOutgoingContext(c, md)
	grpcclient := rpcclient.NewUserV1Client(u.grpcConn)
	if _, err := grpcclient.Logout(ctx, &rpcRequest); err != nil {
		return err
	}

	if err := u.tokendat.DelToken(token); err != nil {
		log.Warn(log.LogFields{}, err.Error())
	}
	return nil
}

func (u *UserBiz) GetUserInfo(c context.Context, userid string) (*User, error) {
	//TODO: auth check
	rpcRequest := pbuserv1.GetUserRequest{}
	rpcRequest.Userid = userid

	md := metadata.Pairs("rid", c.Value("rid").(string))
	ctx := metadata.NewOutgoingContext(c, md)

	grpcclient := rpcclient.NewUserV1Client(u.grpcConn)
	user, err := grpcclient.GetUser(ctx, &rpcRequest)
	if err != nil {
		return nil, err
	}

	userinfo := &User{
		UserId:   user.Userid, //TODO: set to ""
		UserName: &user.Username,
		Desc:     &user.Desc,
	}
	return userinfo, nil
}

func (u *UserBiz) SetUserInfo(c context.Context, userinfo *User) error {
	rpcRequest := pbuserv1.SetUserRequest{Userid: userinfo.UserId}
	if userinfo.UserName != nil {
		rpcRequest.Username = &pbuserv1.StringNil{Data: *userinfo.UserName}
	}
	if userinfo.Desc != nil {
		rpcRequest.Desc = &pbuserv1.StringNil{Data: *userinfo.Desc}
	}

	md := metadata.Pairs("rid", c.Value("rid").(string))
	ctx := metadata.NewOutgoingContext(c, md)

	grpcclient := rpcclient.NewUserV1Client(u.grpcConn)
	_, err := grpcclient.SetUser(ctx, &rpcRequest)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserBiz) AddUserInfo(c context.Context, userinfo *User) (string, error) {
	rpcRequest := pbuserv1.AddUserRequest{}
	if userinfo.UserName != nil {
		rpcRequest.Username = &pbuserv1.StringNil{Data: *userinfo.UserName}
	}
	if userinfo.Desc != nil {
		rpcRequest.Desc = &pbuserv1.StringNil{Data: *userinfo.Desc}
	}

	md := metadata.Pairs("rid", c.Value("rid").(string))
	ctx := metadata.NewOutgoingContext(c, md)

	grpcclient := rpcclient.NewUserV1Client(u.grpcConn)
	reply, err := grpcclient.AddUser(ctx, &rpcRequest)
	if err != nil {
		return "", err
	}

	return reply.Userid, nil
}
