package main

import (
	"fmt"
)

func main() {

	//Excercise 1: Dem so phan tu trong mang
	// fruits := []string{"apple", "banana", "orange", "banana", "apple"}

	// // map nay luu key là tên trái cây, int: là số lần xuất hiện
	// counts := make(map[string]int)
	// // "apple" : 0

	// for _, value := range fruits {
	// 	counts[value]++
	// 	// mỗi lần thấy value có tồn tại trong counts thì sẽ tự cộng lên
	// }

	// fmt.Println("Số lần xuất hiện của mỗi loại trái cây có trong slice là: ")
	// for name, count := range counts {
	// 	fmt.Printf("%s xuất hiện số %d lần\n", name, count)
	// }

	//Excercise 2: Tra cứu thông tin sinh viên, xóa sinh viên

	// students := map[string]string{
	// 	"Anh":     "10a1",
	// 	"Binh":    "10a2",
	// 	"Lan Anh": "10a2",
	// 	"Van":     "10a3",
	// 	"Hung":    "10a6",
	// }

	// var confrim string
	// fmt.Print("Input name: ")
	// fmt.Scan(&confrim)

	// for key, value := range students {
	// 	if confrim == key {
	// 		fmt.Printf("%s' class: %s", key, value)
	// 		break
	// 	} else {
	// 		fmt.Println("Cannot found!!")
	// 		break
	// 	}
	// }

	// var confrimV2 string
	// fmt.Print("\nInput name delete: ")
	// fmt.Scan(&confrimV2)

	// for key := range students {
	// 	if confrimV2 == key {
	// 		delete(students, key)
	// 		fmt.Println("Deleted!!!")
	// 		for key, value := range students {
	// 			fmt.Println(key + "|" + value)
	// 		}
	// 		break
	// 	} else {
	// 		fmt.Println("Cannot found!!")
	// 		break
	// 	}
	// }

	// Excercise 4: GPA Student

	// gpaStudents := map[string][]float64{
	// 	"Tung": {7, 8, 8, 10, 8},
	// 	"Hien": {9.8, 7, 8, 10, 7.5},
	// 	"Van":  {8.8, 8.7, 8, 10, 9.5},
	// }

	// for key, val := range gpaStudents {
	// 	var total float64 = 0
	// 	for _, score := range val {
	// 		total += score
	// 	}
	// 	gpa := total / float64(len(val))
	// 	fmt.Printf("%s has gpa: %f \n", key, gpa)
	// }

	a := map[string]int{"apple": 3, "banana": 2}
	b := map[string]int{"apple": 1, "orange": 2}

	for keyB, valueB := range b {
		a[keyB] += valueB
	}

	fmt.Println(a)

	//Excercise 5:Đảo map

	ages := map[string]int{"An": 25, "Hung": 25, "Oanh": 28}

	reversed := map[int][]string{}
	for key, value := range ages {
		reversed[value] = append(reversed[value], key)
	}
	fmt.Println(reversed)

	//Excercise 6: Thống kế kí tự xuất hiện trong chuỗi

	var input string
	fmt.Println("Input string: ")
	fmt.Scan(&input)

	counts := make(map[rune]int)

	for _, char := range input {
		counts[char]++
	}

	fmt.Println("Thống kê kí tự xuất hiện")
	for key, value := range counts {
		fmt.Printf("%c: %d\n", key, value)
	}

	//GPA  pro

	scores := map[string]map[string]float64{
		"Hung": {"Math": 8, "English": 9, "Biology": 8.7},
		"Son":  {"Math": 7, "English": 10, "Biology": 7.7},
		"Uyen": {"Math": 7.5, "English": 7, "Biology": 9.7},
	}

	var topGPA float64
	var topStudent string
	for name, subjects := range scores {
		var total float64 = 0
		for _, score := range subjects {
			total += score
		}
		gpa := total / float64(len(subjects))
		fmt.Printf("%s has gpa: %.1f\n", name, gpa)
		if gpa > topGPA {
			topGPA = gpa
			topStudent = name
		}
	}
	fmt.Printf("%s is top studdent", topStudent)

}
