package model

type RequestMessage struct {
	Gram         string  `json:"gram"`
	Harga        string  `json:"harga"`
	Norek        string  `json:"norek"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	ReffId  string `json:"reff_id"`
}

type CheckSaldo struct {
	CustomerId string `json:"norek"`
}

type CheckSaldoResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	CustomerId string  `json:"norek"`
	Saldo      float64 `json:"saldo"`
}

type CheckHarga struct {
	AdminId      string  `json:"admin_id"`
	HargaTopup   float64 `json:"harga_topup"`
	HargaBuyback float64 `json:"harga_buyback"`
}

type CheckHargaResponse struct {
	Error   string    `json:"error"`
	Message string    `json:"message"`
	Data    HargaResp `json:"data"`
}

type HargaResp struct {
	Topup   float64 `json:"harga_topup"`
	Buyback float64 `json:"harga_buyback"`
}
