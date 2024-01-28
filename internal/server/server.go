package server

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/hitesharora1997/testassignment/internal/counter"
	"github.com/hitesharora1997/testassignment/pkg/persistance"
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

func (s *Server) PersistData() {
	persistance.PeriodicSave(s.dataFile, s.counter, 5)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count := s.counter.RecordAndCount()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Total request in the last 60 seconds: " + strconv.Itoa(count)))
}
