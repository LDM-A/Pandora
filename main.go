package main

import "github.com/LDM-A/pandoratree/pkg/store"

func main() {

	kvStore := store.NewKVStorage()
	kvStore.Get("foo")

}
