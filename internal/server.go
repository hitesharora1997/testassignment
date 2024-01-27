package server

import "sync"

type Server struct {
	sync.Mutex
	dataFile string
}

func NewServer(dataFile string) *Server {
	return &Server{
		dataFile: dataFile,
	}
}
