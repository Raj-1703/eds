package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	x := "kjj"
	h1 := sha1.New()
	h2 := sha256.New()
	h3 := sha512.New()
	//h1.Write([]byte(x))
	//h2.Write([]byte(x))
	//h3.Write([]byte(x))
	bx1 := h1.Sum([]byte(x))
	bx2 := h2.Sum([]byte(x))
	bx3 := h3.Sum([]byte(x))
	fmt.Printf("%x\n", bx1)
	fmt.Printf("%x\n", bx2)
	fmt.Printf("%x\n", bx3)
}
