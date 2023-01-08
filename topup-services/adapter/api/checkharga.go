package api

import (
	"devimaya/goldshop/topupservices/adapter/db/model"
	"devimaya/goldshop/topupservices/internal/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CheckHarga() (interface{}, error) {
	cfg := config.GetConfig()

	url := cfg.CheckHargaEndpoint
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var checkHarga *model.CheckHargaResponse
	err = json.Unmarshal(responseData, &checkHarga)
	if err != nil {
		return "", err
	}

	return checkHarga, nil

}
