package main

import (
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//DominosURL is root URL
const DominosURL = "https://order.dominos.ca/power/"

func main() {
	r := mux.NewRouter()

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//ORDERS
	oh := CreateOrdersHandler(db)
	r.HandleFunc("/orders", oh.GetOrdersHandler).Methods("GET")
	r.HandleFunc("/orders", oh.PostOrdersHandler).Methods("POST")
	r.HandleFunc("/orders/price", oh.PriceOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{id}", oh.PostOrdersOnIdHandler).Methods("POST")

	//STORES
	sh := StoresHandler{}
	r.HandleFunc("/store/{storeId}/menu", sh.GetStoreMenuHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
