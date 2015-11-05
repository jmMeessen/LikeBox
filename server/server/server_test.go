package server

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"log"
)

func TestUnauthorized(t *testing.T) {

	server := NewAPIServer(StubBackend{})
	ts := httptest.NewServer(server)
	defer ts.Close()

	res, err := http.Get(ts.URL+"/status/1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Status code "+res.Status)
	if res.StatusCode != http.StatusUnauthorized {
		log.Fatal("Should have been rejected as unahtuorized")
	}
}


type StubBackend struct {
}

func (s StubBackend) getStatus(box string) (Status, error) {
	return Status{12345, 67890}, nil
}

func (s StubBackend) vote(box string, session string, timestamp int64, vote int) error {
	return nil
}