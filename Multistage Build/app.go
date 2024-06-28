package main

import (
    "fmt"
    "net/http"
    "os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var name, = os.Hostname()

	fmt.Fprintf(w, "<h1>Hello! This request was processed by host: %s</h1>\n", name)
}

func main() {
    http.HandleFunc("/", Handler)

    port := "8080"
    if envPort := os.Getenv("PORT"); envPort != "" {
        port = envPort
    }

    fmt.Printf("Server starting on port %s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
        os.Exit(1)
    }
}
