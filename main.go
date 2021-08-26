package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("request received")
	delay := r.URL.Query().Get("d")
	if delay == "" {
		delay = "5s"
	}
	d, err := time.ParseDuration(delay)
	if err != nil {
		d, _ = time.ParseDuration("5s")
	}
	log.Printf("Sleeping for %s ...", delay)
	time.Sleep(d)
	log.Printf(".. done")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Slept for %s", delay)))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Printf("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
