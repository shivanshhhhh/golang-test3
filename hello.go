package main

import (
    "fmt"
    "net/http"
)

// healthzHandler handles the /healthz route.
func healthzHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World")
}

func main() {
    http.HandleFunc("/healthz", healthzHandler)

    fmt.Println("Hello World")
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
