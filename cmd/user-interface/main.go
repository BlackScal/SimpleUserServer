package main

import (
	"flag"
	"fmt"
	configs "userserver/configs"
	"userserver/internal/pkg/cache/redis"
	log "userserver/internal/pkg/log"
	userbiz "userserver/internal/user-interface/biz/userbiz"
	"userserver/internal/user-interface/data/tokendat"
	service "userserver/internal/user-interface/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	//cmd params
	debug    bool
	conffile string

	//config
	config configs.UserServerConf

	//server
	server *service.Server
)

func init() {
	flag.Parse()
	debug = *flag.Bool("debug", true, "set debug mode, true | false") //TODO: Set default value to false
	conffile = *flag.String("conf", "../../configs/userserver.yaml", "User HTTP Server Config File Path")

	var err error
	if err = configs.Parse(conffile, &config); err != nil {
		panic(err)
	}

	//log
	log.SetFormatter(config.Log.Format)
	log.SetOutput(config.Log.Output)
	log.SetLevel(config.Log.Level)
	log.SetAppid(config.AppID)
}

func main() {
	exitfunc := func(err error) {
		log.Error(log.LogFields{}, err.Error())
		panic(err)
	}

	//cache data, redis
	redisAddr := fmt.Sprintf("%s:%d", config.Redis.IP, config.Redis.Port)
	redis := redis.NewRedis(redisAddr)
	redis.SetDataBase(config.Redis.DB)
	//redis.SetXXX

	tokendat := tokendat.NewTokenDat(redis)

	//user biz, rpc conn
	//TODO:
	rpcServerAddr := fmt.Sprintf("%s:%d", config.UserService.IP, config.UserService.Port)
	rpcConn, err := grpc.Dial(rpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		exitfunc(err)
	}
	userbiz := userbiz.NewUserBiz(rpcConn, tokendat)

	//server
	server = service.NewServer(config.Addr.IP, config.Addr.Port)
	if !debug {
		server.SetMode("release")
	}
	server.SetUserBiz(userbiz)
	if err := server.Init(); err != nil {
		exitfunc(err)
	}

	err = server.Run()
	if err != nil {
		log.Error(log.LogFields{}, err.Error()) //TODO
	}
}
