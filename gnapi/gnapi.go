package gnapi

import (
	"encoding/json"
	back "gonotify/gnbackend"
	"net/http"
)

func PostItem(r *http.Request) (int64, error) {
	decoder := json.NewDecoder(r.Body)
	i := back.NewItem()
	err := decoder.Decode(&i.Mex)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	id, err := back.PostItem(i)
	i.Id = id
	return id, err
}

func DeleteItem(id int64) error {
	err := back.DeleteItem(id)
	return err
}
