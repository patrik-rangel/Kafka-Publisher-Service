package domain

type PurchaseDomain struct {
	Id			    int64
	ClientName	    string
	DatePurchase    string
	ShippingAddress Address
	Products        []Product
    TotalAmount     float64
    PaymentMethod   string
}

type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
    Country string
}

type Product struct {
    ID       int64
    Name     string
    Quantity int
    Price    float64
}