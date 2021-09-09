package main

import (
	"fmt"
	"log"
	"math/rand"
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

type healthServer struct{}

func (s *healthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("request received")

	// Sleep for a random period
	min := 1
	max := 5000
	delay := rand.Intn(max-min+1) + min
	d, err := time.ParseDuration(fmt.Sprintf("%dms", delay))
	if err != nil {
		d, _ = time.ParseDuration("5s")
	}
	log.Printf("Sleeping for %dms ...", delay)
	time.Sleep(d)

	// Are we going to return OK, or error
	min = 1
	max = 100
	threshold := rand.Intn(max-min+1) + min
	if threshold < 66 {
		log.Printf("Returning 200")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"health": "ok"}`))
	} else {
		log.Printf("Returning 500")
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Server error`))
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	h := &healthServer{}
	http.Handle("/health", h)
	s := &server{}
	http.Handle("/", s)
	log.Printf("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
