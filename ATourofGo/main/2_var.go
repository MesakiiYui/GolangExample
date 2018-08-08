package main

import "fmt"

// The var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level. We see both in this example.
var c = 1
var d, e bool

// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, j int = 1, 2

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

// n:=4 => := not work outside
func main() {
	fmt.Println(c, d, e, i, j)
	m:=3
	fmt.Println(m)
}
