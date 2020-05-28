package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("15h04m05s", "0h02m03s")
	d, err := time.ParseDuration("1h2m")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t)
	fmt.Println(d, d.Seconds())
}
