package main

import (
	"flag"
	pbuserv1 "userserver/api/proto/user/v1"
	configs "userserver/configs"
	log "userserver/internal/pkg/log"
	service "userserver/internal/user-service/service"

	"google.golang.org/grpc"
)

var (
	//cmd params
	debug    bool
	conffile string

	//config
	config configs.UserServiceConf

	//server
	server *service.Server
)

func init() {
	flag.Parse()
	debug = *flag.Bool("debug", true, "set debug mode, true | false") //TODO: Set default value to false
	conffile = *flag.String("conf", "../../configs/userservice.yaml", "User RPC Service Config File Path")

	if err := configs.Parse(conffile, &config); err != nil {
		panic(err)
	}

	//log
	log.SetFormatter(config.Log.Format)
	log.SetOutput(config.Log.Output)
	log.SetLevel(config.Log.Level)
	log.SetAppid(config.AppID)

}

func main() {
	//server
	server = service.NewServer(config.Addr.IP, config.Addr.Port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pbuserv1.RegisterUserServiceServer(grpcServer, server)

	server.SetRPCServer(grpcServer)
	err := server.Run()
	if err != nil {
		log.Error(log.LogFields{"server": "userrpcserver"}, err.Error()) //TODO
	}
}
