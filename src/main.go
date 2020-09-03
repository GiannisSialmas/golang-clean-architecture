package main

import (
	httpController "application/controllers/http/gin"
	"application/infrastructure/database"
	postgresRepository "application/repositories/postgres"
	"application/services"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

// How this should be
// Repository which implements our own interface with our own methods
// Services which implement our own method
// Controllers dont have a point because the depend on the framework, but if we use the contructor method we can write unit tests for them easier

func main() {
	fmt.Printf("Program Starting")
	db := database.NewDB()
	defer db.Close()

	repository := postgresRepository.NewUserRepository(db)
	serviceLayer := services.NewUserService(repository)
	controllerLayer := httpController.NewControllerLayer(serviceLayer)

	httpHandler := controllerLayer.GetHttpHandler()

	httpServer := newHTTPWebserver(httpHandler)

	log.Fatal(httpServer.ListenAndServe())

}

func newHTTPWebserver(handler http.Handler) *http.Server {

	// Get port from env variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		fmt.Printf("No port given. Defaulting to port %s\n", port)
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("Invalid port", port)
	}
	//Get interface to bind to from env variable
	bindInterface := os.Getenv("BIND_INTERFACE")
	var bindAdress string
	if bindInterface == "" {
		bindAdress = "0.0.0.0"
		fmt.Printf("Defaulting to bindInterface %s\n", bindInterface)
	} else {

		ief, err := net.InterfaceByName(bindInterface)

		if err != nil {
			log.Fatal("Error binding interface:", err)
		}
		addrs, err := ief.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		bindAdress = addrs[0].(*net.IPNet).IP.String()
	}
	listenAdress := fmt.Sprintf("%s:%s", bindAdress, port)

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	return &http.Server{
		Addr:         listenAdress,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
