package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/boltdb/bolt"
)

func main() {
    r := mux.NewRouter()

    db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    oh := OrdersHandler{db: db}
    // Routes consist of a path and a handler function.
    r.HandleFunc("/orders", oh.GetOrdersHandler).Methods("GET")
    r.HandleFunc("/orders", oh.PostOrdersHandler).Methods("POST")
    r.HandleFunc("/orders/{id}", oh.PostOrdersOnIdHandler).Methods("POST")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))

}
