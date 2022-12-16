package gollazo

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// map roman numerals to the alphabet
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

// the Roman numerals used for the encoding
var roman_characters_str_array = []string{
	"X",
	"V",
	"I",
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

	var plaintext string
	var plainbyte string

	// do an initial check whether the passed cipher looks like
	// a Collazo cipher. If it is, identify the A and B strings
	U, A, B, err := CheckCipher(cipher)
	if err != nil {
		return " ", err
	}

	// split the A string into an array of integers.
	A_arr, err := splitAtoIntArray(A)
	if err != nil {
		return " ", err
	}

	// split the B string into an array of strings
	// (in order to retain higher/lower term)
	B_arr, err := splitBtoStrArray(B)
	if err != nil {
		return " ", err
	}

	// iterate through the arrays to find the corresponding
	// Roman numeral reputation (and there for the plaintext)
	for i := 0; i < U; i++ {
		plainbyte, err = translateABPair2Plaintext(A_arr[i], B_arr[i], private_key)
		plaintext = plaintext + plainbyte
	}
	if err != nil {
		return " ", err
	}

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

// translateABPair2Plaintext takes an integer, A_elem, returns the plaintext character encoded
// by a AB pair in a Collazo cipher.
// A_elem: number of roman numerals required to express B_elem
// B_elem: numer expressed as string. If leng(B_elem) == 3: "O" or "9" indicates lower/upper
func translateABPair2Plaintext(A_elem int, B_elem string, private_key []int) (string, error) {

	// variables
	var B_int int
	var triplet_flag bool = false
	var key_length int = len(private_key)
	var roman_numeral_arr []int
	var roman_representation string = ""

	// parse B_elem to flag and integer value
	// set a flag if it is a triplet
	if (string(B_elem[0]) == "O") || (string(B_elem[0]) == "9") {
		triplet_flag = true
		B_int, _ = strconv.Atoi(B_elem[1:3])
	} else {
		B_int, _ = strconv.Atoi(B_elem[0:2])
	}

	//

	// initialise the array to zeros
	for i := 0; i < key_length; i++ {
		roman_numeral_arr = append(roman_numeral_arr, 0)
	}

	// recursively iterate over all of the possible combinations with the
	// correct length (brute force). Note that roman numeral array is passed
	// as a pointer and is therefor modified by the function.
	solution_found := findRomanRecursion(A_elem, B_int, roman_numeral_arr, private_key, key_length)
	if solution_found != true {
		return " ", errors.New("gollazo.translateABPair2Plaintext: Not a Collazo cipher. Couldn't find Roman suitable numeral representation")
	}

	// construct the string repr. of the found Roman numeral
	for i := 0; i < key_length; i++ {
		roman_representation = roman_representation + strings.Repeat(roman_characters_str_array[i], roman_numeral_arr[i])
	}

	// account for the triplet higher/lower indicator
	roman_representation_length := len(roman_representation)
	if triplet_flag && (string(B_elem[0]) == "O") {
		roman_representation = roman_representation[:roman_representation_length-2] +
			roman_representation[roman_representation_length-1:roman_representation_length] +
			roman_representation[roman_representation_length-2:roman_representation_length-1]
	}

	return roman_2_letter[roman_representation], nil
}

// sumIntArray takes an array of integers as input and returns the
// sum of all elements
func sumIntArray(int_arr []int) int {
	var sum int
	for i := 0; i < len(int_arr); i++ {
		sum = sum + int_arr[i]
	}
	return sum
}

// sumElementwiseProduct takes two integer arrays, and returns the the sum
// of the products of each pair of  corresponding elements.
func sumElementwiseProduct(int_arr []int, private_key []int) (int, error) {
	var sum int

	if len(int_arr) != len(private_key) {
		return -1, errors.New("gollazo.sumElementwiseProduct: Passed arrays must have equal length")
	}

	for i := 0; i < len(int_arr); i++ {
		sum = sum + int_arr[i]*private_key[i]
	}
	return sum, nil
}

// findRomanRecursion performs a brute force recursive search through plausible combinations
// of the private key. Note that the cipher doesn't use a proper roman numeral system
// (i.e. it doesn't use "IV" or "IX" etc as bases; this type of Roman base is only used
// via the "O" string indicating that the final two terms should be switched).
func findRomanRecursion(A_elem int, B_elem int, roman_numeral []int, private_key []int, depth int) bool {

	// the roman_numeral int array is holding the number of each numeral.
	// e.g. [2,1,2] would mean 2 "X"s, 1 "V", and 2 "I"s, i.e. "XXVII"

	// A_elem sets an upper limit as once any individual character has
	// more than A_elem (the number of characters required) then there is
	// no need to search.
	for i := 0; i <= A_elem; i++ {

		roman_numeral[len(roman_numeral)-depth] = i
		if depth > 1 {
			solution_found := findRomanRecursion(A_elem, B_elem, roman_numeral, private_key, depth-1)
			if solution_found {
				return true
			}
		} else {
			num_numerals := sumIntArray(roman_numeral)
			sum_total, _ := sumElementwiseProduct(roman_numeral, private_key)

			if num_numerals < A_elem {
				continue
			} else if num_numerals > A_elem {
				break
			} else if (sum_total == B_elem) && (num_numerals == A_elem) {
				return true
			}
		}
	}

	return false
}
