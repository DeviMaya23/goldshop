package main

import (
	"devimaya/goldshop/checkhargaservices/adapter/db"
	"devimaya/goldshop/checkhargaservices/adapter/db/model"
	"devimaya/goldshop/checkhargaservices/internal/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckHarga(w http.ResponseWriter, r *http.Request) {
	var harga *model.Harga
	err := db.GoldShopDB.First(&harga).Error
	if err != nil {
		fmt.Println("error")
	}

	response := &model.Response{}
	response.Error = "false"
	response.Data = model.HargaResp{
		Topup:   harga.Topup,
		Buyback: harga.Buyback,
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
	r.HandleFunc("/api/check-harga", CheckHarga)

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
