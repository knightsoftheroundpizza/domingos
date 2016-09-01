package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type StoresHandler struct {
}

func (sh *StoresHandler) GetStoreMenuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]
	req, _ := http.NewRequest("GET", DominosURL+"/store/"+storeId+"/menu?lang=en&structured=true", nil)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
