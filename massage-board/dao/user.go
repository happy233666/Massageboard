package dao

import (
	"gin/main/massage-board/model"
)

func Insertuser(u model.User) (err error) {
	if u.Username == "" || u.Password == "" {
		return nil
	}
	_, err = DB.Exec("insert into user(id,username,password)values (?,?,?)", u.Id, u.Username, u.Password)
	return err
}
func Searchid(id string) (u model.User, err error) {
	row := DB.QueryRow("select id,username,password from user where id = ?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.Username, &u.Password)
	return
}
