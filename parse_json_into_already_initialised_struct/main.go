package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dummy struct {
	A, B string
}

func main() {
	d := Dummy{
		A: "A value",
	}

	onlyBStruct := struct {
		B string
	}{
		B: "B value",
	}

	data, error := json.Marshal(onlyBStruct)
	if error != nil {
		log.Fatal(error)
	}

	error = json.Unmarshal(data, &d)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%#v\n", d)

	dummy2 := Dummy{
		B: "B value",
	}

	data, error = json.Marshal(dummy2)
	if error != nil {
		log.Fatal(error)
	}

	error = json.Unmarshal(data, &d)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%#v", d)
}
