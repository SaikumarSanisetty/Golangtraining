package main

import (
	"fmt"
)

//Package level varibale
var data map[int]Customer
var customer Customer

type datas map[int]Customer

type Customer struct {
	CustomerId int
	Name       string
	Age        int
	Mobile     int
}

// init function is used to initialise package level variables
func init() {

	data = make(map[int]Customer)
}
func (c Customer) add(k int, v Customer) {
	// Check if key exists or not
	if _, check := data[k]; check {
		fmt.Println(k, "Customer not added..")

	} else {
		data[k] = v
		fmt.Println(k, "Customer added")
	}
}
func (c Customer) delete(k int) {
	if _, check := data[k]; check {
		delete(data, k)
		fmt.Println(k, "Customer deleted")
	} else {
		fmt.Println("Customer Not Deleted")
	}
}
func (c Customer) update(k int, v Customer) {
	if _, check := data[k]; check {
		data[k] = v
		fmt.Println(k, "Customer updated")
	} else {
		fmt.Println("Customer Not Updated")
	}

}
func (c Customer) get(k int) (Customer, string) {
	if _, check := data[k]; check {
		return data[k], "Get data"
	}
	return Customer{}, " Not Found"
}
func (c Customer) getAll() datas {
	return data
}

func main() {
	c1 := Customer{101, "Saikumar", 23, 9866954863}
	c2 := Customer{102, "Sruthi", 25, 8241688652}
	c3 := Customer{103, "Kundana", 8, 9677246416}
	fmt.Println(data)

	customer.add(101, c1)
	customer.add(102, c2)
	customer.add(103, c3)
	customer.add(108, Customer{108, "Venkat", 30, 9575647483})
	fmt.Println(data)

	customer.update(106, Customer{102, "Sreehari", 36, 6584648488})
	fmt.Println(data)

	customer.delete(104)
	fmt.Println(data)

	customers := customer.getAll()
	fmt.Println(customers)
}
