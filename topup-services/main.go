package main

import (
	"devimaya/goldshop/topupservices/adapter"
	"devimaya/goldshop/topupservices/adapter/api"
	"devimaya/goldshop/topupservices/adapter/db/model"
	"devimaya/goldshop/topupservices/internal/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/teris-io/shortid"
)

func TopupProducer(w http.ResponseWriter, r *http.Request) {
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

	inputHarga, err := strconv.ParseFloat(request.Harga, 64)
	if err != nil {
		newErr := errors.New("gagal mengambil harga")
		returnError(newErr, w)
		return
	}

	if inputHarga != checkHarga.Data.Topup {
		newErr := errors.New("harga tidak sesuai harga topup")
		returnError(newErr, w)
		return
	}

	arr := strings.Split(request.Gram, ".")
	if len(arr) != 2 {
		newErr := errors.New("format gram salah (contoh yg benar : 1.23)")
		returnError(newErr, w)
		return
	}
	if len(arr[1]) > 3 {
		newErr := errors.New("gram bukan kelipatan 0.001")
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
	err = adapter.Produce("topup", id, newReq)
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

	r.HandleFunc("/api/topup", TopupProducer)

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
