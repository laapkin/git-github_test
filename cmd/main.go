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
	cache.Put(5, "кер")
	cache.Put(6, "aaa")
	cache.Put(7, "ada")

	val, ok := cache.Get(2)
	fmt.Println(val, ok)

	val, ok = cache.Get(4)
	fmt.Println(val, ok)

	fmt.Printf("%+v\n", cache.Data)
	fmt.Printf("%+v\n", cache.Order)
}
