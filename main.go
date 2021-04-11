package main

import "fmt"

type Item struct {
	Id     string  `json:"id"`
	Field1 string  `json:"field1"`
	Field2 float32 `json:"field2"`
}

func main() {
	item := Item{Id: "uuid1"}
	fmt.Println(item)
}
