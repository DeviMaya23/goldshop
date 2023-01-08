package repository

import (
	"devimaya/goldshop/inputhargastorage/adapter/db"
	"devimaya/goldshop/inputhargastorage/adapter/db/model"
	"fmt"
	"strconv"
)

func SetHarga(req *model.Request) error {
	var harga *model.Harga

	topup, _ := strconv.ParseFloat(req.HargaTopup, 32)
	buyback, _ := strconv.ParseFloat(req.HargaBuyback, 32)

	err := db.GoldShopDB.First(&harga).Error
	if err != nil {
		fmt.Println("error : " + err.Error())
		return err
	}

	harga.Topup = float32(topup)
	harga.Buyback = float32(buyback)
	db.GoldShopDB.Where("id = 1").Save(&harga) // TODO fix query
	return nil
}
