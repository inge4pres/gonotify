package gnbackend

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB
var DbLog *DbParam
var DbItem *DbParam
var DbUser *DbParam
var DbSess *DbParam

type DbParam struct {
	user, pass, url, name, table string
	params                       []string
}

func newConn(DbUser, dbpass, dbaddr, dbname, dbtable string, params []string) *DbParam {
	return &DbParam{
		user:   DbUser,
		pass:   dbpass,
		url:    dbaddr,
		name:   dbname,
		table:  dbtable,
		params: params,
	}
}

func (l *DbParam) WriteLog(mex bytes.Buffer, level string) error {
	db, err := openConn(l)
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+l.table+
		" VALUES (NULL, NULL, ?, ?)",
		level, mex.String())
	if err != nil {
		return err
	}
	return nil
}
func (u *DbParam) InsertUser(us *User) (*User, error) {
	db, err := openConn(u)
	defer db.Close()
	res, err := db.Exec("INSERT INTO "+u.table+" VALUES (null,?,?,?,?,?,?)", us.Modified, us.Uname, us.Rname, us.Mail, us.Pwd, false)
	if err != nil {
		return us, err
	}
	id, err := res.LastInsertId()
	us.Id = id
	return us, err
}
func (i *DbParam) GetItem(id int64) (Item, error) {
	db, err := openConn(i)
	defer db.Close()
	item := NewItem()
	err = db.QueryRow("SELECT * from "+i.table+" WHERE id = ?", id).Scan(&item.Id, &item.Time, &item.Notify.Level, &item.Notify.Rcpnt, &item.Notify.Sndr, &item.Notify.Subject, &item.Notify.Message, &item.Archived)
	return item, err
}
func (i *DbParam) InsertItem(item Item) (int64, error) {
	db, err := openConn(i)
	defer db.Close()
	res, err := db.Exec("INSERT INTO "+i.table+" VALUES (null , ?, ?, ?, ?, ?, ?, ?)",
		item.Time, item.Notify.Level, item.Notify.Rcpnt, item.Notify.Sndr, item.Notify.Subject, item.Notify.Message, item.Archived)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
func (o *DbParam) DeleteById(id int64) error {
	db, err := openConn(o)
	defer db.Close()
	_, err = db.Exec("DELETE FROM "+o.table+" WHERE id = ?", id)
	return err
}
func (u *DbParam) GetUserByField(field string, value interface{}) (*User, error) {
	db, _ := openConn(u)
	defer db.Close()
	user := NewUser()
	err := db.QueryRow("SELECT * from "+u.table+" WHERE "+field+" = ?", value).Scan(&user.Id, &user.Modified, &user.Uname, &user.Rname, &user.Mail, &user.Pwd, &user.IsLogged)
	return user, err
}
func (i *DbParam) GetUserItems(user *User) ([]Item, error) {
	db, err := openConn(i)
	defer db.Close()
	var items []Item
	res, err := db.Query("SELECT * FROM "+i.table+" WHERE recipient = ?", user.Mail)
	defer res.Close()
	for res.Next() {
		item := NewItem()
		if err := res.Scan(&item.Id, &item.Time, &item.Notify.Level, &item.Notify.Rcpnt, &item.Notify.Sndr,
			&item.Notify.Subject, &item.Notify.Message, &item.Archived); err != nil {
			return items, err
		}
		items = append(items, item)
	}
	return items, err
}
func (u *DbParam) UpdateFieldById(id int64, field string, value interface{}) error {
	db, err := openConn(u)
	defer db.Close()
	res, err := db.Exec("UPDATE "+u.table+" SET "+field+" = ? WHERE id = ?", value, id)
	txid, _ := res.LastInsertId()
	Logg.Printf("Updating %s for user with ID %d, setting on %b; TXID %d", field, id, value, txid)
	return err
}
func (s *DbParam) InsertSession(uid int64, scookie string, expires time.Time) (int64, error) {
	db, err := openConn(s)
	defer db.Close()
	res, err := db.Exec("INSERT INTO "+s.table+" VALUES( null, null, ?, ?, ?)", expires, uid, scookie)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
func (s *DbParam) SelectUserIdFromSessionCookie(value string) (int64, error) {
	db, _ := openConn(s)
	defer db.Close()
	var id int64
	err := db.QueryRow("SELECT id FROM gn_session where scookie = ?", value).Scan(&id)
	return id, err
}

func openConn(db *DbParam) (*sql.DB, error) {
	var p string
	if db.params != nil {
		p += "?"
		for s := range db.params {
			p += db.params[s]
		}
	}
	return sql.Open("mysql", db.user+":"+db.pass+"@"+db.url+"/"+db.name+p)
}
