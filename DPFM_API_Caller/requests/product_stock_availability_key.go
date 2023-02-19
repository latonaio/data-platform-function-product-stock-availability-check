package requests

type ProductStockAvailabilityKey struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	Product                      string `json:"Product"`
	Plant                        string `json:"Plant"`
	ProductStockAvailabilityDate string `json:"ProductStockAvailabilityDate"`
}
