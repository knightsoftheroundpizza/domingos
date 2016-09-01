package main

import (
    "net/http"
    "github.com/boltdb/bolt"
)

type OrdersHandler struct {
	db *bolt.DB
}

func (oh *OrdersHandler) Init() {

}

func (oh *OrdersHandler) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"id\": \"test\"}"))

}

func (oh *OrdersHandler) PostOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"id\": \"test\"}"))
}

func (oh *OrdersHandler) PostOrdersOnIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"id\": \"test\"}"))
}