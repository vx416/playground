package main

import "log"

func main() {
	log.Printf("SELECT * FROM test WHERE id = ?", 1)
}
