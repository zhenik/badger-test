package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	item1 := Item{Id: "uuid1", Field1: "f2", Field2: 32.4}

	// marshalling
	jsonMarshalled, err := json.Marshal(&item1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonMarshalled))

	// un-marshalling
	var item2 Item
	jsonPayload := `{"id":"uuid1","field1":"new value","field2":42}`
	err = json.Unmarshal([]byte(jsonPayload), &item2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(item2)

	store, err := NewBadgerDBFromPath("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	store.Upsert("uuid1", item1)
	item3, err := store.GetItem("uuid1")
	fmt.Println(item3)
	store.Upsert("uuid1", item2)
	item4, err := store.GetItem("uuid1")
	fmt.Println(item4)

	store.Upsert("mcrs", Item{Id: "mcrs", Field1: "some other stuf", Field2: 224})

	items, err := store.GetItems()
	fmt.Println(items)

	keys, err := store.GetAllKeys()
	fmt.Println(keys)

	fmt.Println("Deleting entry `uuid1`")
	store.DeleteItem("uuid1")

	items2, err := store.GetItems()
	fmt.Println(items2)

	keys2, err := store.GetAllKeys()
	fmt.Println(keys2)
}
