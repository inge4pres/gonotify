package gnapi

import (
	"encoding/json"
	back "gonotify/gnbackend"
	"net/http"
	"strconv"
)

type Response struct {
	Action string    `json:"action"`
	Item   back.Item `json:"item"`
	Status int       `json:"status"`
	Err    string    `json:"error"`
}

func NewResponse() Response {
	return Response{Action: "", Item: back.NewItem(), Status: 0, Err: ""}
}

func GetItem(id string) ([]byte, error) {
	r := NewResponse()
	intid, _ := strconv.Atoi(id)
	item, err := back.GetItem(int64(intid))
	r.Item = item
	r.Action = "GET"
	if err != nil {
		r.Err = err.Error()
		r.Status = http.StatusNotFound
		return json.Marshal(r)
	}
	r.Status = http.StatusOK
	return json.Marshal(r)

}

func PostItem(r *http.Request) ([]byte, error) {
	decoder := json.NewDecoder(r.Body)
	i := back.NewItem()
	err := decoder.Decode(&i.Notify)
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
