package main

import (
	"fmt"
)

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {

	employees := map[string]Employee{
		"emp1": {"Khanh", 21, "Student"},
		"emp2": {"Hien", 22, "Teacher"},
		"emp3": {"Van", 27, "Developer"},
	}

	fmt.Println(employees)

	fmt.Printf("Name: %s \n", employees["emp1"].Name)
	fmt.Printf("Age: %d \n", employees["emp1"].Age)
	fmt.Printf("Role: %s \n", employees["emp1"].Role)

	for _, value := range employees {
		fmt.Printf("Name: %s | Age: %d | Role: %s\n", value.Name, value.Age, value.Role)
	}

	studentSub := map[string][]string{
		"Khoi": {"Golang", "HTML", "CSS", "JS"},
		"Anh":  {"Toan", "Ly", "Hoa", "Su"},
		"Tuan": {"Van", "Anh", "Sinh", "Cong Nghe"},
	}

	for key, value := range studentSub {
		for _, subject := range value {
			fmt.Printf("Mon hoc cua %s la: %s \n", key, subject)
		}
		fmt.Printf("Mon hoc cua %s la: %s \n", key, value[0])
	}

	// drink := map[string]int{
	// 	"coffe": 500,
	// 	"tea":   200,
	// }

	// fmt.Printf("%+v", drink)

	// student := map[int]string{
	// 	10: "Thanh",
	// 	9:  "Yen",
	// 	8:  "Khanh",
	// }

	// for key, value := range student {
	// 	fmt.Printf("Student[%d] is: %s\n", key, value)
	// }

	// //create map with make

	// m := make(map[string]int)

	// m["a"] = 1
	// m["b"] = 2
	// m["c"] = 3
	// m["d"] = 4
	// m["e"] = 5

	// fmt.Println(m)

	// var monAn map[string]int

	// monAn = make(map[string]int)

	// monAn["pho"] = 10
	// monAn["bun"] = 50
	// monAn["com"] = 30
	// monAn["banh mi"] = 20

	// fmt.Printf("%+v \n", monAn)

	// value, exists := monAn["bun"]

	// if exists {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Empty")
	// }

}
