package main

import "fmt"

type Person struct {
	Name	string
	Age	int
	City	map[string]any
	Country	string
}

type Shop struct {
	Employees []Person
	Manager	[]Person
	Address	string
	YearlyRevenue	float64
}

type Family struct {
	Parent	[]Person
	Childreen	[]Person
	House	map[string]any
}

func bruh() {
	fmt.Println("Experimental stuff ...")
}

