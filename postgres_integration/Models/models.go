package model

type Stock struct{
	StockId int `json:"stockid"`
	Name string `json:"name"`
	Price int `json:"price"`
	Company string `json:"company"`
}

type Response struct{
	ID string `json:"id"`
	Message string `json:"messsage"`
}