package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello from Mondoo Engineer!\n")
        })

        log.Println("Server listening on :8080") 
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err) 
        }
}
