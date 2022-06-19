package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"novan","MiddleName":"java","LastName":"golang","Age":20,"Merried":false}`
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
	fmt.Println(customer.Age)
	fmt.Println(customer.Merried)
}
