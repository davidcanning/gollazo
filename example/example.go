package main

import (
	"fmt"
	"log"

	"github.com/davidcanning/gollazo/gollazo"
)

func main() {

	// the private key
	private_key := []int{24, 22, 12}

	// check a cipher
	num, A, B, err := gollazo.CheckCipher("84581248O6096095854123337")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("U = %d\tA = %s\tB = %s\terr = %v\n", num, A, B, err)

	// decrypt
	Z, err := gollazo.Decrypt("84581248O6096095854123337", private_key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Z)

}
