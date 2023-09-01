package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"testapi/datastore"
	"testapi/handlers"
	"testapi/utils"
)

func main() {
	port := flag.Int("port", 2969, "Write the port name")
	flag.Parse()

	// Mongo Connection

	mongoCollection := utils.ConnectMongoDB("contact_book", "contacts")

	mux := http.NewServeMux()


	// h := handlers.NewContactHandler(&datastore.InMemory{ Contacts: make(map[string]string) })
	h := handlers.NewContactHandler(&datastore.MongoDB{Collection: mongoCollection})

	mux.Handle("/api/contact", h)

	fmt.Println("Starting server at :", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}



