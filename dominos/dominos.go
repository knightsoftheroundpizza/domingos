package dominos

import (
	"log"
	"encoding/json"
)

type MenuItem struct {
	Name string
	Code string
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
