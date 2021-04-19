package main

import (
	"fmt"
)

func main() {
	fmt.Println("Apenas um teste!")

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for k, v := range commits {
		fmt.Printf("key[%s] value[%d]\n", k, v)
	}
}
