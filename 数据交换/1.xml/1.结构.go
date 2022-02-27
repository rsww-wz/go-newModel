package main

type Book struct {
	Name   string
	Author string
	Price  float64
}

type Person struct {
	Name      string
	Gender    bool
	Age       int
	Telephone string
	Book      Book
}
