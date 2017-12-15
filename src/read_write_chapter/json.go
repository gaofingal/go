package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// json encoding
	js, _ := json.Marshal(vc)
	fmt.Printf("Json Ending Result:%s\n", js)

	// json decoding
	bJson := []byte(js)
	vc1 := VCard{}
	json.Unmarshal(bJson, &vc1)
	fmt.Printf("Json Decoding Result:%t\n", vc1)

	// output json file
	file, _ := os.OpenFile("vc.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

	// input json file
	inputJson, err := os.Open("vc.json")
	vc2 := VCard{}
	dnc := json.NewDecoder(inputJson)
	dnc.Decode(&vc2)
	fmt.Printf("FileDecodin-Json Decoding Result:%t\n", vc2)

}
