package main

import (
	"log"
)

func main() {
	log.Printf("Fetching O365Endpoints")
	eps, err := getO365Endpoints()
	if err != nil {
		log.Fatal("Failed to get O365Endpoints: ", err)
	}
	log.Printf("EPS: %+v", eps)
}