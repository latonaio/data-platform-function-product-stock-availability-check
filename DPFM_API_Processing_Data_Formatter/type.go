package api_processing_data_formatter

type SDC struct {
	ProductStockAvailability *ProductStockAvailability `json:"ProductStockAvailability"`
}

// 1. Product Stock Availability
type ProductStockAvailabilityKey struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	Product                      string `json:"Product"`
	Plant                        string `json:"Plant"`
	ProductStockAvailabilityDate string `json:"ProductStockAvailabilityDate"`
}

type ProductStockAvailability struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Product                      string  `json:"Product"`
	Plant                        string  `json:"Plant"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}
