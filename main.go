package main

import (
	"fmt"
	"github.com/chazari-x/computers-count/computers-count"
)

func main() {
	for i := range 50 {
		fmt.Println(computers_count.ComputersCount(i))
	}
}
