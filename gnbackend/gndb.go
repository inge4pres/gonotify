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
	var p string
	if l.params != nil {
		p += "?"
		for s := range l.params {
			p += l.params[s]
		}
	}
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name+p)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+l.table+
		" VALUES (NULL, NULL, ?, ?)",
		level, mex.String())
	if err != nil {
		return err
	}
	return nil
}

func (l *DbParam) GetItem(id int64) (Item, error) {
	var item Item
	var p string
	if l.params != nil {
		p += "?"
		for s := range l.params {
			p += l.params[s]
		}
	}
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name+p)
	if err != nil {
		return item, err
	}
	defer db.Close()
	err = db.QueryRow("SELECT * from "+l.table+" WHERE id = ?", id).Scan(&item.Id, &item.Time, &item.Notify.Level, &item.Notify.Rcpnt, &item.Notify.Sndr, &item.Notify.Subject, &item.Notify.Message, &item.Archived)
	return item, err
}

func (l *DbParam) InsertItem(item Item) (int64, error) {
	var p string
	if l.params != nil {
		p += "?"
		for s := range l.params {
			p += l.params[s]
		}
	}
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name+p)
	if err != nil {
		return -1, err
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO "+l.table+" VALUES (null , ?, ?, ?, ?, ?, ?, ?)",
		item.Time, item.Notify.Level, item.Notify.Rcpnt, item.Notify.Sndr, item.Notify.Subject, item.Notify.Message, item.Archived)
	if err != nil {
		return -1, err
	}
	lid, err := res.LastInsertId()
	return lid, err
}

func (l *DbParam) DeleteById(id int64) error {
	var p string
	if l.params != nil {
		p += "?"
		for s := range l.params {
			p += l.params[s]
		}
	}
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name+p)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM "+l.table+" WHERE id = ?", id)
	return err
}
