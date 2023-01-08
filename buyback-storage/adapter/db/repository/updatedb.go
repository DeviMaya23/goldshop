package repository

import (
	"devimaya/goldshop/buybackstorage/adapter/db"
	"devimaya/goldshop/buybackstorage/adapter/db/model"
	"strconv"
	"time"
)

func InsertBuyback(req *model.Request, balance float64, trxKey string) error {

	now := time.Now()

	date := int(now.Unix())

	newBuyback := model.Buyback{
		TransactionId:   trxKey,
		CustomerId:      req.Norek,
		HargaTopup:      req.HargaTopup,
		HargaBuyback:    req.HargaBuyback,
		Gram:            req.Gram,
		Saldo:           balance,
		TransactionDate: date,
	}

	err := db.GoldShopDB.Create(&newBuyback).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateSaldo(req *model.Request) (float64, error) {
	rekening := model.Rekening{CustomerId: req.Norek}
	err := db.GoldShopDB.First(&rekening).Error
	if err != nil {
		return 0.00, err
	}
	gram, err := strconv.ParseFloat(req.Gram, 64)
	if err != nil {
		return 0.00, err
	}
	rekening.Balance = rekening.Balance - gram

	err = db.GoldShopDB.Save(&rekening).Error
	if err != nil {
		return 0.00, err
	}

	return rekening.Balance, nil
}
