package gnbackend

import (
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

func encPwd(input string) []byte {
	return sha512.New().Sum([]byte(input))
}

func (u *User) VerifyPwd(input string) bool {
	if compare(sha512.New().Sum([]byte(input)), u.Pwd) {
		u.IsLogged = true
		if err := u.updateLogin(u.IsLogged); err != nil {
			return false
		}
		return true
	}
	return false
}

func compare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (u *User) updateLogin(logged bool) error {
	return dbitem.UpdateFieldById(u.Id, "islogged", logged)
}

func GetUserByName(uname string) (*User, error) {
	return dbuser.GetUserByField("username", uname)
}
