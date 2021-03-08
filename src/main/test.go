package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	key := "hihi"
	h := fnv.New32a()
	h.Write([]byte(key))
	a := int(h.Sum32() & 0x7fffffff)
	fmt.Println(a) 
}
