package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	//generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)

	publicKey := privateKey.PublicKey
	fmt.Println(publicKey)

	x := "kjjinHBVJDCNM"
	h1 := sha256.New()
	h1.Write([]byte(x))
	bx1 := h1.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, bx1, nil)
	CheckError(err)

	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, bx1, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return
	}

	fmt.Println("signature verified")
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}
