package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"net/http"
	"github.com/knightsoftheroundpizza/domingos/dominos"
)

const ORDER_BUCKET = "orders"

type Order struct {
	Id       string `json:id`
	Status   string `json:status`
	Customer string `json:customer`
}

type OrdersDb struct {
	db *bolt.DB
}

func (o *OrdersDb) Init() {
	o.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(ORDER_BUCKET))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func (o *OrdersDb) Find(id string) Order {
	result := Order{}
	o.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ORDER_BUCKET))
		v := b.Get([]byte(id))
		json.Unmarshal(v, &result)
		return nil
	})
	return result
}

func (o *OrdersDb) FindAll() []Order {
	result := []Order{}
	o.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(ORDER_BUCKET))
		b.ForEach(func(k, v []byte) error {
			e := Order{}
			json.Unmarshal(v, &e)
			result = append(result, e)
			return nil
		})
		return nil
	})
	return result
}

func (o *OrdersDb) Remove(id string) bool {
	success := true
	o.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ORDER_BUCKET))
		err := b.Delete([]byte(id))
		success = (err != nil)
		return err
	})
	return success
}

func (o *OrdersDb) Insert(order Order) bool {
	success := true
	o.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ORDER_BUCKET))
		fmt.Println(b)
		m, _ := json.Marshal(order)
		err := b.Put([]byte(order.Id), m)
		success = (err != nil)
		return err
	})
	return success
}

type OrdersHandler struct {
	ordersDb *OrdersDb
}

func CreateOrdersHandler(db *bolt.DB) *OrdersHandler {
	odb := &OrdersDb{db: db}
	odb.Init()
	return &OrdersHandler{ordersDb: odb}
}

func (oh *OrdersHandler) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"id\": \"test\"}"))

}

func (oh *OrdersHandler) PostOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"id\": \"test\"}"))
}

func (oh *OrdersHandler) PriceOrderHandler(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("POST", DominosURL+"/price-order", r.Body)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	order := dominos.ParseOrder(body)

	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(map[string]float32{
		"Net": order.Amounts["Net"],
		"Tax": order.Amounts["Tax"],
		"Total": order.Amounts["Payment"],
	})
	w.Write(result)
}

func (oh *OrdersHandler) PostOrdersOnIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"id\": \"test\"}"))
}
