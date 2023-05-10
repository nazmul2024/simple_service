package main

import "fmt"

func main() {
	fmt.Println(run())
	fmt.Println("This is First Travis Build")
}

func run() string {
	return "Setup Travis CI for Golang project"
}
