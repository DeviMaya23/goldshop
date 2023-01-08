package api

import (
	"bytes"
	"devimaya/goldshop/buybackservices/adapter/db/model"
	"devimaya/goldshop/buybackservices/internal/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CheckSaldo(req *model.RequestMessage) (interface{}, error) {
	cfg := config.GetConfig()

	requestSaldo := model.CheckSaldo{CustomerId: req.Norek}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(requestSaldo)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	url := cfg.CheckSaldoEndpoint
	client := &http.Client{}
	reqApi, err := http.NewRequest("GET", url, &buf)
	if err != nil {
		return "", err
	}

	response, err := client.Do(reqApi)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var checkHarga *model.CheckSaldoResponse
	err = json.Unmarshal(responseData, &checkHarga)
	if err != nil {
		return "", err
	}

	return checkHarga, nil

}
