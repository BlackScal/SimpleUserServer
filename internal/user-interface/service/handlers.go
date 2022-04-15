package service

import (
	"net/http"

	"userserver/api/proto/rpc"
	"userserver/internal/pkg/log"
	"userserver/internal/user-interface/biz/userbiz"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

type jsondata map[string]string

func genJsonBody(code codes.Code, msg string, data interface{}) gin.H {
	if data == nil {
		return gin.H{
			"code": code,
			"msg":  msg,
		}
	}
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func _getServer(c *gin.Context) *Server {
	return c.MustGet("s").(*Server)
}

func getPing(c *gin.Context) {
	code := http.StatusOK
	c.JSON(code, genJsonBody(codes.Code(rpc.ERR_OK), "Pong", nil))
}

func postLogin(c *gin.Context) {
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	if username == "" || passwd == "" {
		c.JSON(400, genJsonBody(400, "Invalid Parameters", nil))
		return
	}

	s := _getServer(c)
	userid, token, err := s.userbiz.Login(c.Request.Context(), username, passwd)
	if err != nil {
		log.Error(log.LogFields{"rid": c.Request.Context().Value("rid")}, err.Error())
		c.JSON(500, genJsonBody(500, "", nil))
		return
	}
	c.SetCookie("token", token, 30, "/", s.ip, false, true)
	c.JSON(200, genJsonBody(200, "OK", jsondata{"userid": userid}))
}

func postLogout(c *gin.Context) {
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(400, genJsonBody(400, "Invalid Parameters", nil))
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(400, genJsonBody(400, "Invalid Parameters", nil))
		return
	}

	s := _getServer(c)
	if err := s.userbiz.Logout(c.Request.Context(), userid, token); err != nil {
		log.Error(log.LogFields{"rid": c.Request.Context().Value("rid")}, err.Error())
		c.JSON(500, genJsonBody(500, "", nil))
		return
	}
	c.JSON(200, genJsonBody(200, "OK", nil))
}

func getUserInfo(c *gin.Context) {
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(400, genJsonBody(400, "Invalid Parameters", nil))
		return
	}

	s := _getServer(c)
	userinfo, err := s.userbiz.GetUserInfo(c.Request.Context(), userid)
	if err != nil {
		log.Error(log.LogFields{"rid": c.Request.Context().Value("rid")}, err.Error())
		c.JSON(500, genJsonBody(500, "", nil))
		return
	}
	c.JSON(200, genJsonBody(200, "OK",
		jsondata{
			"userid":   userinfo.UserId,
			"username": *userinfo.UserName,
			"desc":     *userinfo.Desc,
		}))
}

func postUserInfo(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.JSON(400, genJsonBody(400, "Invalid Parameters", nil))
		return
	}
	userinfo := userbiz.User{}
	userinfo.UserName = &username

	if desc, ok := c.GetPostForm("desc"); ok {
		userinfo.Desc = &desc
	}

	s := _getServer(c)
	userid, err := s.userbiz.AddUserInfo(c.Request.Context(), &userinfo)
	if err != nil {
		log.Error(log.LogFields{"rid": c.Request.Context().Value("rid")}, err.Error())
		c.JSON(500, genJsonBody(500, "", nil))
		return
	}
	c.JSON(200, genJsonBody(200, "OK", jsondata{"userid": userid}))
}

func putUserInfo(c *gin.Context) {
	userid := c.PostForm("userid")
	if userid == "" {
		c.JSON(400, genJsonBody(400, "Invalid Parameters", nil))
		return
	}

	userinfo := userbiz.User{UserId: userid}
	if username, ok := c.GetPostForm("username"); ok {
		userinfo.UserName = &username
	}
	if desc, ok := c.GetPostForm("desc"); ok {
		userinfo.Desc = &desc
	}

	s := _getServer(c)
	err := s.userbiz.SetUserInfo(c.Request.Context(), &userinfo)
	if err != nil {
		log.Error(log.LogFields{"rid": c.Request.Context().Value("rid")}, err.Error())
		c.JSON(500, genJsonBody(500, "", nil))
		return
	}
	c.JSON(200, genJsonBody(200, "OK", nil))
}
