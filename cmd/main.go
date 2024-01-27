package main

import (
	"github.com/hitesharora1997/testassignment/internal/server"
)

func main() {
	dataFile := "server_data.json"
	srv := server.NewServer(dataFile)
	srv.RestoreData()
}
