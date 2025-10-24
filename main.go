package main

import (
	"fmt"

	"github.com/Xlussov/HashTable/hashtable"
)

func main() {
	ht := hashtable.New[string, int](8)
	ht.Put("alah1", 1)
	ht.Put("alah2", 2)
	ht.Put("alah3", 3)
	ht.Put("alah4", 4)
	ht.Put("alah5", 5)
	ht.Put("alah6", 6)
	ht.Put("alah7", 7)
	ht.Print()
	fmt.Println("==============================")

	// ht.Put("alah7", 7)
	ht.Put("alah8", 8)

	// if v, ok := ht.Get("alah1"); ok {
	// 	fmt.Println(v)
	// }

	// if v, ok := ht.Remove("alah1"); ok {
	// 	fmt.Println(v)
	// }

	// if v, ok := ht.Get("alah1"); ok {
	// 	fmt.Println(v)
	// }

	ht.Print()

	// ht.Put("alah1", 1488)

	// ht.Print()
}
