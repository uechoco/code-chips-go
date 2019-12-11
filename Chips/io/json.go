package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Sample01 struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	jsonDecodeFromFile()
	jsonDecodeFromString()
	jsonDecodeFromBytes()
	jsonUnmarshal()
}

func jsonDecodeFromFile() {
	file, err := os.Open("fixtures/sample01.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	var sample01 Sample01
	if err := json.NewDecoder(file).Decode(&sample01); err != nil {
		log.Fatalf("failed to decode json: %v", err)
	}

	fmt.Printf("decode json to struct: %+v\n", sample01)
}

func jsonDecodeFromString() {
	jsonStr := `{
	"id": 13,
	"name": "スズキ"
}`
	strReader := strings.NewReader(jsonStr)

	var sample01 Sample01
	if err := json.NewDecoder(strReader).Decode(&sample01); err != nil {
		log.Fatalf("failed to decode json: %v", err)
	}

	fmt.Printf("decode json to struct: %+v\n", sample01)
}

func jsonDecodeFromBytes() {
	jsonBytes := []byte(`{
		"id": 13,
		"name": "スズキ"
	}`)

	byteReader := bytes.NewReader(jsonBytes)

	var sample01 Sample01
	if err := json.NewDecoder(byteReader).Decode(&sample01); err != nil {
		log.Fatalf("failed to decode json: %v", err)
	}

	fmt.Printf("decode json to struct: %+v\n", sample01)
}

func jsonUnmarshal() {
	jsonBytes := []byte(`{
		"id": 13,
		"name": "スズキ"
	}`)
	var sample01 Sample01
	if err := json.Unmarshal(jsonBytes, &sample01); err != nil {
		log.Fatalf("failed to unmarshal json: %v", err)
	}

	fmt.Printf("unmarshal json to struct: %+v\n", sample01)
}
