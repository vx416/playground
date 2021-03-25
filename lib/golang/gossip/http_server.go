package main

import (
	"encoding/json"
	"net/http"

	"github.com/hashicorp/memberlist"
)

func (node *Node) handler() *handler {
	return &handler{node}
}

type handler struct {
	node *Node
}

func (h *handler) addHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := r.Form.Get("val")
	h.node.add(key, val)

	msg := &msg{
		Action: add, Key: key, Val: val,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	h.node.broadcasts.QueueBroadcast(&broadcast{
		msg:    b,
		notify: nil,
	})
	w.Write([]byte("ok"))
}

func (h *handler) delHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	h.node.del(key)

	b, err := json.Marshal([]*msg{{
		Action: add,
		Key:    key,
	}})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	h.node.broadcasts.QueueBroadcast(&broadcast{
		msg:    b,
		notify: nil,
	})
}

func (h *handler) getHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := h.node.get(key)
	w.Write([]byte(val))
}

type broadcast struct {
	msg    []byte
	notify chan<- struct{}
}

func (b *broadcast) Invalidates(other memberlist.Broadcast) bool {
	return false
}

func (b *broadcast) Message() []byte {
	return b.msg
}

func (b *broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}
