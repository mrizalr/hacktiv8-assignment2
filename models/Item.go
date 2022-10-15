package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"type:varchar(5)" json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
