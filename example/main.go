package main

import (
	"log"
	"net/http"

	_ "github.com/nogoegst/cabin/magic"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Printf("status: %s", resp.Status)

}
