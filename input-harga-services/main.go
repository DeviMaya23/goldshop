package main

import (
	"devimaya/goldshop/inputhargaservices/adapter"
	"devimaya/goldshop/inputhargaservices/adapter/db/model"
	"devimaya/goldshop/inputhargaservices/internal/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
)

func InputHargaProducer(w http.ResponseWriter, r *http.Request) {
	id, _ := shortid.Generate()

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

	err = adapter.Produce("input-harga", id, reqBody)
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

	r.HandleFunc("/api/input-harga", InputHargaProducer)

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
