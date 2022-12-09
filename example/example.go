package main

import (
	"fmt"

	"github.com/davidcanning/gollazo/gollazo"
)

func main() {

	// the private key
	private_key := []int{24, 22, 12}

	// decrypt
	fmt.Printf("Cipher: %s -> Plaintext: %s.\n", "84581248O6096095854123337", gollazo.Decrypt("84581248O6096095854123337", private_key))

	// encrypt
	fmt.Println(gollazo.Encrypt("Test", private_key))
}
