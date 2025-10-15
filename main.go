package main

import "fmt"

func main() {
	drink := map[string]int{
		"coffe": 500,
		"tea":   200,
	}

	fmt.Printf("%+v", drink)

	student := map[int]string{
		10: "Thanh",
		9:  "Yen",
		8:  "Khanh",
	}
	fmt.Printf("%+v", student)

	//create map with make

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5

	fmt.Println(m)

	var monAn map[string]int

	monAn = make(map[string]int)

	monAn["pho"] = 10
	monAn["bun"] = 50
	monAn["com"] = 30
	monAn["banh mi"] = 20

	fmt.Printf("%+v \n", monAn)
}
