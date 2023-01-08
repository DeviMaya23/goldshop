package main

import (
	"devimaya/goldshop/checkmutasiservices/adapter/db"
	"devimaya/goldshop/checkmutasiservices/adapter/db/model"
	"devimaya/goldshop/checkmutasiservices/internal/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CheckMutasi(w http.ResponseWriter, r *http.Request) {
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

	var buybackList []*model.Buyback
	err = db.GoldShopDB.Where("transaction_date BETWEEN ? AND ?", request.StartDate, request.EndDate).Find(&buybackList, model.Buyback{CustomerId: request.CustomerId}).Debug().Error
	if err != nil {
		returnError(err, w)
		return
	}
	var topupList []*model.Topup
	err = db.GoldShopDB.Where("transaction_date BETWEEN ? AND ?", request.StartDate, request.EndDate).Find(&topupList, model.Topup{CustomerId: request.CustomerId}).Debug().Error
	if err != nil {
		returnError(err, w)
		return
	}

	dataList := []model.Transaction{}
	for _, buyback := range buybackList {
		dataList = append(dataList, model.Transaction{
			Date:         buyback.TransactionDate,
			Type:         "buyback",
			Gram:         buyback.Gram,
			HargaTopup:   buyback.HargaTopup,
			HargaBuyback: buyback.HargaBuyback,
			Saldo:        buyback.Saldo,
		})
	}
	for _, topup := range topupList {
		dataList = append(dataList, model.Transaction{
			Date:         topup.TransactionDate,
			Type:         "topup",
			Gram:         topup.Gram,
			HargaTopup:   topup.HargaTopup,
			HargaBuyback: topup.HargaBuyback,
			Saldo:        topup.Saldo,
		})
	}

	response := &model.Response{}
	response.Data = dataList
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
	r.HandleFunc("/api/mutasi", CheckMutasi)

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
