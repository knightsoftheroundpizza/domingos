package dominos

import (
	"encoding/json"
	"log"
)

type MenuItem struct {
	Name string
	Code string
}

type Address struct {
	Type         string `json:",omitempty"`
	Street       string `json:",omitempty"`
	City         string `json:",omitempty"`
	Region       string `json:",omitempty"`
	PostalCode   string `json:",omitempty"`
	StreetNumber string `json:",omitempty"`
	StreetName   string `json:",omitempty"`
}

type Product struct {
	Id           int     `json:",omitempty"`
	Code         string  `json:",omitempty"`
	Qty          int     `json:",omitempty"`
	CategoryCode string  `json:",omitempty"`
	Price        float32 `json:",omitempty"`
	Name         string  `json:",omitempty"`
}

type Order struct {
	Id                   string             `json:"OrderID,omitempty"`
	Address              *Address           `json:",omitempty"`
	Email                string             `json:",omitempty"`
	FirstName            string             `json:",omitempty"`
	LastName             string             `json:",omitempty"`
	Phone                string             `json:",omitempty"`
	ServiceMethod        string             `json:",omitempty"`
	StoreId              string             `json:"StoreID,omitempty"`
	EstimatedWaitMinutes string             `json:",omitempty"`
	Products             []Product          `json:",omitempty"`
	Amounts              map[string]float32 `json:",omitempty"`
}

type orderResult struct {
	Order *Order `json:",omitempty"`
}

func ParseOrder(data []byte) *Order {
	result := orderResult{}
	json.Unmarshal(data, &result)
	return result.Order
}

func ParseMenu(data []byte) []MenuItem {
	var i interface{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		log.Printf("Couldn't unmarshal menu - error was %v\n", err)
		return []MenuItem{}
	}
	r := i.(map[string]interface{})
	rs := r["result"]
	items := rs.([]interface{})
	menu := []MenuItem{}
	for _, item := range items {
		switch item.(type) {
		case map[string]interface{}:
			menu = append(menu, *createMenuItemFromItem(item.(map[string]interface{})))
		default:
			log.Printf("Menu incorrectly formatted")
			return []MenuItem{}
		}
	}
	return menu
}

func createMenuItemFromItem(item map[string]interface{}) *MenuItem {
	for k, v := range item {
		// There's only one item in here, but have to use range...
		return &MenuItem{
			Name: k,
			Code: v.(string),
		}
	}
	return nil
}
