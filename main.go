package main

import "fmt"

func main() {

	var a = "initial" // Não precisa colocar o tipo
	fmt.Println(a)

	var b, c int = 1, 2 // Mas também dá para tipar
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	var f string
	fmt.Println(f)

	g := "maçã"
	fmt.Println(g)

}
