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

func (l *DbParam) GetItem(id int64) (Item, error) {
	var item Item
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name+"?parseTime=true")
	if err != nil {
		return item, err
	}
	defer db.Close()
	err = db.QueryRow("SELECT * from "+l.table+" WHERE id = ?", id).Scan(&item.Id, &item.Time, &item.Mex.Level, &item.Mex.Rcpnt, &item.Mex.Sndr, &item.Mex.Subject, &item.Mex.Message, &item.Archived)
	return item, err
}

func (l *DbParam) PostItem(item Item) (int64, error) {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO "+l.table+" VALUES (null , ?, ?, ?, ?, ?, ?, ?)",
		item.Time, item.Mex.Level, item.Mex.Rcpnt, item.Mex.Sndr, item.Mex.Subject, item.Mex.Message, item.Archived)
	if err != nil {
		return 0, err
	}
	lid, err := res.LastInsertId()
	return lid, err
}

func (l *DbParam) DeleteById(id int64) error {
	db, err := sql.Open("mysql", l.user+":"+l.pass+"@"+l.url+"/"+l.name)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM "+l.table+" WHERE id = ?", id)
	return err
}
