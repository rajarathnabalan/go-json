package main

import (
	"log"

	"github.com/rajarathnabalan/go-json/jsonnode"
)

func main() {
	n := &jsonnode.Node{}
	n.Add(jsonnode.NewKeyValue("k1", "v"))
	log.Println(n.Value)
	log.Println(n.Get("k2"))
}
