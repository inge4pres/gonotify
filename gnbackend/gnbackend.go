package gnbackend

import (
	"bytes"
	"log"
	"time"
)

var Logg *log.Logger
var Logbuf bytes.Buffer

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
	Logg = log.New(&Logbuf, "", log.Lshortfile)
	DbLog = newConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_log", []string{"parseTime=true"})
	DbItem = newConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_item", []string{"parseTime=true"})
	DbUser = newConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_user", []string{"parseTime=true"})
	DbSess = newConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_session", []string{"parseTime=true"})
}

func NewItem() Item {
	return Item{
		Id:       -1,
		Time:     time.Now().Local(),
		Archived: false,
	}
}

func GetUserItems(user *User) ([]Item, error) {
	return DbItem.GetUserItems(user)
}
func GetItem(id int64) (item Item, err error) {
	item, err = DbItem.GetItem(id)
	if err != nil {
		log.Printf("GET ITEM with id %d FAILED! Cause : %s", id, err.Error())
		DbLog.WriteLog(Logbuf, "ERROR")
		return NewItem(), err
	}
	return item, err
}
func PostItem(item Item) (int64, error) {
	id, err := DbItem.InsertItem(item)
	if err != nil {
		log.Printf("Not adding new item from %s to %s because of %s", item.Notify.Sndr, item.Notify.Rcpnt, err.Error())
		_ = DbLog.WriteLog(Logbuf, "ERROR")
	}
	return id, err
}
func DeleteItem(id int64) error {
	err := DbItem.DeleteById(id)
	if err != nil {
		log.Printf("NOT Deleting ITEM with id %d FAILED! Cause : %s", id, err.Error())
		_ = DbLog.WriteLog(Logbuf, "ERROR")
		return err
	}
	log.Printf("Deleting ITEM with ID %d", id)
	_ = DbLog.WriteLog(Logbuf, "INFO")
	return err
}
func ArchiveItem(id int64) (err error) {
	if err = DbItem.UpdateFieldById(id, "archived", true); err != nil {
		Logg.Printf("ARCHIVE for ITEM with id %d FAILED! Cause : %s", id, err.Error())
		DbLog.WriteLog(Logbuf, "ERROR")
	}
	return
}
