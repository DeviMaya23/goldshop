package model

type Request struct {
	AdminId      string `json:"admin_id"`
	HargaTopup   string `json:"harga_topup"`
	HargaBuyback string `json:"harga_buyback"`
}

type Harga struct {
	Topup   float32 `gorm:"column:harga_topup"`
	Buyback float32 `gorm:"column:harga_buyback"`
}

func (Harga) TableName() string {
	return "tbl_harga"
}
