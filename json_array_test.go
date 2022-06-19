package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Novan",
		MiddleName: "java",
		LastName:   "golang",
		Hobbies: []string{
			"Gaming",
			"Reading",
			"Coding",
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Novan","MiddleName":"java","LastName":"golang","Age":0,"Merried":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplexEncode(t *testing.T) {
	customer := Customer{
		FirstName: "novan",
		Address: []Address{
			{
				Street:     "jl buntu",
				Country:    "Tegal",
				PostalCode: "52122",
			},
			{
				Street:     "jl panggung",
				Country:    "Tegal timur",
				PostalCode: "52122",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"novan","MiddleName":"","LastName":"","Age":0,"Merried":false,"Hobbies":null,"Address":[{"Street":"jl buntu","Country":"Tegal","PostalCode":"52122"},{"Street":"jl panggung","Country":"Tegal timur","PostalCode":"52122"}]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Address)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"jl buntu","Country":"Tegal","PostalCode":"52122"},{"Street":"jl panggung","Country":"Tegal timur","PostalCode":"52122"}]`
	jsonByte := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonByte, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplexEncode(t *testing.T) {
	addresses := []Address{
		{
			Street:     "jl buntu",
			Country:    "Tegal",
			PostalCode: "52122",
		},
		{
			Street:     "jl panggung",
			Country:    "Tegal timur",
			PostalCode: "52122",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
