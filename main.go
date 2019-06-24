package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port = getEnv("PORT", ":3000")
)

func getEnv(key, defaultValue string) string {
	if result := os.Getenv(key); result != "" {
		return result
	}

	return defaultValue
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Something happens...")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world! I am changing from here!"))
	})

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}

	fmt.Println("Server is starting at", time.Now())
}
