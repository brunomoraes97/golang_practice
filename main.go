package main

import "fmt"

// função que será chamada
func olaMundo() {
	fmt.Println("Olá Mundo!")
}

func dois() {
	fmt.Println("Dois")
}

func tres() {
	fmt.Println("Tres")
}

// função principal
// entrypoint -> A função que será chamada
func main() {
	olaMundo()
	dois()
	tres()
}
