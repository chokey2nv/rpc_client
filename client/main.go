package main

import (
	"log"
	"net/rpc"

	"github.com/chokey2nv/rpc_client/types"
)

func main() {
	var (
		reply types.Item
		db    []types.Item
	)
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	a := types.Item{"first", "first test item"}
	b := types.Item{"second", "second test item"}
	c := types.Item{"third", "third test item"}

	client.Call("API.AddItem", a, &reply)
	log.Println("Reply: ", reply)
	client.Call("API.AddItem", b, &reply)
	log.Println("Reply: ", reply)
	client.Call("API.AddItem", c, &reply)
	log.Println("Reply: ", reply)
	client.Call("API.GetDB", "", &db)
	log.Println("Reply: ", db)
}
