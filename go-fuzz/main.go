package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	// using bytes causing problem
	//The entire seed corpus used strings in which every character was a single byte.
	// However, characters such as 泃 can require several bytes.
	// Thus, reversing the string byte-by-byte will invalidate multi-byte characters.
	// r := []byte(s)
	// change from byte-by-byte with rune-by-rune
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	//will return an error if the input string contains characters which are not valid UTF-8.
	return string(r), nil
	// The key difference is that Reverse is now iterating over each rune in the string, rather than each byte.
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
