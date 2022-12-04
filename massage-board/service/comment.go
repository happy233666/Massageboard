package service

import (
	"gin/main/massage-board/dao"
	"gin/main/massage-board/model"
)

func NewComment(m model.Massage) (err error) {
	err = dao.Insertcomment(m)
	return err
}
func ChangeuserCommentByGMID(gid string, mid string, comment string) {

}