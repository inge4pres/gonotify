package gnbackend

import (
	//	"database/sql"
	"log"
	"time"
)

var dblog *DbParam
var dbitem *DbParam
var l log.Logger

type Item struct {
	Time                   time.Time
	Level                  string
	Rcpnt, Sndr, Rcpt_mail string
	Message                map[string][]rune
}

type Notify interface {
}

func init() {
	l.SetPrefix("BACKEND")
	dblog = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_log")
	dbitem = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_item")
}

func RecvItem(item Item) error {
	err := dbitem.WriteItem(item)
	if err == nil {
		err = dblog.WriteLog("New ITEM", "INFO")
	}
	return err
}
