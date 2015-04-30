package gnbackend

import (
	//	"database/sql"
	"bytes"
	json "encoding/json"
	"log"
	"time"
)

var dblog *DbParam
var dbitem *DbParam
var l *log.Logger
var lbuf bytes.Buffer

type Item struct {
	Id          int64
	Time        time.Time
	Level       string
	Rcpnt, Sndr string
	Message     json.RawMessage
}

type Notify interface {
}

func init() {
	l = log.New(&lbuf, "BACKEND", log.Lshortfile)
	dblog = NewDbConn("gonotify", "n0tifyM3", "", "gonotify", "gn_log")
	dbitem = NewDbConn("gonotify", "n0tifyM3", "", "gonotify", "gn_item")
}

func PostItem(item Item) error {
	err := dbitem.PostItem(item)
	if err == nil {
		l.Printf("Adding new item from %s to %s", item.Sndr, item.Rcpnt)
		err = dblog.WriteLog(lbuf, "INFO")
	} else {
		l.Printf("Not adding new item from %s to %s because of %T", item.Sndr, item.Rcpnt, err.Error())
		_ = dblog.WriteLog(lbuf, "ERROR")
	}
	return err
}
