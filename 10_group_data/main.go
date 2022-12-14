package main

import (
	"fmt"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(groupTemp(temps))
}

func groupTemp(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, val := range temps {
		group := int(val/10) * 10

		groups[group] = append(groups[group], val)
	}

	return groups
}
