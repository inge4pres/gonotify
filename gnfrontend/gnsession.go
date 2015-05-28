package gnfrontend

import (
	"crypto/sha512"
	"encoding/base64"
	b "gonotify/gnbackend"
	"math/rand"
	"net/http"
	"time"
)

type Session struct {
	Id, Uid int64
	Scookie *http.Cookie
	Expire  time.Time
}

func NewSession() *Session {
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
	id, err := b.DbSess.InsertSession(u.Id, s.Scookie.Value, s.Expire)
	if err != nil {
		b.Logg.Printf("Session creation failed for user with ID %d", s.Scookie.Value, u.Id)
		b.DbLog.WriteLog(b.Logbuf, "ERROR")
	}
	if id > 0 {
		s.Id = id
	}
	return
}
func createCookieValue(dest time.Time, val string) string {
	rand.Seed(time.Now().Local().UnixNano())
	key := string(rand.Int63n(dest.UnixNano()))
	h := sha512.New()
	h.Write([]byte(val))
	h.Write([]byte(key))
	return base64.StdEncoding.EncodeToString(h.Sum([]byte(nil)))
}
func VerifyCookie(c *http.Cookie) bool {
	uid, err := b.DbSess.SelectUserIdFromSessionCookie(c.Value)
	if err != nil || uid < 1 {
		return false
	}
	return true
}
func Logout(c *http.Cookie) error {
	return b.StopSession(c.Value)
}
