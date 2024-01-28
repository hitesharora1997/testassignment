package server

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/hitesharora1997/testassignment/internal/counter"
)

func TestServer_ServeHTTP(t *testing.T) {
	s := &Server{
		counter: counter.NewRequestCounter(),
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.ServeHTTP)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
