package main

import (
	"fmt"
	"strconv"
)

// Package-level variable declaration
// := format does not work here
var l int = 99

// Variables like this clutter the code
// so use a var block instead
// multiple blocks are allowed
// so you can group related items into blocks
var (
	thirdBase     string = "Carney Lansford"
	leftField     string = "Rickey Henderson"
	acePitcher    string = "Dave Stewart"
	managerNumber int    = 10
	titleYear            = 1989
)

func main() {
	// Declaration of variables
	// var name type
	// Useful for counters in loops
	var i int
	i = 42
	fmt.Println(i)

	// Change the value
	i = 27
	fmt.Println(i)

	// Declare the variable and assign value
	// on the same line
	// Useful for finer-grained control over type
	var j int = 39
	fmt.Println(j)

	// Go can figure out the type on its own
	// For numbers, will only do into or float64
	k := 84
	fmt.Println(k)

	// Replace Println with Printf for formatting
	// %v is a value. %T is the variable type
	// Each is replaced with variables outside of quotes
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// Print the package-level variable
	fmt.Printf("%v, %T\n", l, l)

	// innermost declared variable takes precedence
	// meaning same variable inside a function replaces package-level
	// This is called 'shadowing'

	// Variables that are declared but not used produce errors

	// Changing types
	var m int = 31
	fmt.Printf("%v, %T\n", m, m)

	var n float32
	// float32 here is a built-in conversion function
	n = float32(m)
	fmt.Printf("%v, %T\n", n, n)

	// float to int loses accuracy
	// n = m would fail without conversion because two different types

	// Converstion to string
	var o int = 102
	fmt.Printf("%v, %T\n", o, o)

	var p string
	p = string(o)
	fmt.Printf("%v, %T\n", p, p)
	// fmt.Printf("%v, %T\n", o, o)
	// The above produces a unicode value not a string
	// Conversion to string requires strconv package
	// https://golang.org/pkg/strconv/ for useage instructions
	p = strconv.Itoa(o)
	fmt.Printf("%v, %T\n", p, p)

}
