package main

import (
)

func main() {
	var array_in []float64
	average_of_array := average(array_in)
	max_val_in_array := max(array_in)
}

// Find average of array
func average(array []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(array); i++ {
	   sum += array[i]
	}
	return sum / float64(len(array))
}

// Finds max value in an array
func max(array []float64) (max_val float64){
	for i:=0;i<len(array);i++{
		if array[i] > max_val {
			max_val = array[i]
		}
	}
	return max_val
}