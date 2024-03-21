package main

import (
	"math"
	"os"
	"fmt"
	"time"
	"math/rand"
	// "encoding/csv"
	// "strconv"
	"encoding/json"
	"io/ioutil"
)

const data_dir = "../data"

type OutStuff struct {
	NumOrbits int
	ProcessTime float64
}

func main() {
	all_num_sats := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	var out_data []OutStuff
	for j:=0;j<len(all_num_sats);j++{
		num_orbits := all_num_sats[j]
		var process_time []float64
		for i:=0;i<1000;i++{	// For certainty
			orbits := generateUniqueRandomNumbers(num_orbits)
			process_time = append(process_time, run_process(orbits))
		}
		fmt.Println(num_orbits, "satellites picked giving an average time of", average(process_time), "orbits")
		out_data = append(out_data, OutStuff{
			NumOrbits: num_orbits,
			ProcessTime:  average(process_time),
		})
	}


	os.MkdirAll(data_dir, 0700)	// Should probably catch error here
	
	json_bits, err := json.Marshal(out_data)
	check(err)
	// fmt.Println(string(json_bits))
 
	err = ioutil.WriteFile("../data/simple_orbit_times.json", json_bits, 0644)
	check(err)
}

// Find average of array
func average(array []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(array); i++ {
	   sum += array[i]
	}
	return sum / float64(len(array))
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


func run_process(orbits []float64) (process_time float64) {
	
	// fmt.Println("All orbits:", orbits)
	primary_id := 0

	primary := orbits[primary_id]
	// fmt.Println("Chosen primary orbit:", primary)

	orbits_primary_removed := get_orbits_primary_removed(orbits, primary_id)
	// fmt.Println("Orbits with primary removed:", orbits_primary_removed)

	pairs := generate_pairs(len(orbits))
	// fmt.Println("Orbit Combinations:", pairs)

	theta_per := get_orbit_theta_steps(orbits)
	// fmt.Println("Theta Steps:", theta_per)

	// Inputs sender=primary recievers=primary_removed
	var t_now float64
	t_pre_prepare := get_step_times(t_now, orbits, pairs, theta_per, []float64{primary}, orbits_primary_removed)
	// Inputs sender=primary_removed recievers=orbits
	t_now = max(t_pre_prepare)
	t_prepare := get_step_times(t_now, orbits, pairs, theta_per, orbits_primary_removed, orbits)
	// Inputs sender=orbits recievers=orbits
	t_now = max(t_prepare)
	t_commit := get_step_times(t_now, orbits, pairs, theta_per, orbits, orbits)
	// Inputs sender=primary_removed recievers=primary
	t_now = max(t_commit)
	t_reply := get_step_times(t_now, orbits, pairs, theta_per, orbits_primary_removed, []float64{primary})

	// fmt.Println(t_pre_prepare)
	// fmt.Println(t_prepare)
	// fmt.Println(t_commit)
	// fmt.Println(t_reply)

	process_time = max(t_reply)
	// fmt.Println("Time of PBFT:", process_time)
	return process_time
}





func get_step_times(c_start float64, orbits []float64, pairs [][]int, theta_per []float64, senders []float64, recievers []float64) (t_step []float64) {
	var send float64
	var reci float64
	var pair0 float64
	var pair1 float64
	var next_val float64
	for i:=0;i<len(senders);i++{
		send = senders[i]
		for j:=0;j<len(recievers);j++{
			reci = recievers[j]
			for k:=0;k<len(pairs);k++{
				pair0 = orbits[pairs[k][0]]
				pair1 = orbits[pairs[k][1]]
				if (pair0 == send || pair0 == reci) && (pair1 == send || pair1 == reci) {
					next_val = 0
					for next_val <= c_start{
						next_val += theta_per[k]
					}
					t_step = append(t_step, next_val)
				}
			}
		}
	}
	return t_step
}





func get_orbit_theta_steps(orbits []float64) (theta_per []float64) {
	pairs := generate_pairs(len(orbits))

	var pair []int
	var x float64
	var y float64
	var theta_per_step float64
	for i:=0;i<len(pairs);i++ {
		pair = pairs[i]
		x = orbits[pair[0]]
		y = orbits[pair[1]]
		theta_per_step = (x*y)/(x-y)
		if theta_per_step < 0 {
			theta_per_step = theta_per_step * -1
		}
		theta_per = append(theta_per, theta_per_step)
	}
	return theta_per
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

// Gets all orbits minus the primary orbit
func get_orbits_primary_removed(orbits []float64, primary_id int)  (orbits_primary_removed []float64) {
	for i:=0;i<len(orbits);i++ {
		if i != primary_id {
			orbits_primary_removed = append(orbits_primary_removed, orbits[i])
		}
	}
	return orbits_primary_removed
}



// Generates all unique combinations of orbits
func generate_pairs(num int) (pairs [][]int) {
    for i:=0;i<num;i++ {
        for  j:=i+1;j<num;j++ {
			pairs = append(pairs, []int{i, j})
		}
	}
	return pairs
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}