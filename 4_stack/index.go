package main

import "fmt"

func main() {
	stack()
}

func stack() {
	hi()
	hwu()
	bye()
}

func hi() {
	fmt.Println("Hello!")
}

func hwu() {
	fmt.Println("How are you?")
}

func bye() {
	fmt.Println("Bye!")
}