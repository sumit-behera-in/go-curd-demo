package models

type Address struct {
	State   string
	City    string
	Pincode int
}

type User struct {
	Name    string  `json:"name" bson:"userName"`
	Age     int     `json:"age" bson:"userAge"`
	Address Address `json:"address" bson:"userAddress"`
}
