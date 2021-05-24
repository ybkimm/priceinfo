package main

type PriceInfo struct {
	MallName    string  `json:"mallName"`
	Price       float64 `json:"price"`
	ShippingFee float64 `json:"shippingFee"`
	URL         string  `json:"url"`
}
