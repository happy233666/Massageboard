package api

import (
	"database/sql"
	"gin/main/massage-board/model"
	"gin/main/massage-board/service"
	"gin/main/massage-board/util"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	id := c.PostForm("id")
	userName := c.PostForm("username")
	password := c.PostForm("password")
	err, u := service.SearchuserByID(id)
	if err != nil && err != sql.ErrNoRows {
		util.RestpInternalErr(c)
		return
	}
	if u.Id != "" {
		util.Normalerr(c, 3, "账号已存在")
	}
	err = service.Creatuser(model.User{
		Id:       id,
		Username: userName,
		Password: password,
	})
	if id == "" || userName == "" || password == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		util.RestpInternalErr(c)
		return
	}
	util.Respok(c)

}
