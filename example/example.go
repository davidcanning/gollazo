package main

import (
	"fmt"

	"github.com/davidcanning/gollazo/gollazo"
)

func main() {

	// the private key
	private_key := []int{24, 22, 12}

	// check a cipher
	result := gollazo.IsCollazoCipher("11223344556677881110111213777777777777713")
	fmt.Println("Result = ", result)

	// decrypt
	test0_plain, _ := gollazo.Decrypt("84581248O6096095854123337", private_key)
	fmt.Printf("Cipher: %s -> Plaintext: %s.\n", "84581248O6096095854123337", test0_plain)

	// encrypt
	fmt.Println(gollazo.Encrypt("Test", private_key))
}
