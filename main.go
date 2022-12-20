
package main

import (
    "fmt"
    "log"
    "net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "i am working")
    log.Println("i am working")
}

func health(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "OK")
    log.Println("GET /health")
}

func main() {


    http.HandleFunc("/health", health)
    http.HandleFunc("/api", api)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
