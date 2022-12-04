package api

import (
	"database/sql"
	"gin/main/massage-board/model"
	"gin/main/massage-board/service"
	"gin/main/massage-board/util"
	"github.com/gin-gonic/gin"
	"log"
)

func SendComment(c *gin.Context) {
	gid := c.PostForm("gid")
	mid := c.PostForm("mid")
	comment := c.PostForm("comment")
	err := service.NewComment(model.Massage{
		Gid:      gid,
		Mid:      mid,
		Comments: comment,
	})
	if gid == "" || mid == "" || comment == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		util.RestpInternalErr(c)
		return
	}
	util.Respok(c)
}
func UpdateComment(c *gin.Context) {
	gid := c.PostForm("gid")
	mid := c.PostForm("mid")
	comment := c.PostForm("comment")
	err, m := service.SearchmassageByMID(mid)
	if mid == "" || gid == "" {
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
	if m.Comments == comment {
		util.Normalerr(c, 5, "留言修改重复")
		return
	}
	service.ChangeuserCommentByGMID(gid, mid, comment)
	util.Respok(c)
}
func DeleteComment(c *gin.Context) {

}
