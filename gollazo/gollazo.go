package gollazo

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
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

var str_of_digits string = "012345678" // note 9 is excluded as it is used for something else

// Decrypt takes an encoded cipher (string) and the private key (array of integers) and produces
// the plain text equivalent using the Collazo cryptosystem.
//
// The private key array should be given in reverse order from largest to smallest Roman numeral, i.e:
// private_key[0] := 24  (X)
// private_key[0] := 22  (V)
// private_key[0] := 12  (I)
//
// The function will use the length of the key to map to roman numerals.
func Decrypt(cipher string, private_key []int) (string, error) {

	// first check if passed cipher is consistent with
	// the rules of the cryptosystem.
	if !IsCollazoCipher(cipher) {
		return " ", errors.New("Decrypt: Passed string is not a valid Collazo cipher")
	}

	// identify the seperate data strings within cipher
	A, B, err := extractAB(cipher)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A = %s\nB = %s \n", A, B)

	plaintext := convertAB2Plain(A, B, private_key)

	return plaintext, nil
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

// IsCollazoCipher checks whether a passed cipher is consistent with the rules
// of the Collazo cipher. It checks whether there is any configuration of the final
// digits (U) which produces A and B with length U.
func IsCollazoCipher(cipher string) bool {

	l_c := len(cipher)

	// there is no rule on the length of U
	// Check over all possible lengths of U, l_u.
	l_u := 1
	for {
		// define U assuming it has l_u digits
		U_str := cipher[len(cipher)-l_u:]

		// if leading digit is 0, this cannot be valid
		// for this l_u (but could be as part of a longer U,
		// so don't return yet)
		if U_str[0:1] == "0" {
			l_u++
			continue
		}

		// cast U, with length l_u, to integer
		U, _ := strconv.Atoi(U_str)

		// return false if number of bytes in the cipher subtract
		// number of bytes in U is less than the minimum length
		// required (i.e. 3 times the lower power). This assumes U
		// never has any leading zeroes (i.e "03" to mean "3")
		if l_c-l_u <= 3*int(math.Pow10(l_u-1)) {
			return false
		}

		// using this value of U, extract B and its length
		B := cipher[:l_c-l_u-U]
		l_b := len(B)

		// check that that total length of B is consistent with
		// number of triplets and doublets
		num_triplets := strings.Count(B, "O") + strings.Count(B, "9")
		num_doublets := (len(B) - 3*num_triplets) / 2
		if ((l_b-3*num_triplets)%2 != 0) || (num_triplets+num_doublets) != U {
			l_u++
			continue
		}

		// check that every triplet indicator ("O" or "9") has two
		// non-triplet indicators after it
		i := 0
		for {

			// if we have got this far then the B string
			// is consistent with a Collazo cipher
			if i == l_b {
				return true
			}

			if (string(B[i]) == "O") || (string(B[i]) == "9") {
				if !(strings.Contains(str_of_digits, string(B[i+1])) && strings.Contains(str_of_digits, string(B[i+2]))) {
					l_u++
					break
				} else {
					i += 3
					continue
				}
			} else if !(strings.Contains(str_of_digits, string(B[i])) && strings.Contains(str_of_digits, string(B[i+1]))) {
				l_u++
				break
			} else {
				i += 2
			}

		}

	}

}

// extract AB takes a cipher encrypted via the Collazo cryptosystem and returns two strings
// containing A, the number of Roman numerals in the Private Key whose decimal values form
// the corresponding value (i.e. same index) in the B array
// Note that the length of these arrays corresponds to the length of the encoded message.
func extractAB(cipher string) (string, string, error) {

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
