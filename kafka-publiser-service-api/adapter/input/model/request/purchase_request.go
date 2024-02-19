package request

import "time"

type PurchaseRequest struct {
	Id 				int64			 `json:"id" binding:"required" copier:"Id"` 	
	ClientName      string           `json:"clientName" binding:"required" copier:"ClientName"`  
	DatePurchase    time.Time        `json:"datePurchase" binding:"required" time_format:"2006-01-02" copier:"DatePurchase"`
	Products        []ProductRequest `json:"products" binding:"required" copier:"Products"`
	PaymentMethod   string           `json:"paymentMethod" binding:"required" copier:"PaymentMethod"`
	ShippingAddress AddressRequest   `json:"shippingAddress" binding:"required" copier:"ShippingAddress"`
}

type ProductRequest struct {
	Id       int64   `json:"id" binding:"required" copier:"Id"`
	Name     string  `json:"name" binding:"required" copier:"Name"`
	Quantity int     `json:"quantity" binding:"required" copier:"Quantity"`
	Price    float64 `json:"price" binding:"required" copier:"Price"`
}

type AddressRequest struct {
	Street  string `json:"street" binding:"required" copier:"Street"`
	City    string `json:"city" binding:"required" copier:"City"`
	State   string `json:"state" binding:"required" copier:"State"`
	ZipCode string `json:"zipCode" binding:"required" copier:"ZipCode"`
	Country string `json:"country" binding:"required" copier:"Country"`
}
