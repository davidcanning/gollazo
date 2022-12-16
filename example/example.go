package main

import (
	"fmt"
	"log"

	"github.com/davidcanning/gollazo/gollazo"
)

func main() {

	// the private key given in the challenge
	private_key := []int{24, 22, 12}

	// the test examples to be decoded (test 0 is given example):
	test0 := "84581248O6096095854123337"
	test1 := "704696084O36O583235236"
	test2 := "O58O36362224462432311227"
	test3 := "60124858O36O60934960O583124232339"

	// decrypt the example
	plaintext0, err := gollazo.Decrypt(test0, private_key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cipher0: %s\t\tPlaintext: %s\n", test0, plaintext0)

	// decrypt the tests
	plaintext1, err := gollazo.Decrypt(test1, private_key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cipher1: %s\t\t\tPlaintext: %s\n", test1, plaintext1)

	plaintext2, err := gollazo.Decrypt(test2, private_key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cipher1: %s\t\tPlaintext: %s\n", test2, plaintext2)

	plaintext3, err := gollazo.Decrypt(test3, private_key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cipher1: %s\tPlaintext: %s\n", test3, plaintext3)
}
