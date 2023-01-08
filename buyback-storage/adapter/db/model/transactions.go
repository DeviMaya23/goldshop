package model

type Request struct {
	Gram         string  `json:"gram"`
	Harga        string  `json:"harga"`
	Norek        string  `json:"norek"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}

type Buyback struct {
	TransactionId   string  `gorm:"column:transaction_id;primaryKey"`
	CustomerId      string  `gorm:"column:customer_id"`
	TransactionDate int     `gorm:"column:transaction_date"`
	HargaTopup      float64 `gorm:"column:harga_topup"`
	HargaBuyback    float64 `gorm:"column:harga_buyback"`
	Gram            string  `gorm:"column:gram"`
	Saldo           float64 `gorm:"column:saldo"`
}

func (Buyback) TableName() string {
	return "tbl_transaksi"
}

type Rekening struct {
	CustomerId string  `gorm:"column:customer_id;primaryKey"`
	Balance    float64 `gorm:"column:balance"`
}

func (Rekening) TableName() string {
	return "tbl_rekening"
}
