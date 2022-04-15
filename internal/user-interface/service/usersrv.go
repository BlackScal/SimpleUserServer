package service

import (
	"context"
	"fmt"
	"userserver/internal/user-interface/biz/userbiz"
	"userserver/pkg/utils"

	"github.com/pkg/errors"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	ip      string
	port    int
	router  *gin.Engine
	userbiz userbiz.UserBizInterface
}

func NewServer(ip string, port int) *Server {
	router := gin.Default()
	return &Server{router: router, ip: ip, port: port}
}

func (s *Server) SetMode(mode string) {
	gin.SetMode(mode)
}

func (s *Server) SetUserBiz(userbiz userbiz.UserBizInterface) {
	s.userbiz = userbiz
}

func (s *Server) Init() error {
	if s.userbiz == nil {
		return errors.New("Init failed: userbiz is empty.")
	}
	if err := s.userbiz.Init(); err != nil {
		return errors.WithMessage(err, "userbiz init failed.")
	}

	//TODO: Pass *Server in a better way
	s.router.Use(func(c *gin.Context) {
		c.Set("s", s)
		c.Next()
	})

	s.router.Use(ginHookSetRequestID())

	s.router.GET("/ping", getPing)
	v1 := s.router.Group("/v1")
	{
		//TODO: invalid path handler
		v1.POST("/login", postLogin)
		v1.POST("/logout", postLogout)
		v1.GET("/userinfo", getUserInfo)
		v1.POST("/userinfo", postUserInfo)
		v1.PUT("/userinfo", putUserInfo)
	}

	return nil
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.ip, s.port)
	return s.router.Run(addr) //TODO: HTTPS
}

func ginHookSetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := utils.NewUUID()
		ctx := context.WithValue(c.Request.Context(), "rid", uuid)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
