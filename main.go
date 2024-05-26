package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Valid bool `json:"valid"`
}

func creditCardValidator(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var cardNumber struct {
		Number string `json:"number"`
	}
	err := json.NewDecoder(request.Body).Decode(&cardNumber)
	if err != nil {
		http.Error(writer, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	isValid := luhnAlgorithm(cardNumber.Number)
	response := Response{Valid: isValid}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(writer, "Error creating response", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonResponse)
}

func main() {
	args := os.Args
	port := args[1]

	http.HandleFunc("/", creditCardValidator)
	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
