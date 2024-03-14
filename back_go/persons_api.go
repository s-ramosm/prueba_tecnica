package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const api_persons string = "http://localhost:8000/persons/"

type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/api/persons", func(w http.ResponseWriter, r *http.Request) {

		response, err := http.Get(api_persons)
		if err != nil {
			fmt.Println("Error al hacer la solicitud:", err)
			return
		}

		resBody, err := io.ReadAll(response.Body)

		// message := Message{Text: "Hello, world!"}

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			log.Fatal(err)
		}

		w.Write(resBody)

	})

	log.Println("Start Server localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
