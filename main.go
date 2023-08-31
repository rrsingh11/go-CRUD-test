package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"testapi/handlers"
)





func main() {
	port := flag.Int("port", 2969, "Write the port name")
	flag.Parse()
	mux := http.NewServeMux()

	h := testapi.NewHandler()
	mux.Handle("/api/contact", &h)

	fmt.Println("Starting server at :", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}



