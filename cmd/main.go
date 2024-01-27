package main

import (
	"fmt"

	server "github.com/hitesharora1997/testassignment/internal"
)

func main() {
	dataFile := "server_data.json"
	srv := server.NewServer(dataFile)
	fmt.Println(srv)
}
