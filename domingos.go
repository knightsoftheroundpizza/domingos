package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/boltdb/bolt"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)
    db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
