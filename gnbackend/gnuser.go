package gnbackend

import (
	"crypto/sha512"
	"encoding/base64"
	"time"
)

type User struct {
	Id                 int64
	Modified           time.Time
	Uname, Rname, Mail string
	Pwd                string
	IsLogged           bool
}

func NewUser() *User {
	return &User{
		Id:       -1,
		Modified: time.Now().Local(),
		IsLogged: false,
	}
}

func RegisterUser(uname, rname, mail, pwd string) (*User, error) {
	user := NewUser()
	user.Uname = uname
	user.Rname = rname
	user.Mail = mail
	user.Pwd = encPwd(pwd)
	return DbUser.InsertUser(user)
}

func GetUserByName(uname string) (*User, error) {
	return DbUser.GetUserByField("uname", uname)
}
func GetUserByCookieValue(sessid string) (*User, error) {
	uid, err := DbSess.SelectUserIdFromSessionCookie(sessid)
	if err != nil {
		return NewUser(), err
	}
	return DbUser.GetUserByField("id", uid)
}
func GetUserById(uid int64) (*User, error) {
	return DbUser.GetUserByField("id", uid)
}

func (u *User) VerifyPwd(pwd string) bool {
	if encPwd(pwd) == u.Pwd {
		return true
	}
	Logg.Printf("LOGIN FAILED for USER %s with PASSWORD %s", u.Uname, pwd)
	DbLog.WriteLog(Logbuf, "ERROR")
	return false
}

func (u *User) UpdateLogin(islogged bool) error {
	return DbUser.UpdateFieldById(u.Id, "islogged", islogged)
}
func encPwd(input string) string {
	h := sha512.New()
	h.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.Sum([]byte{}))
}
