package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwzyz"

	varMinWait  = "ABC_MIN_WAIT"
	varMaxWait  = "ABC_MAX_WAIT"
	varEchoURL  = "ABC_URL"
	varEchoPort = "ABC_PORT"
	varHost     = "ABC_HOST"

	paramAlphabet = "alphabet"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		minWait = getDurVar(varMinWait)
		maxWait = getDurVar(varMaxWait)
		port    = uint(getIntVar(varEchoPort))
		host    = getVar(varHost)
	)

	log.Println("Starting server")
	go startServer(host, port)

	for {
		// wait min-max range and start the chain again
		r := rand.Intn(int(maxWait)-int(minWait)) + int(minWait)
		time.Sleep(time.Duration(r))

		log.Println("Starting alphabet")
		startAlphabet(host, port)
	}
}

func startServer(host string, port uint) {
	http.HandleFunc("/continue", func(w http.ResponseWriter, r *http.Request) {
		prev := r.URL.Query().Get(paramAlphabet)
		if prev == "" {
			log.Fatal("url form does not contain ?alphabet=")
		}

		if prev == alphabet {
			log.Println("Done.")
			return
		}

		cont := prev + string(alphabet[len(prev)])

		go makeRequest(host, port, cont)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func startAlphabet(host string, port uint) {
	makeRequest(host, port, "a")
}

func makeRequest(host string, port uint, s string) {
	url := fmt.Sprintf(
		"http://%s:%d/continue?%v=%s", host, port, paramAlphabet, s,
	)

	_, err := http.Get(url)
	if err != nil {
		log.Println(":(", err)
	}
}

func getDurVar(v string) time.Duration {
	val := getVar(v)

	dur, err := time.ParseDuration(val)
	if err != nil {
		log.Fatal("could not parse value as duration", dur)
	}

	return dur
}

func getIntVar(v string) int {
	val := getVar(v)

	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Can't convert %q to int", val)
	}

	return i
}

func getVar(v string) string {
	val := os.Getenv(v)
	if val == "" {
		log.Fatalf("Required var %q is missing", v)
	}

	return val
}
