package gollazo

import (
	"errors"
	"fmt"
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

var str_of_valid_digits string = "012345678"  // note 9 is excluded as it indicates higher Roman numeral value
var str_of_valid_bytes string = "0123456789O" // the full set of allowed bytes

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

	U, A, B, err := CheckCipher(cipher)
	if err != nil {
		return " ", err
	}

	A_arr, err := splitAtoIntArray(A)
	if err != nil {
		return " ", err
	}

	B_arr, err := splitBtoStrArray(B)
	if err != nil {
		return " ", err
	}

	fmt.Printf("%d\t%v\t%v\n", U, A_arr, B_arr)

	return " ", nil
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

// CheckCipher checks whether a passed cipher is consistent with the rules
// of the Collazo cipher. It checks whether there is any configuration of the final
// digits (U) which produces A and B with length U.
func CheckCipher(cipher string) (int, string, string, error) {

	l_c := len(cipher)

	// first check that there are no wrong characters
	for i := 0; i < l_c; i++ {
		if !strings.Contains(str_of_valid_bytes, cipher[i:i+1]) {
			return -1, " ", " ", errors.New("gollazo.CheckCipher: Not a Collazo cipher. Contains at least one invalid character")
		}
	}

	// there is no rule on the length of U
	// Check over all possible lengths of U, l_u.
	l_u := 1
	for {
		// define U assuming it has l_u digits
		U_str := cipher[len(cipher)-l_u:]

		// if leading digit is 0, this cannot be valid
		// for this l_u (but could be as part of a longer U,
		// so don't return yet)
		//
		// Also if there
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
			return -1, " ", " ", errors.New("gollazo.CheckCipher: Not a Collazo cipher. No possible U")
		}

		// using this value of U, extract A and B
		// also store length of B for use later in this function.
		A := cipher[l_c-l_u-U : l_c-l_u]
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
				// check that the A array only contains integers
				if strings.Contains(A, "O") {
					return -1, " ", " ", errors.New("gollazo.Decrypt: Not a Collazo cipher. A contains 9 or O")
				}
				return U, A, B, nil
			}

			// iterate through the string from the first byte.
			// If the leading byte indicates a triplet (O or 9), check that the next two bytes are decimal,
			// and move forward 3 bytes to the next leading byte.
			// If the leading byte indicates a doubled (0-8), check that the next two byte is decimal,
			// and move forward w bytes to the next leading byte.
			if (string(B[i]) == "O") || (string(B[i]) == "9") {
				if !(strings.Contains(str_of_valid_digits, string(B[i+1])) && strings.Contains(str_of_valid_digits, string(B[i+2]))) {
					l_u++
					break
				} else {
					i += 3
					continue
				}
			} else if !(strings.Contains(str_of_valid_digits, string(B[i])) && strings.Contains(str_of_valid_digits, string(B[i+1]))) {
				l_u++
				break
			} else {
				i += 2
			}
		}
	}
}

// splitAtoIntArray takes a string of integers and returns an integer array
// with each byte of the string assigned to each entry of the array.
func splitAtoIntArray(A string) ([]int, error) {
	var A_arr []int
	var A_int int
	var err error

	for i := 0; i < len(A); i++ {
		A_int, err = strconv.Atoi(A[i : i+1])
		A_arr = append(A_arr, A_int)
		if err != nil {
			return []int{}, errors.New("gollazo.Decrypt: Not a Collazo cipher. A contains non-integer")
		}
	}

	return A_arr, nil
}

// splitBtoStrArray takes a Collazo B string and returns an array of
// doublet and triplet strings.
// IMPORTANT: this function does not do any checking and works on the
// assumption that B has been returned by CheckCipher (i.e. that is a
// valid Collazo string). This is to avoid repeated logic.
func splitBtoStrArray(B string) ([]string, error) {

	var B_arr []string
	var l_b int = len(B)

	// loop through the B string, identifying doublets and triplets
	// and assign them to new array
	i := 0
	for {

		// if we have got this far then the B string
		// is consistent with a Collazo cipher
		if i == l_b {
			return B_arr, nil
		}

		// extract each term (2 bytes standard, 3 bytes if
		// triplet indicated by O or 9)
		if (string(B[i]) == "O") || (string(B[i]) == "9") {
			B_arr = append(B_arr, B[i:i+3])
			i += 3
			continue
		} else {
			B_arr = append(B_arr, B[i:i+2])
			i += 2
		}
	}
}
