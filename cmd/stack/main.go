package main

import (
	"fmt"

	"github.com/laapkin/git-github_test/internal/stack"
)

func main() {
	s := stack.New(5)
	// fmt.Println(s.Size())

	fmt.Println(s.Push("5"))
	// fmt.Println(s.Size())

	fmt.Println(s.Push("11"))
	// fmt.Println(s.Size())

	fmt.Println(s.Push("222"))
	// fmt.Println(s.Size())
	// s.Pop()
	// fmt.Println(s.Size())
	// fmt.Println(s)

	fmt.Println(s.Push("2"))
	// fmt.Println(s.Size())

	fmt.Println(s.Push("3333"))
	// fmt.Println(s.Size())

	fmt.Println(s.Push("55555"))
	// fmt.Println(s.Size())

	s.Pop()
	// fmt.Println(s.Size())
	fmt.Println(s)

}
