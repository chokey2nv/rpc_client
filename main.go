package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/chokey2nv/rpc_client/types"
)

type API int

var database []types.Item

func (a *API) GetDB(title string, reply *[]types.Item) error {
	*reply = database
	return nil
}
func (a *API) GetByTitle(title string, reply *types.Item) error {
	var item types.Item
	for _, val := range database {
		if val.Title == title {
			item = val
		}
	}
	*reply = item
	return nil
}

func (a *API) CreateItem(item types.Item, reply *types.Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) AddItem(item types.Item, reply *types.Item) error {
	database = append(database, item)
	*reply = item
	return nil
}
func (a *API) EditItem(item types.Item, reply *types.Item) error {
	var changedItem types.Item
	for idx, val := range database {
		if val.Title == item.Title {
			database[idx] = types.Item{
				Title: item.Title,
				Body:  item.Body,
			}
			changedItem = database[idx]
		}
	}
	*reply = changedItem
	return nil
}

func (a *API) DeleteItem(item types.Item, reply *types.Item) error {
	var del types.Item
	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
		}
	}
	*reply = del
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering api: ", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listener error")
	}

	log.Printf("Serving rpc on port: %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("http server error: ", err)
	}
	// fmt.Println("Initial database", database)
	// a := Item{"first", "first test item"}
	// b := Item{"second", "second test item"}
	// c := Item{"third", "third test item"}

	// // AddItem(a)
	// AddItem(b)
	// AddItem(c)

	// fmt.Println(" Second database ", database)

	// DeleteItem(b)
	// fmt.Println(" Third database ", database)

	// EditItem("third", Item{"fourth", "fourth item"})
	// fmt.Println("Fourth database ", database)

	// x := GetByTitle("fourth")
	// y := GetByTitle("first")

	// fmt.Println(x, y)
}
