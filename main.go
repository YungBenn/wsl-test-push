package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HelloRuben(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Ruben")
}

func HC(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok!")
}

type resp struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HandleJson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	data := resp{
		Name: "Ruben Adisuryo Nugroho",
		Age:  21,
	}

	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshaling data", http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func main() {
	http.HandleFunc("/json", HandleJson)
	http.HandleFunc("/hello", HelloRuben)
	http.HandleFunc("/", HC)

	port := "8000"
	http.ListenAndServe(":"+port, nil)
	log.Println("Server is running on http://localhost:8000")
}
