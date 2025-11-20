package main

import (
	"fmt"

	"github.com/laapkin/git-github_test/internal/lru"
)

func main() {

	cache := lru.New(3)

	cache.Put(1, "знач")
	cache.Put(2, "пер")
	cache.Put(3, "тер")
	cache.Put(4, "мер")
	fmt.Println(cache.Size())

	cache.Put(3, "sf")
	fmt.Println(cache.Size())

	val, ok := cache.Get(1)
	fmt.Println(val, ok)

	// val, ok = cache.Get(2)
	// fmt.Println(val, ok)

	val, ok = cache.Get(2)
	fmt.Println(val, ok)

	val, ok = cache.Get(4)
	fmt.Println(val, ok)

	cache.Put(5, "sf")
	val, ok = cache.Get(3)
	fmt.Println(val, ok)

	fmt.Println(cache.Size())
}
