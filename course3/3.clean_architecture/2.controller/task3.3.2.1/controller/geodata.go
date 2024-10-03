package controller

type ResponseAddress struct {
	Address Address `json:"address"`
}

type Address struct {
	Road        string `json:"road"`
	Town        string `json:"town"`
	County      string `json:"county"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ResponseAddressGeocode struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
