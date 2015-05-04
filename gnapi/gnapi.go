package gnapi

import (
	"encoding/json"
	back "gonotify/gnbackend"
	"net/http"
)

func GetItem(id int64) ([]byte, error) {
	item, err := back.GetItem(id)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(item)
	return resp, err
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
	resp, err := json.Marshal(i)
	return resp, err
}

func DeleteItem(id int64) ([]byte, error) {
	err := back.DeleteItem(id)
	if err != nil {
		return nil, err
	}
	return []byte("{\"response\":\"OK\"}"), nil
}
