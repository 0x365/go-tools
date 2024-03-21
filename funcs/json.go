package main

import (
)

// Be careful must be capital
type OutStuff struct {
	Item1 int
	Item2 float64
}

const data_dir = "../data"

func main() {
	open_from_json()
	save_to_json()
}


func open_from_json() {
	// WIP
}




func save_to_json() {
	var out_data []OutStuff
	out_data = append(out_data, OutStuff{
		Item1: 324
		Item2: 2532.0243
	})

	// This makes sure file exists and creates it if not
	os.MkdirAll(data_dir, 0700)	// Should probably catch error here
	
	json_bits, err := json.Marshal(out_data)
	check(err)
	// fmt.Println(string(json_bits))
 
	err = ioutil.WriteFile("../data/file_name_example.json", json_bits, 0644)
	check(err)
}
