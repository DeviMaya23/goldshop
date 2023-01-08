package main

import (
	"devimaya/goldshop/buybackservices/adapter"
	"devimaya/goldshop/buybackservices/adapter/api"
	"devimaya/goldshop/buybackservices/adapter/db/model"
	"devimaya/goldshop/buybackservices/internal/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/teris-io/shortid"
)

func BuybackProducer(w http.ResponseWriter, r *http.Request) {
	id, _ := shortid.Generate()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnError(err, w)
		return
	}

	var request *model.RequestMessage
	json.Unmarshal(reqBody, &request)
	if err != nil {
		returnError(err, w)
		return
	}

	var checkHarga *model.CheckHargaResponse
	cekHarga, err := api.CheckHarga()
	if err != nil {
		returnError(err, w)
		return
	}

	err = mapstructure.Decode(cekHarga, &checkHarga)
	if err != nil {
		returnError(err, w)
		return
	}

	var checkSaldo *model.CheckSaldoResponse
	cekSaldo, err := api.CheckSaldo(request)
	if err != nil {
		returnError(err, w)
		return
	}

	err = mapstructure.Decode(cekSaldo, &checkSaldo)
	if err != nil {
		returnError(err, w)
		return
	}

	if checkSaldo.Error == "true" {
		newErr := errors.New(checkSaldo.Message)
		returnError(newErr, w)
		return
	}

	gram, err := strconv.ParseFloat(request.Gram, 64)
	if err != nil {
		newErr := errors.New("gagal convert gram")
		returnError(newErr, w)
		return
	}

	if gram > checkSaldo.Data.Saldo {
		newErr := errors.New("saldo tidak cukup")
		returnError(newErr, w)
		return
	}

	request.HargaBuyback = checkHarga.Data.Buyback
	request.HargaTopup = checkHarga.Data.Topup
	newReq, err := json.Marshal(request)
	if err != nil {
		returnError(err, w)
		return
	}
	err = adapter.Produce("buyback", id, newReq)
	if err != nil {
		returnError(err, w)
		return
	}
	response := model.Response{
		Error:  "false",
		ReffId: id,
	}

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

	r.HandleFunc("/api/buyback", BuybackProducer)

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
