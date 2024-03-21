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
	file_name := "this_is_a_file.csv"
	open_from_csv()
	save_to_csv()
}


func open_from_csv(file_name string) (a,x,y,z float64) {
	// Open the CSV file
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()
	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	check(err)
	for _, row := range data {
		a, _ := strconv.ParseFloat(row[0], 64)
		x, _ := strconv.ParseFloat(row[1], 64)
		y, _ := strconv.ParseFloat(row[2], 64)
		z, _ := strconv.ParseFloat(row[3], 64)
	}
	return a,x,y,z
}


func save_to_csv(data_1d []string) {
	// Save data to csv
	file, err := os.Create(data_dir+"/out_file_name.csv")
	defer file.Close()
	check(err)
	w := csv.NewWriter(file)
	defer w.Flush()
	w.Write(data_1d)	// Can iterate over this one
	// w.WriteAll(data_2d)
}


