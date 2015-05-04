package gnapi

import (
	"encoding/json"
	back "gonotify/gnbackend"
	"net/http"
	"strconv"
)

func GetItem(id string) ([]byte, int) {
	intid, _ := strconv.Atoi(id)
	item, err := back.GetItem(int64(intid))
	if err != nil {
		return []byte(err.Error()), http.StatusInternalServerError
	}
	resp, _ := json.Marshal(item)
	return resp, http.StatusOK
}

func PostItem(r *http.Request) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)
	i := back.NewItem()
	err := decoder.Decode(&i.Message)
	if err != nil {
		return []byte(err.Error()), err
	}
	id, err := back.PostItem(i)
	i.Id = id
	resp, _ := json.Marshal(i)
	return resp, err
}

func DeleteItem(id int64) ([]byte, error) {
	err := back.DeleteItem(id)
	if err != nil {
		return nil, err
	}
	return []byte("{\"response\":\"OK\"}"), nil
}
