package main

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

type eventDelegate struct{}

func (ed *eventDelegate) NotifyJoin(node *memberlist.Node) {
	fmt.Println("A node has joined: " + node.String())
	fmt.Printf("Node's meta: %s\n at %s", string(node.Meta), node.Address())
}

func (ed *eventDelegate) NotifyLeave(node *memberlist.Node) {
	fmt.Println("A node has left: " + node.String())
	fmt.Printf("Node's meta: %s\n at %s", string(node.Meta), node.Address())
}

func (ed *eventDelegate) NotifyUpdate(node *memberlist.Node) {
	fmt.Println("A node was updated: " + node.String())
	fmt.Printf("Node's meta: %s\n at %s", string(node.Meta), node.Address())
}
