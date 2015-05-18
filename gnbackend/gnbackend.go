package gnbackend

import (
	"bytes"
	"log"
	"time"
)

var dblog *DbParam
var dbitem *DbParam
var dbuser *DbParam
var dbsess *DbParam
var logg *log.Logger
var logbuf bytes.Buffer

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
	logg = log.New(&logbuf, "", log.Lshortfile)
	dblog = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_log", []string{"parseTime=true"})
	dbitem = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_item", []string{"parseTime=true"})
	dbuser = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_user", []string{"parseTime=true"})
	dbsess = NewDbConn("gonotify", "n0tifyM3", "(sviluppo.mtl.it:3306)", "gonotify", "gn_session", []string{"parseTime=true"})
}

func NewItem() Item {
	return Item{
		Id:       -1,
		Time:     time.Now().Local(),
		Archived: false,
	}
}

func GetUserItems(user *User) ([]Item, error) {
	return dbitem.GetUserItems(user)
}
func GetItem(id int64) (item Item, err error) {
	item, err = dbitem.GetItem(id)
	if err != nil {
		log.Printf("GET ITEM with id %d FAILED! Cause : %s", id, err.Error())
		dblog.WriteLog(logbuf, "ERROR")
		return NewItem(), err
	}
	return item, err
}
func PostItem(item Item) (int64, error) {
	id, err := dbitem.InsertItem(item)
	if err != nil {
		log.Printf("Not adding new item from %s to %s because of %s", item.Notify.Sndr, item.Notify.Rcpnt, err.Error())
		_ = dblog.WriteLog(logbuf, "ERROR")
	}
	return id, err
}
func DeleteItem(id int64) error {
	err := dbitem.DeleteById(id)
	if err != nil {
		log.Printf("NOT Deleting ITEM with id %d FAILED! Cause : %s", id, err.Error())
		_ = dblog.WriteLog(logbuf, "ERROR")
		return err
	}
	log.Printf("Deleting ITEM with ID %d", id)
	_ = dblog.WriteLog(logbuf, "INFO")
	return err
}
func ArchiveItem(id int64) (err error) {
	if err = dbitem.UpdateFieldById(id, "archived", true); err != nil {
		logg.Printf("ARCHIVE for ITEM with id %d FAILED! Cause : %s", id, err.Error())
		dblog.WriteLog(logbuf, "ERROR")
	}
	return
}
func StartSession(uid int64, scookie string, expires time.Time) (id int64, err error) {
	if id, err = dbsess.insertSession(uid, scookie, expires); err != nil {
		logg.Printf("Session creation failed for user with ID %d", scookie, uid)
		dblog.WriteLog(logbuf, "ERROR")
	}
	logg.Printf("Session cookie %s created for user with ID %d", scookie, uid)
	dblog.WriteLog(logbuf, "INFO")
	return
}

func GetSession(value string) error {
	return dbsess.searchSession(value)
}
