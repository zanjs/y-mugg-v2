package models

type (
	// Inventory is 库存
	Inventory struct {
		BaseModel
		WareroomID int `json:"wareroom_id"`
		ProductID  int `json:"product_id"`
		Quantity   int `json:"quantity"`
	}
)
