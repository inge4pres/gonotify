package gnbackend

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DbParam struct {
	user, pass, url, name, table string
	params                       []string
}

func NewDbConn(dbuser, dbpass, dbaddr, dbname, dbtable string, params []string) *DbParam {
	return &DbParam{
		user:   dbuser,
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

func (u *DbParam) InsertUser(us *User) error {
	db, err := openConn(u)
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+u.table+" VALUES (null,?,?,?,?,?,?)", us.Modified, us.Uname, us.Rname, us.Mail, us.Pwd, false)
	return err
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
	lid, err := res.LastInsertId()
	return lid, err
}

func (o *DbParam) DeleteById(id int64) error {
	db, err := openConn(o)
	defer db.Close()
	_, err = db.Exec("DELETE FROM "+o.table+" WHERE id = ?", id)
	return err
}

func (u *DbParam) GetUserByField(field, value string) (*User, error) {
	db, err := openConn(u)
	defer db.Close()
	user := NewUser()
	logg.Println("EXECUTING : SELECT * from " + u.table + " WHERE " + field + " = " + value)
	dblog.WriteLog(logbuf, "INFO")
	err = db.QueryRow("SELECT * from "+u.table+" WHERE "+field+" = ?", value).Scan(&user.Id, &user.Modified, &user.Uname, &user.Rname, &user.Mail, &user.Pwd, &user.IsLogged)
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
	_, err = db.Exec("UPDATE "+u.table+" SET "+field+" = ? WHERE id = ?", value, id)
	return err
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
