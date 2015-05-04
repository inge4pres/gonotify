package gnbackend

import (
	"bytes"
	"log"
	"time"
)

var dblog *DbParam
var dbitem *DbParam
var l *log.Logger
var lbuf bytes.Buffer

type JReq struct {
	Level   string `json:"level"`
	Rcpnt   string `json:"recipient"`
	Sndr    string `json:"sender"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type Item struct {
	Id       int64
	Time     time.Time
	Mex      JReq
	Archived bool
}

func init() {
	l = log.New(&lbuf, "", log.Lshortfile)
	dblog = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_log")
	dbitem = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_item")
}

func NewItem() Item {
	return Item{
		Time:     time.Now().Local(),
		Archived: false,
	}
}

func GetItem(id int64) (Item, error) {
	item, err := dbitem.GetItem(id)
	if err != nil {
		l.Printf("GET failed for ID %d", id)
		dblog.WriteLog(lbuf, "ERROR")
	}
	return item, err
}

func PostItem(item Item) (int64, error) {
	id, err := dbitem.PostItem(item)
	if err == nil {
		l.Printf("Adding new item from %s to %s", item.Mex.Sndr, item.Mex.Rcpnt)
		err = dblog.WriteLog(lbuf, "INFO")
	} else {
		l.Printf("Not adding new item from %s to %s because of %s", item.Mex.Sndr, item.Mex.Rcpnt, err.Error())
		_ = dblog.WriteLog(lbuf, "ERROR")
	}
	return id, err
}

func DeleteItem(id int64) error {
	err := dbitem.DeleteById(id)
	if err != nil {
		l.Printf("NOT Deleting ITEM with ID %d because of %s", id, err.Error())
		_ = dblog.WriteLog(lbuf, "INFO")
		return err
	}
	l.Printf("Deleting ITEM with ID %d", id)
	_ = dblog.WriteLog(lbuf, "INFO")
	return err
}
