package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Dummy struct {
	A time.Time
	B *time.Time
}

func main() {
	none := Dummy{}

	data, error := json.Marshal(none)
	if error != nil {
		log.Fatal(error)
	}

	var d Dummy

	error = json.Unmarshal(data, &d)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%#v", d)
}
