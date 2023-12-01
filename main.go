package main

import (
	"fmt"

	"github.com/LDM-A/pandoratree/pkg/store"
)

func main() {

	kvStore := store.NewKVStorage()
	if err := kvStore.Put("foo", []byte("Secret stuff")); err != nil {
		fmt.Println("Doesnt work")
	}
	data, err := kvStore.Get("foo")
	if err != nil {
		fmt.Println("Put doesnt work")
	}
	fmt.Printf("%s", string(data))

}
