package model

type Request struct {
	CustomerId string `json:"customer_id"`
	StartDate  int    `json:"start_date"`
	EndDate    int    `json:"end_date"`
}

type Response struct {
	Error   string        `json:"error"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

type Transaction struct {
	Date         int    `json:"date"`
	Type         string `json:"type"`
	Gram         string `json:"gram"`
	HargaTopup   string `json:"harga_topup"`
	HargaBuyback string `json:"harga_buyback"`
	Saldo        string `json:"saldo"`
}

type Buyback struct {
	TransactionId   string `gorm:"column:transaction_id;primaryKey"`
	CustomerId      string `gorm:"column:customer_id"`
	TransactionDate int    `gorm:"column:transaction_date"`
	HargaTopup      string `gorm:"column:harga_topup"`
	HargaBuyback    string `gorm:"column:harga_buyback"`
	Gram            string `gorm:"column:gram"`
	Saldo           string `gorm:"column:saldo"`
}

func (Buyback) TableName() string {
	return "tbl_transaksi"
}

type Topup struct {
	TopupId         string `gorm:"column:topup_id;primaryKey"`
	CustomerId      string `gorm:"column:customer_id"`
	TransactionDate int    `gorm:"column:transaction_date"`
	HargaTopup      string `gorm:"column:harga_topup"`
	HargaBuyback    string `gorm:"column:harga_buyback"`
	Gram            string `gorm:"column:gram"`
	Saldo           string `gorm:"column:saldo"`
}

func (Topup) TableName() string {
	return "tbl_topup"
}
