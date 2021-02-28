package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const url = "localhost:8080"
const endpoint = "/endpoint"

func main() {
	fmt.Printf("pid: %d\n", os.Getpid())

	http.HandleFunc(endpoint, func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("OK"))

		if err != nil {
			log.Fatalln(err)
		}
	})

	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatalln(err)
	}
}
