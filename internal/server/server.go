package server

import (
	"log"
	"sync"

	"github.com/hitesharora1997/testassignment/internal/counter"
	"github.com/hitesharora1997/testassignment/internal/persistance"
)

type Server struct {
	sync.Mutex
	dataFile string
	counter  *counter.RequestCounter
}

func NewServer(dataFile string) *Server {
	return &Server{
		dataFile: dataFile,
		counter:  counter.NewRequestCounter(),
	}
}

func (s *Server) RestoreData() {
	if err := persistance.Restore(s.dataFile, s.counter); err != nil {
		log.Println("Error restoring the data", err)
	}
}
