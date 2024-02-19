package domain

type PurchaseDomain struct {
	Id			    int64	   `copier:"Id"`
	ClientName	    string     `copier:"ClientName"`
	DatePurchase    string     `copier:"DatePurchase"`
	ShippingAddress Address    `copier:"ShippingAddress"`
	Products        []Product  `copier:"Products"`
    TotalAmount     float64    `copier:"TotalAmount"`
    PaymentMethod   string     `copier:"PaymentMethod"`
}

type Address struct {
    Street  string	`copier:"Street"`
    City    string	`copier:"City"`
    State   string	`copier:"State"`
    ZipCode string	`copier:"ZipCode"`
    Country string	`copier:"Country"`
}

type Product struct {
    Id       int64	  `copier:"Id"`
    Name     string	  `copier:"Name"`
    Quantity int	  `copier:"Quantity"`
    Price    float64  `copier:"Price"`
}