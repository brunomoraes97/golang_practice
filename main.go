package main

import (
	"fmt"
	"slices"
)

func main() {

	// Take variable S, which will be an array of strings
	var s []string
	// Print "uninit", whatever that means => uninit:
	// Print the value of the variable s, which will be exactly [] (an empty array) => []
	// Check if the variable s is equal to nil, which is true => true
	// Check if the variable S' length is equal to zero, which is true => true
	// uninit: [] true true
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Take variable S and "make" it an array of three empty spaces
	s = make([]string, 3)
	// Print "emp:" => emp:
	// Print the value of the variable of S, which consists of an array of three empty spaces => [   ]
	// Print "len:" => len:
	// Print the length of the variable of S, which should be three => 3
	// Print "cap:" => cap:
	// Print the capability of the S array, which should be three (three empty spaces) => 3
	// emp: [   ] len: 3 cap: 3
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Take the variable S and put the value "a" into the position 0 of the empty array
	s[0] = "a"
	// Take the variable S and put the value "b" into the position 1 of the not-empty-anymore array
	s[1] = "b"
	// Take the variable S and put the value "c" into the position 2 of the array
	s[2] = "c"
	// Print "set:" => set:
	// Print the value of S => [a b c]
	// Result => set: [a b c]
	fmt.Println("set:", s)
	// Print "get:" => get:
	// Print the value of whatever is inside the position [2] of the S array => c
	// Result => get: c
	fmt.Println("get:", s[2])
	// Print "len:" => len:
	// Print the length of the S array => 3
	// Result => len: 3
	fmt.Println("len:", len(s))
	// Append is a built-in function that takes two arguments: the original array variable and the new value(s) you want to append
	// We will call the append function and add the S array and the value "d" as arguments
	// The value returned by this function will be stored in the variable S
	s = append(s, "d")
	// We're also appending "e" and "f" to the same array and storing them again in the S variable
	s = append(s, "e", "f")
	// Print "apd:" => apd:
	// Print the value of the variable S
	// Result => apd: [a b c d e f]
	fmt.Println("apd:", s)

	// Take the variable C and "make" it an empty array with the same capability as the length of the S array
	c := make([]string, len(s))
	// Copy the elements present in the variable S to the variable C
	copy(c, s)
	// Print "cpy:" => cpy
	// Print the new value of the variable C, which will be the same as the variable S => [a b c d e f]
	// Result => cpy: [a b c d e f]
	fmt.Println("cpy:", c)

	// Take the variable L and store whatever there is in between the positions [2] and [5] of the S array
	// Print "sl1:"
	// Print the value of L => [c d e]
	// Result => sl1: [c d e]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Take the variable L and store whatever there is in positions from 0 to 4 in the S array (stop storing values from the [5] position)
	l = s[:5]
	// Print "sl2:"
	// Print the value of L => [a b c d e]
	// Result => sl2: [a b c d e]
	fmt.Println("sl2:", l)

	// Take the variable L and store whatever there is starting from the position [2] of the S array
	l = s[2:]
	// Print "sl3:"
	// Print the value of L, which will be [c d e f]
	// Result => sl3: [c d e f]
	fmt.Println("sl3:", l)

	// Take the variable T and store an array of strings, adding the values "g", "h" and "i" to it
	t := []string{"g", "h", "i"}
	// Print "dcl:" => dcl:
	// Print the value of T => [g h i]
	// Result => dcl: [g h i]
	fmt.Println("dcl:", t)

	// Take the variable T2 and store an array of strings, adding the values "g", "h and "i" to it
	t2 := []string{"g", "h", "i"}
	// Use the "Equal" function that belongs to the "slices" library, which takes two arrays as arguments and compare them
	// If they are equal, the function will return true
	// If they are not equal, it will return false
	// What we are doing here is checking whether slices.Equal will return true
	// If it returns true, then "t == t2" should be printed in the terminal
	// In this case, both arrays (t and t2) store the same values, so the function will return true
	// Result => t == t2
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
