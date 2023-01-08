package model

type Request struct {
	CustomerId string `json:"norek"`
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	CustomerId string  `json:"norek"`
	Saldo      float64 `json:"saldo"`
}

type Rekening struct {
	CustomerId string  `gorm:"column:customer_id;primaryKey"`
	Balance    float64 `gorm:"column:balance"`
}

func (Rekening) TableName() string {
	return "tbl_rekening"
}
