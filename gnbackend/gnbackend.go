package gnbackend

import (
	"bytes"
	"log"
	"time"
)

var dblog *DbParam
var dbitem *DbParam
var dbuser *DbParam
var l *log.Logger
var lbuf bytes.Buffer

type Note struct {
	Level   string `json:"level"`
	Rcpnt   string `json:"recipient"`
	Sndr    string `json:"sender"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type Item struct {
	Id       int64
	Time     time.Time
	Notify   Note
	Archived bool
}

func init() {
	l = log.New(&lbuf, "", log.Lshortfile)
	dblog = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_log", []string{"parseTime=true"})
	dbitem = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_item", []string{"parseTime=true"})
	dbuser = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_user", []string{"parseTime=true"})
}

func NewItem() Item {
	return Item{
		Id:       -1,
		Time:     time.Now().Local(),
		Archived: false,
	}
}

func GetUserItems(user User) ([]Item, error) {
	return dbitem.GetUserItems(user)
}

func GetItem(id int64) (item Item, err error) {
	item, err = dbitem.GetItem(id)
	if err != nil {
		l.Printf("GET failed for ID %d", id)
		dblog.WriteLog(lbuf, "ERROR")
		return NewItem(), err
	}
	return item, err
}

func PostItem(item Item) (int64, error) {
	id, err := dbitem.InsertItem(item)
	if err != nil {
		l.Printf("Not adding new item from %s to %s because of %s", item.Notify.Sndr, item.Notify.Rcpnt, err.Error())
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
