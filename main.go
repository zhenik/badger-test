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
	jsonPayload := `{"id":"uuid1","field1":"f2","field2":32.4}`
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

	item3, err := store.Get("uuid1")

	fmt.Println(item3)


}
