package requests

type ProductStockAvailabilityKeyBylotto struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	Product                      string `json:"Product"`
	Plant                        string `json:"Plant"`
	Batch                        string `json:"Batch"`
	ProductStockAvailabilityDate string `json:"ProductStockAvailabilityDate"`
}
