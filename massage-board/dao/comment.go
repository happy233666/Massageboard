package dao

import "gin/main/massage-board/model"

func Insertcomment(m model.Massage) (err error) {
	if m.Mid == "" || m.Gid == "" || m.Comments == "" {
		return nil
	}
	_, err = DB.Exec("insert into comment(gid,mid,comments)values (?,?,?)", m.Gid, m.Mid, m.Comments)
	return err
}
