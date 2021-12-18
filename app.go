package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Loading config")
	err := LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config:", err)
		return
	}

	addr, ok := os.LookupEnv("BIND_ADDR")
	if !ok {
		addr = ":5000"
	}

	log.Println("Starting web server on", addr)
	err = http.ListenAndServe(addr, WebHandler)
	if err != nil {
		log.Fatalln("Unable to listen:", err)
		return
	}
}
