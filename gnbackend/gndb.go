package gnbackend

import (
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

func (l *DbParam) WriteLog(mex, level string) error {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+l.table+" VALUES (NULL, NULL, ?, ?)", level, mex)
	if err != nil {
		return err
	}
	return nil
}

func (l *DbParam) WriteItem(item Item) error {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO "+l.table+" VALUES (NULL, NULL, ?, ?, NULL, NULL)", item.Level, item.Rcpt_mail)
	if err != nil {
		return err
	}
	return nil
}
