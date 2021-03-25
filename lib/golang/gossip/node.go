package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/hashicorp/memberlist"
)

type action int8

const (
	del action = iota + 1
	add
)

type msg struct {
	Key    string `json:"key"`
	Val    string `json:"val"`
	Action action `json:"action"`
}

type Node struct {
	addr       string
	port       string
	env        string
	rw         sync.RWMutex
	kv         map[string]string
	broadcasts *memberlist.TransmitLimitedQueue
}

func (n *Node) add(key string, val string) {
	n.rw.Lock()
	defer n.rw.Unlock()
	n.kv[key] = val
}

func (n *Node) del(key string) {
	n.rw.Lock()
	defer n.rw.Unlock()
	delete(n.kv, key)
}

func (n *Node) get(key string) string {
	n.rw.RLock()
	defer n.rw.RUnlock()
	return n.kv[key]
}

func (n *Node) NodeMeta(limit int) []byte {
	return make([]byte, 0, limit)
}

func (n *Node) NotifyMsg(b []byte) {
	if len(b) == 0 {
		return
	}

	// log.Printf("notify msg rev msg: %s", string(b))
	msg := &msg{}

	err := json.Unmarshal(b, msg)
	if err != nil {
		log.Printf("notify msg: umarshal failed, err:%+v", err)
	}
	log.Printf("notify msg rev msg: %+v", msg)
	switch msg.Action {
	case add:
		n.add(msg.Key, msg.Val)
	case del:
		n.del(msg.Key)
	}
}

func (n *Node) GetBroadcasts(overhead, limit int) [][]byte {
	return n.broadcasts.GetBroadcasts(overhead, limit)
}

func (n *Node) LocalState(join bool) []byte {
	n.rw.RLock()
	defer n.rw.RUnlock()

	if join {
		data, err := json.Marshal(n.kv)
		if err != nil {
			return []byte{}
		}
		return data
	}
	return []byte{}
}

func (n *Node) MergeRemoteState(buf []byte, join bool) {
	if !join {
		return
	}

	n.rw.Lock()
	defer n.rw.Unlock()
	json.Unmarshal(buf, &n.kv)
}

func (node *Node) Serve() {
	handler := node.handler()
	http.HandleFunc("/add", handler.addHandler)
	http.HandleFunc("/del", handler.delHandler)
	http.HandleFunc("/get", handler.getHandler)
	httpPort := "1" + node.port
	fmt.Printf("Listening on :%s\n", httpPort)
	if err := http.ListenAndServe(":"+httpPort, nil); err != nil {
		fmt.Println(err)
	}
}

func New(addr, port string) (*Node, error) {
	node := &Node{
		addr: addr,
		port: port,
		env:  "local",
		kv:   make(map[string]string),
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	config := memberlist.DefaultLocalConfig()
	config.AdvertisePort = p
	config.BindPort = p
	config.Events = &eventDelegate{}
	config.Delegate = node
	hostname, _ := os.Hostname()
	config.Name = hostname + "-" + port
	m, err := memberlist.Create(config)
	if err != nil {
		return nil, err
	}
	_, err = m.Join([]string{"127.0.0.1:8001", "127.0.0.1:8002"})
	if err != nil {
		return nil, err
	}

	node.broadcasts = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return m.NumMembers()
		},
		RetransmitMult: 3,
	}

	return node, nil
}
