package gollazo

import (
	"errors"
	"fmt"
	"log"
)

// using map literals
var roman_2_letter = map[string]string{
	"I":     "A",
	"II":    "B",
	"III":   "C",
	"IV":    "D",
	"V":     "E",
	"VI":    "F",
	"VII":   "G",
	"VIII":  "H",
	"IX":    "I",
	"X":     "J",
	"XI":    "K",
	"XII":   "L",
	"XIII":  "M",
	"XIV":   "N",
	"XV":    "O",
	"XVI":   "P",
	"XVII":  "Q",
	"XVIII": "R",
	"XIX":   "S",
	"XX":    "T",
	"XXI":   "U",
	"XXII":  "V",
	"XXIII": "W",
	"XXIV":  "X",
	"XXV":   "Y",
	"XXVI":  "Z",
}

// Decrypt takes an encoded cipher (string) and the private key (array of integers) and produces
// the plain text equivalent using the Collazo cryptosystem.
//
// The private key array should be given in reverse order from largest to smallest Roman numeral, i.e:
// private_key[0] := 24  (X)
// private_key[0] := 22  (V)
// private_key[0] := 12  (I)
//
// The function will use the length of the key to map to roman numerals.
func Decrypt(cipher string, private_key []int) string {

	// identify the seperate data strings within cipher
	A, B, err := extractAB(cipher)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A = %v\nB = %v \n", A, B)

	plaintext := convertAB2Plain(A, B, private_key)

	return plaintext
}

// Encrypt takes a plain text string (plaintext) and a private key (array of integers) and produces
// an encoded cipher (string) using the Collazo cryptosystem.
//
// The private key array should be given in reverse order from largest to smallest Roman numeral, i.e:
// private_key[0] := 24  (X)
// private_key[0] := 22  (V)
// private_key[0] := 12  (I)
//
// The function will use the length of the key to map to roman numerals.
func Encrypt(plaintext string, private_key []int) string {
	return "Encrypt Return"
}

func IsCollazoCipher(cipher string) bool {
	return true
}

// extract AB takes a cipher encrypted via the Collazo cryptosystem and returns two strings
// containing A, the number of Roman numerals in the Private Key whose decimal values form
// the corresponding value (i.e. same index) in the B array
// Note that the length of these arrays corresponds to the length of the encoded message.
func extractAB(cipher string) (string, string, error) {

	if !IsCollazoCipher(cipher) {
		return "", "", errors.New("extractAB: Passed string is not a valid Collazo cipher")
	}

	A := "5412333"
	B := "84581248O60960958"

	return A, B, nil
}

func convertAB2Plain(A string, B string, private_key []int) string {

	// split the strings up into arrays
	A_split := splitAtoIntArray(A)
	B_split := splitBtoStrArray(B)

	// decrypt each AB pair one by one
	// e.g. A=5 B="84" -> "XXIII" -> "W" with example key
	var plaintext string = ""
	var B_split_roman []string

	for i := 0; i < len(A_split); i++ {
		decrypted_roman := translateAB2Roman(A_split[i], B_split[i], private_key)
		B_split_roman = append(B_split_roman, decrypted_roman)
		plaintext += roman_2_letter[decrypted_roman]
	}

	return plaintext
}

func splitAtoIntArray(A string) []int {
	A_split := []int{5, 4, 1, 2, 3, 3, 3}
	return A_split
}

func splitBtoStrArray(B string) []string {
	B_split := []string{"84", "58", "12", "48", "O60", "960", "958"}
	return B_split
}

func translateAB2Roman(num_numerals int, sum_decimal string, private_key []int) string {
	return "XXIII"
}
