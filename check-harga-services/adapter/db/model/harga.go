package model

type Response struct {
	Error   string    `json:"error"`
	Message string    `json:"message"`
	Data    HargaResp `json:"data"`
}

type HargaResp struct {
	Topup   float64 `json:"harga_topup"`
	Buyback float64 `json:"harga_buyback"`
}

type Harga struct {
	Topup   float64 `gorm:"column:harga_topup"`
	Buyback float64 `gorm:"column:harga_buyback"`
}

func (Harga) TableName() string {
	return "tbl_harga"
}
