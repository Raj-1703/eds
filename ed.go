package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	//generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	CheckError(err)

	publicKey := privateKey.PublicKey //public key can be generated from private key
	initialMessage := "i am Raj Kumar"

	encryptedMessage := RSA_OAEP_Encrypt(initialMessage, publicKey)

	fmt.Println("Encrypted Text:", encryptedMessage)

	RSA_OAEP_Decrypt(encryptedMessage, *privateKey)

}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
} //this function verifies whether the private key is generated or not.

func RSA_OAEP_Encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")                                                            //[]byte: represents byte stream
	rng := rand.Reader                                                                           //Random reader used for generating random bits so that no two input has same output
	encryptedtext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label) //key: public and private keys are generated previously.
	CheckError(err)
	return base64.StdEncoding.EncodeToString(encryptedtext)
}

func RSA_OAEP_Decrypt(encryptedtext string, privateKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(encryptedtext)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privateKey, ct, label)
	CheckError(err)
	fmt.Println("Plaintext:", string(plaintext))
	return string(plaintext)
}
