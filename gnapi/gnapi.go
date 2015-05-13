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

func NewResponse() *Response {
	return &Response{Action: "", Item: back.NewItem(), Status: 0, Err: ""}
}

func GetItem(id string) *Response {
	r := NewResponse()
	intid, _ := strconv.Atoi(id)
	item, err := back.GetItem(int64(intid))
	r.Item = item
	r.Action = "GET"
	if err != nil {
		r.Err = err.Error()
		r.Status = http.StatusNotFound
		return r
	}
	r.Status = http.StatusOK
	return r
}

func PostItem(req *http.Request) *Response {
	r := NewResponse()
	decoder := json.NewDecoder(req.Body)
	i := back.NewItem()
	r.Action = "POST"
	err := decoder.Decode(&i.Notify)
	if err != nil {
		r.Err = err.Error()
		r.Status = http.StatusInternalServerError
		return r
	}
	id, err := back.PostItem(i)
	i.Id = id
	r.Item = i
	r.Status = http.StatusCreated
	return r
}

func DeleteItem(id string) *Response {
	r := NewResponse()
	intid, _ := strconv.Atoi(id)
	err := back.DeleteItem(int64(intid))
	r.Action = "DELETE"
	if err != nil {
		r.Err = err.Error()
		r.Status = http.StatusNotFound
		return r
	}
	r.Status = http.StatusAccepted
	return r
}

func ArchiveItem(id string) *Response {
	r := NewResponse()
	intid, _ := strconv.Atoi(id)
	r.Action = "PATCH"
	if err := back.ArchiveItem(int64(intid)); err != nil {
		r.Err = err.Error()
		r.Status = http.StatusNotFound
		return r
	}
	r.Item, _ = back.GetItem(int64(intid))
	r.Status = http.StatusAccepted
	return r
}
