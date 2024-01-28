package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hitesharora1997/testassignment/internal/server"
)

func main() {
	dataFile := "server_data.json"
	fmt.Println("Starting the server")
	srv := server.NewServer(dataFile)
	srv.RestoreData()
	go srv.PersistData()

	http.Handle("/", srv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
