package models

type (
	// Sale is 销售
	Sale struct {
		BaseModel
		WareroomID int      `json:"wareroom_id"`
		ProductID  int      `json:"product_id"`
		Quantity   int      `json:"quantity"`
		Product    Product  `json:"product"`
		Wareroom   Wareroom `json:"wareroom"`
	}
)
