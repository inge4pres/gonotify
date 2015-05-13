package gnbackend

import (
	"bytes"
	"crypto/sha512"
	"time"
)

type User struct {
	Id                 int64
	Modified           time.Time
	Uname, Rname, Mail string
	Pwd                []byte
	IsLogged           bool
}

func NewUser() *User {
	return &User{
		Id:       -1,
		Modified: time.Now().Local(),
		Uname:    "",
		Rname:    "",
		Mail:     "",
		Pwd:      nil,
		IsLogged: false,
	}
}

func RegisterUser(uname, rname, mail, pwd string) (*User, error) {
	user := NewUser()
	user.Uname = uname
	user.Rname = rname
	user.Mail = mail
	user.Pwd = encPwd(pwd)
	return dbuser.InsertUser(user)
}

func GetUserByName(uname string) (*User, error) {
	return dbuser.GetUserByField("uname", uname)
}

func (u *User) VerifyPwd(pwd string) bool {
	if bytes.Equal(encPwd(pwd), u.Pwd) {
		u.IsLogged = true
		if err := u.updateLogin(u.IsLogged); err != nil {
			logg.Println("LOGIN FAILED for USER " + u.Uname + " Cause: " + err.Error())
			dblog.WriteLog(logbuf, "ERROR")
			return false
		}
	}
	return u.IsLogged
}

func (u *User) updateLogin(islogged bool) error {
	return dbuser.UpdateFieldById(u.Id, "islogged", islogged)
}

func encPwd(input string) []byte {
	h := sha512.New()
	h.Write([]byte(input))
	return h.Sum([]byte{})
}
