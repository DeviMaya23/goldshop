package main

import (
	"devimaya/goldshop/checksaldoservices/adapter/db"
	"devimaya/goldshop/checksaldoservices/adapter/db/model"
	"devimaya/goldshop/checksaldoservices/internal/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckSaldo(w http.ResponseWriter, r *http.Request) {
	// response := &model.Response{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnError(err, w)
		return
	}

	var request *model.Request
	json.Unmarshal(reqBody, &request)
	if err != nil {
		returnError(err, w)
		return
	}

	rekening := model.Rekening{CustomerId: request.CustomerId}
	err = db.GoldShopDB.First(&rekening).Error
	if err != nil {
		returnError(err, w)
		return
	}

	data := model.Data{
		CustomerId: rekening.CustomerId,
		Saldo:      rekening.Balance,
	}

	response := &model.Response{}
	response.Data = data
	response.Error = "false"

	res, _ := json.Marshal(response)
	fmt.Fprint(w, string(res))
}

func returnError(err error, w http.ResponseWriter) {
	response := &model.Response{}
	response.Error = "true"
	response.Message = err.Error()

	res, _ := json.Marshal(response)
	fmt.Fprint(w, string(res))
}

func main() {
	config := config.GetConfig()

	r := mux.NewRouter()
	r.HandleFunc("/api/saldo", CheckSaldo)

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    config.ServiceHost + ":" + config.ServicePort,
		// WriteTimeout: 10 * time.Second,
		// ReadTimeout:  10 * time.Second,
	}

	fmt.Println("Listening to " + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
