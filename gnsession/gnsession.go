package gnsession

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	b "gonotify/gnbackend"
	"net/http"
	"time"
)

type Session struct {
	Id, Uid int64
	Scookie *http.Cookie
	Expire  time.Time
}

func New() *Session {
	return &Session{
		Uid: -1,
		Scookie: &http.Cookie{
			Name:   "sessionid",
			Value:  "",
			MaxAge: 3600,
		},
		Expire: time.Now().Local().Add(1 * time.Hour),
	}
}

func (s *Session) CreateSession(u *b.User) (err error) {
	s.Scookie.Value = createCookieValue(s.Expire, u.Uname)
	id, err := b.StartSession(u.Id, s.Scookie.Value, s.Expire)
	if id > 0 {
		s.Id = id
	}
	return
}

func createCookieValue(dest time.Time, val string) string {
	key := string(dest.UnixNano())
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(val))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func VerifyCookie(c *http.Cookie) bool {
	if err := b.FindSessionByCookie(c.Value); err != nil {
		return false
	}
	return true
}

func Logout(c *http.Cookie) error {
	c.MaxAge = 0
	return b.StopSession(c.Value)
}
