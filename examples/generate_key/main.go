package main

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	privateKey, err := hedera.GenerateEd25519PrivateKey()
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey()

	fmt.Printf("private = %v\n", privateKey)
	fmt.Printf("public = %v\n", publicKey)
}
