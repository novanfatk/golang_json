package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"macbook pro","price":"20000000"}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "macbook pro",
		"price": "2000000",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
