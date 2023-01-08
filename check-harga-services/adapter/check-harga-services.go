package adapter

import (
	"devimaya/goldshop/checkhargaservices/adapter/db"
	"devimaya/goldshop/checkhargaservices/adapter/db/model"
	"encoding/json"
	"fmt"
)

func GetHarga() []byte {
	var harga *model.Harga
	err := db.GoldShopDB.First(&harga).Error
	if err != nil {
		fmt.Println("error")
	}

	res, _ := json.Marshal(harga)

	return res
}
