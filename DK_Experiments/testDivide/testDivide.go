package main

import (
	"fmt"
	"math"
)

func main() {

	as := []int{25, 36, 47, 48, 49}
	bs := []int{24, 24, 24, 24, 24}

	for i := 0; i < len(as); i++ {
		fmt.Printf("%d / %d = %d \n", as[i], bs[i], as[i]/bs[i])
		fmt.Printf("%d / %d = %d \n", as[i], bs[i], int(math.Round(float64(as[i]/24))))
		fmt.Printf("%d / %d = %d \n", as[i], bs[i], int(math.Round(float64(as[i])/24)))
		fmt.Println(`\\\\\\\\\\\\\\\\\`)
	}
}
