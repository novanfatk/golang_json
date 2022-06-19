package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"url"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Mac",
		ImageUrl: "https://example.com",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac","url":"https://example.com"}`
	jsonByte := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonByte, product)
	fmt.Println(product)
}
