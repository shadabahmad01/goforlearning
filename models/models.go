package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	SmallImage  string  `gorm:"column:smallimg" json:"small_img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` //sql.NullFloat64
	ProductName string  `json:"productName" gorm:"column:productname"`
	Description string  `json:"desc"`
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	Name      string  `json:"name"`
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `json:"customer_id" gorm:"column:customer_id"`
	ProductID    int       `json:"product_id" gorm:"column:product_id"`
	Price        float64   `json:"sell_price" gorm:"column:price"`
	PurchaseDate time.Time `json:"purchase_date" gorm:"column:purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
