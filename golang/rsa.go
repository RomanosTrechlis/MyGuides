package main 

import (
	//"crypto"
	"crypto/rand"
	"crypto/rsa"
	//"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	bobPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	bobPublicKey := &bobPrivateKey.PublicKey
	alicePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	alicePublicKey := &alicePrivateKey.PublicKey

	fmt.Println("Bob's Private Key:", bobPrivateKey)
	fmt.Println("Bob's Public Key:", bobPublicKey)
	fmt.Println("Alice's Private Key:", alicePrivateKey)
	fmt.Println("Alice's Public Key:", alicePublicKey)
}
