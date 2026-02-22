package main

import (
	"fmt"
	"lld-examples/internal/lru/service"
)

func main() {
	lru := service.NewLRUCache(3)

	lru.Put("a", 1)
	lru.Put("b", 2)
	lru.Put("c", 3)

	val, _ := lru.Get("a") // Access 'a' → now most recently used
	fmt.Println(val)       // 1

	lru.Put("d", 4) // Evicts least recently used ('b')

	_, ok := lru.Get("b")
	fmt.Println(ok) // false, 'b' evicted
}
