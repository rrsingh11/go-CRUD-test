package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"testapi/datastore"
	"testapi/handlers"
)


func main() {
	port := flag.Int("port", 2969, "Write the port name")
	flag.Parse()
	mux := http.NewServeMux()

	h := handlers.NewContactHandler(&datastore.InMemory{ Contacts: make(map[string]string) })
	mux.Handle("/api/contact", h)

	fmt.Println("Starting server at :", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}



