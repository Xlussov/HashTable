package main

import (
	"fmt"

	"github.com/Xlussov/HashTable/hashtable"
)

func main() {
	ht := hashtable.New[int, string](8)

	ht.Put(1, "key1")
	ht.Put(2, "value2")
	ht.Put(3, "value3")
	ht.Put(4, "value4")
	ht.Put(5, "value5")
	ht.Put(6, "value6")
	ht.Put(7, "value7")
	ht.Print()
	fmt.Println("==============================")

	ht.Put(8, "value8")

	if v, ok := ht.Get(1); ok {
		fmt.Println(v)
	}

	if v, ok := ht.Remove(1); ok {
		fmt.Println(v)
	}

	if v, ok := ht.Get(1); ok {
		fmt.Println(v)
	}

	ht.Print()
}
