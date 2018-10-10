package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s %d", e.name, e.currency, e.salary)
}

func main() {
	emp := Employee{
		name:     "z3",
		salary:   5000,
		currency: "￥",
	}
	displaySalary(emp)
}
