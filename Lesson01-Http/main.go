package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is starting...")

	//create server
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Chưa có cái gì đâu", err)

	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	if r.Method != http.MethodGet {
		http.Error(w, "Not support", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"message": "Welcome Golang and Gin",
		"info":    "Kohi pro Golang",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Course", "Coding with Golang")

	// data, err := json.Marshal(response)

	// if err != nil {
	// 	http.Error(w, "Error with change Json", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(data)

	//Cach khac

	json.NewEncoder(w).Encode(response)
}
