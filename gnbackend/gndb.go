package gnbackend

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DbParam struct {
	user, pass, url, name, table string
}

func NewDbConn(dbuser, dbpass, dbaddr, dbname, dbtable string) *DbParam {
	return &DbParam{
		user:  dbuser,
		pass:  dbpass,
		url:   dbaddr,
		name:  dbname,
		table: dbtable,
	}
}

func (l *DbParam) WriteLog(mex bytes.Buffer, level string) error {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
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

func (l *DbParam) PostItem(item Item) error {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+l.table+"VALUES (NULL, ?, ?, ?, ?, ?)",
		item.Time, item.Sndr, item.Level, item.Rcpnt, item.Message)
	if err != nil {
		return err
	}
	return nil
}

func (l *DbParam) DeleteItem(item Item) error {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM "+l.table+" WHERE id = ?", item.Id)
	return err
}
