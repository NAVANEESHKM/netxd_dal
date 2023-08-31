package models

// 👈 SignUpInput struct
type Customer struct {
	Customer_ID int32  `json:"customer_id" bson:"customer_id" `
	First_Name  string `json:"first_name" bson:"first_name" `
	Last_Name   string `json:"last_name" bson:"last_name" `
	Bank_ID     int32  `json:"bank_id" bson:"bank_id" `
	Balance     int32  `json:"balance" bson:"balance"`
	Created_At  string `json:"created_at" bson:"created_at" `
	Updated_At  string `json:"updated_at" bson:"updated_at"`
	ISActive    string `json:"isactive" bson:"isactive" `
}