package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}

type Product struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	InStock int    `json:"in_stock"`
	Unit    string `json:"unit"`
}

type Food struct {
	Id        int             `json:"id"`
	Products  map[Product]int `json:"products"`
	Price     int             `json:"price"`
	Available bool            `json:"available"`
	Limit     int             `json:"limit"`
}

type FoodCount struct {
	Food     Food
	Quantity int
}

type Order struct {
	Id        int         `json:"id"`
	User      User        `json:"user"`
	FoodArray []FoodCount `json:"foodmap"`
	Bill      int         `json:"bill"`
}
