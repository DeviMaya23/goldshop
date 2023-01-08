package model

type Request struct {
	AdminId      string `json:"admin_id"`
	HargaTopup   string `json:"harga_topup"`
	HargaBuyback string `json:"harga_buyback"`
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	ReffId  string `json:"reff_id"`
}
