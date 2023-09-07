package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"testapi/datastore"
	"testapi/handlers"
	"testapi/services"
	"testapi/utils"
)

func main() {

	config, fileErr := utils.LoadConfiguration("config.json")

	if fileErr != nil {
		fmt.Println("Error opening config file")
	}

	port := flag.Int("port", config.Port, "Write the port name")
	flag.Parse()

	var store datastore.ContactBook

	switch config.Memory.Type {
	case "mongodb":
		mongoCollection := utils.ConnectMongoDB(config.Memory.Database, config.Memory.Collection)
		store = &datastore.MongoDB{Collection: mongoCollection}
	case "inmemory":
		store = &datastore.InMemory{Contacts: make(map[string]string)}
	case "redis":
		redisClient := utils.ConnectRedis("localhost", "6379")
		store = &datastore.Redis{Client: *redisClient}
	}

	infoLogger := log.New(os.Stdout, "INFO:", 1)

	// validation service
	vs := services.NewValidationSevice(10)
	//Contact service
	cs := services.NewContactService(store, *vs)
	// Logging service
	lserv := services.NewLoggingService(cs, infoLogger)

	serv := services.NewServices(lserv)

	h := handlers.NewContactHandler(*serv)
	up := handlers.NewUploadHandler(*serv)

	mux := http.NewServeMux()

	mux.Handle("/api/contact", h)
	mux.Handle("/api/contact/upload", up)
	mux.Handle("/api/contact/download", up)

	fmt.Println("Starting server at :", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, *port), mux))
}
