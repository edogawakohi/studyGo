package main

import "fmt"

func main() {

	//Excercise 1: Dem so phan tu trong mang
	fruits := []string{"apple", "banana", "orange", "banana", "apple"}

	// map nay luu key là tên trái cây, int: là số lần xuất hiện
	counts := make(map[string]int)
	// "apple" : 0

	for _, value := range fruits {
		counts[value]++
		// mỗi lần thấy value có tồn tại trong counts thì sẽ tự cộng lên
	}

	fmt.Println("Số lần xuất hiện của mỗi loại trái cây có trong slice là: ")
	for name, count := range counts {
		fmt.Printf("%s xuất hiện số %d lần\n", name, count)
	}

	//Excercise 2: Tra cứu thông tin sinh viên, xóa sinh viên

	students := map[string]string{
		"Anh":     "10a1",
		"Binh":    "10a2",
		"Lan Anh": "10a2",
		"Van":     "10a3",
		"Hung":    "10a6",
	}

	var confrim string
	fmt.Print("Input name: ")
	fmt.Scan(&confrim)

	for key, value := range students {
		if confrim == key {
			fmt.Printf("%s' class: %s", key, value)
			break
		} else {
			fmt.Println("Cannot found!!")
			break
		}
	}

	var confrimV2 string
	fmt.Print("\nInput name delete: ")
	fmt.Scan(&confrimV2)

	for key := range students {
		if confrimV2 == key {
			delete(students, key)
			fmt.Println("Deleted!!!")
			for key, value := range students {
				fmt.Println(key + "|" + value)
			}
			break
		} else {
			fmt.Println("Cannot found!!")
			break
		}
	}

}
