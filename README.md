# gollazo
Exercise to use the Go programming language to solve "A Programming Challenge in Cryptography" by Eric Collazo (and submitted by ) as [accessed online December 2022](http://archive.dimacs.rutgers.edu/drei/1997/classroom/lessons/challenge.html). For reference, the challenge is quoted below (with some added clarificaations)

## The Problem
Write a program to decode a string of digits and the character "O"', and print the hidden message. [Ex. Find the hidden message in 84581248O6096095854123337.]

## The Private Key
Roman numerals: I = 12, V = 22, X = 24,
Characters: O = smaller, 9 = greater

## The Public Key
Assigned Roman numeral value of the letters of alphabet:
A = I, B = II, C = III, D = IV, ..., Z = XXVI

## The Rules
1. The last digit(s) U indicates the number of letters in the message.
[Ex. The last digit in the above example is 7]

2. Count U digits preceding the last digit(s). They are used with the remaining digits.
[Ex. The seven digits are 5412333]

3. The remaining digits are divided into either pairs or triples. The first digit of the triple must be either a "9" or an "O." The total number of pairs and triples should equal to the last digit(s) U. [Ex. 84 58 12 48 O60 960 958]

4. Match the results from Rule #2 and Rule #3 as follows:
[A] 5 4 1 2 3 3 3
[B]84 58 12 48 O60 960 958

Row [Al indicates the number of Roman numerals in the Private Key whose decimal values form the sum directly below it in row [B]. [Ex. In the first set, the 5 Roman numerals ( X X I I I ) whose decimal values ( 24, 24, 12, 12, 12 ) form the sum 84]. The decoded letter is the Roman numeral value in the Public Key. [Ex. XXIII = W]. This example translates to "WHATSUP" which is used for unit testing the decryption function. 

5. The first character in the triple is used to determine the configuration of the resulting Roman numeral. It is not used as a part of the sum. A "9" indicates a higher Roman numeral value; an "O", the lower value. [Ex. In the fifth set, O60 means the lower Roman numeral value formed by adding 3 Private Keys whose decimal value sum is 60:
The 3 Roman numerals whose decimal value ( 24, 24, 12 ) form the sum 60 are X
The "O" means that X X I should be decoded as X I X.] The decoded letter is therefore "S".]

The Test
Decode the following messages:
1) 704696084O36O583235236
2) O58O36362224462432311227
3) 60124858O36O60934960O583124232339