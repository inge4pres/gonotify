package gnapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func NewResponse() *Response {
	return &Response{Action: "", Item: back.NewItem(), Status: 0, Err: ""}
}

func ApiHandler(req *http.Request, w http.ResponseWriter) {
	api := mux.NewRouter()
	var resp []byte
	api.HandleFunc("/items/{id:[0-9]+}", func(wr http.ResponseWriter, r *http.Request) {
		resp, _ = GetItem(mux.Vars(req)["id"])
	}).Methods("GET")
	w.Write(resp)
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

func PostItem(req *http.Request) ([]byte, error) {
	r := NewResponse()
	decoder := json.NewDecoder(req.Body)
	i := back.NewItem()
	r.Action = "POST"
	err := decoder.Decode(&i.Notify)
	if err != nil {
		r.Err = err.Error()
		r.Status = http.StatusInternalServerError
		return json.Marshal(r)
	}
	id, err := back.PostItem(i)
	i.Id = id
	r.Item = i
	r.Status = http.StatusCreated
	return json.Marshal(r)
}

func DeleteItem(id int64) ([]byte, error) {
	r := NewResponse()
	err := back.DeleteItem(id)
	r.Action = "DELETE"
	if err != nil {
		r.Err = err.Error()
		r.Status = http.StatusNotFound
		return json.Marshal(r)
	}
	r.Status = http.StatusAccepted
	return json.Marshal(r)
}

func ArchiveItem(id int64) ([]byte, error) {
	r := NewResponse()
	r.Action = "PATCH"
	if err := back.ArchiveItem(id); err != nil {
		r.Err = err.Error()
		r.Status = http.StatusNotFound
		return json.Marshal(r)
	}
	r.Item, _ = back.GetItem(id)
	r.Status = http.StatusAccepted
	return json.Marshal(r)
}

func RenderJson(t interface{}) []byte {
	res, _ := json.Marshal(t)
	return res
}
