package main

import (
	"fmt"

	"github.com/Xlussov/HashTable/hashtable"
)

func main() {
	ht := hashtable.New[string, int](8)

	ht.Put("value1", 1)
	ht.Put("value2", 2)
	ht.Put("value3", 3)
	ht.Put("value4", 4)
	ht.Put("value5", 5)
	ht.Put("value6", 6)
	ht.Put("value7", 7)
	ht.Print()
	fmt.Println("==============================")

	ht.Put("value8", 8)

	if v, ok := ht.Get("value1"); ok {
		fmt.Println(v)
	}

	if v, ok := ht.Remove("value1"); ok {
		fmt.Println(v)
	}

	if v, ok := ht.Get("value1"); ok {
		fmt.Println(v)
	}

	ht.Print()
}
