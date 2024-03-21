package main

import (
	"math/rand"
)

func main() {
	random_array_out := generateUniqueRandomNumbers(4)
}

// Generates a list of unique floats between 0 and 1 rounded to 2 decimals
func generateUniqueRandomNumbers(n int) []float64 {
	rand.Seed(time.Now().UnixNano())
	set := make(map[float64]bool)
	var result []float64
	for len(set) < n {
	   value := math.Round(rand.Float64()*10000)/1000
	   if !set[value] && value != 0 {
		  set[value] = true
		  result = append(result, value)
	   }
	}
	return result
}