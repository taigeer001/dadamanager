package admin

import (
	"../../../framework"
	"net/http"
)

func GetSessionUser(req *http.Request) *User {
	session, err := xf.Store.Get(req, "g")
	if err != nil {
		return nil
	}
	// 获取数据库链接
	db, err := xf.GetConnect("user.db")
	if err != nil {
		return  nil
	}
	var tmp User
	db.Find(&tmp, session.Values["user-id"])
	if tmp.Id != 0 {
		return &tmp
	}
	return nil
}

func SetSessionUser(user *User, res http.ResponseWriter, req *http.Request) bool {
	if user == nil {
		return false
	}
	session, err := xf.Store.Get(req, "g")
	if err != nil {
		return false
	}
	session.Values["user-id"] = (*user).Id
	return session.Save(req, res) == nil
}

func ClearSessionUser(res http.ResponseWriter, req *http.Request) bool {
	session, err := xf.Store.Get(req, "g")
	if err != nil {
		return false
	}
	delete(session.Values, "user-id")
	return session.Save(req, res) == nil
}

func GetUser(where interface{}) *User {
	// 获取数据库链接
	db, err := xf.GetConnect("user.db")
	if err != nil {
		return nil
	}
	var user User
	db.First(&user, where)
	if user.Id != 0 {
		return &user
	}
	return nil
}