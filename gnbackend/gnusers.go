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
}

func NewUser() *User {
	return &User{
		Id:       -1,
		Modified: time.Now().Local(),
		Uname:    "",
		Rname:    "",
		Mail:     "",
		Pwd:      nil,
	}
}

func encPwd(input string) []byte {
	return sha512.New().Sum([]byte(input))
}

func (u *User) VerifyPwd(input string) bool {
	return compare(sha512.New().Sum([]byte(input)), u.Pwd)
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

//func Register(name, rname, mail, pwd string) (int64, error) {
//	return -1, nil
//}

func GetUser(uname string) (*User, error) {
	return dbuser.GetUserByField("username", uname)
}
