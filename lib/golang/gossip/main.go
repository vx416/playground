package main

import (
	"flag"
	"log"
)

var (
	port = flag.String("port", "8001", "http port")
)

func init() {
	flag.Parse()
}

func main() {
	node, err := New("127.0.0.1", *port)
	if err != nil {
		log.Panicf("new node failed, err:%+v", err)
	}
	node.Serve()
}
