package api

import (
	"database/sql"
	"gin/main/massage-board/model"
	"gin/main/massage-board/service"
	"gin/main/massage-board/util"
	"github.com/gin-gonic/gin"
	"log"
)

func Sendmassage(c *gin.Context) {
	uid := c.PostForm("uid")
	mid := c.PostForm("mid")
	massage := c.PostForm("massage")
	err := service.Creatmassage(model.Massage{
		Uid:      uid,
		Mid:      mid,
		Lmassage: massage,
	})
	if uid == "" || mid == "" || massage == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		util.RestpInternalErr(c)
		return
	}
	util.Respok(c)
}
func ChangeMassage(c *gin.Context) {
	uid := c.PostForm("uid")
	mid := c.PostForm("mid")
	massage := c.PostForm("massage")
	err, m := service.SearchmassageByMID(mid)
	if mid == "" || uid == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "信息不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if m.Uid != uid {
		util.Normalerr(c, 6, "没有权限的用户")
		return
	}
	if m.Lmassage == massage {
		util.Normalerr(c, 5, "留言修改重复")
		return
	}
	service.ChangeuserMassageByMID(mid, massage)
	util.Respok(c)
}
func DeleteMassage(c *gin.Context) {
	uid := c.PostForm("uid")
	mid := c.PostForm("mid")
	massage := c.PostForm("留言已删除")
	err, m := service.SearchmassageByMID(mid)
	if mid == "" || uid == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "信息不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if m.Uid != uid {
		util.Normalerr(c, 6, "没有权限的用户")
		return
	}
	service.ChangeuserMassageByMID(mid, massage)
	util.Respok(c)
}
