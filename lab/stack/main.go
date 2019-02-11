package main

import (
	"fmt"
)

func main() {
	t := stack{
		value: "navid",
	}

	t.Push("David")
	t.Push("Hassan")
	fmt.Println(t.Pull())
	fmt.Println(t.Pull())
	fmt.Println(t.Pull())
}
