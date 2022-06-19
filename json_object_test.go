package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, Country, PostalCode string
}

type Customer struct {
	FirstName, MiddleName, LastName string
	Age                             int
	Merried                         bool
	Hobbies                         []string
	Address                         []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Novan",
		MiddleName: "java",
		LastName:   "golang",
		Age:        20,
		Merried:    false,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
