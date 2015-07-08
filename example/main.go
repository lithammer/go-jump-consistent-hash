package main

import (
	"fmt"

	"github.com/renstrom/go-jump-consistent-hash/jump"
)

func main() {
	h := jump.Hash(256, 1024) // h = 520
	fmt.Println(h)
}
