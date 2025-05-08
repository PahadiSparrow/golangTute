package main

// https://go-proverbs.github.io/

import (
	"fmt"

	"github.com/PahadiSparrow/golangTute/puppy"
)

var (
	m = 23
	n = 56
)

func main() {
	name := "kim"
	age := 19
	fmt.Printf("%s IS %d years old and type of name is %T and age is %T", name, age, name, age)

	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%v %b %x\n", a, a, a)
	fmt.Printf("%v %b %X\n", b, b, b)
	fmt.Printf("%v %b %X\n", c, c, c)
	fmt.Printf("%v %b %X\n", d, d, d)
	fmt.Printf("%v %b %X\n", e, e, e)
	fmt.Printf("%v %b %X\n", f, f, f)
	fmt.Printf("%v %v \n", m, n)
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.Category)
	fmt.Println(puppy.TagVer3())
}
