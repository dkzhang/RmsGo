package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

type StruA struct {
	Name    string
	Age     int
	Content interface{}
}

type StruB struct {
	Bp1 string
	Bp2 int
}

type StruC struct {
	Cp1 int
	Cp2 string
}

func main() {
	a1 := StruA{
		Name: "A",
		Age:  18,
		Content: StruB{
			Bp1: "BB",
			Bp2: 22,
		},
	}
	a2 := StruA{
		Name: "A",
		Age:  18,
		Content: StruC{
			Cp1: 333,
			Cp2: "CCC",
		},
	}

	b1, err := json.Marshal(a1)
	if err != nil {
		logrus.Errorf("json.Marshal(a1) error: %v", err)
	}
	fmt.Printf("%s \n", b1)

	b2, err := json.Marshal(a2)
	if err != nil {
		logrus.Errorf("json.Marshal(a2) error: %v", err)
	}
	fmt.Printf("%s \n", b2)
}
