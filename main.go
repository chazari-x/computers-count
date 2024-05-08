package main

import (
	"fmt"
	cc "github.com/chazari-x/computers-count/computers-count"
)

func main() {
	for i := range 50 {
		fmt.Println(cc.ComputersCount(i))
	}
}
