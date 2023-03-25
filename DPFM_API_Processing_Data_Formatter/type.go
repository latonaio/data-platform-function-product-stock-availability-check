package api_processing_data_formatter

type SDC struct {
	ProductStockAvailability          *ProductStockAvailability          `json:"ProductStockAvailability"`
	ComparisonStock                   *ComparisonStock                   `json:"ComparisonStock"`
	RecalculatedAvailableProductStock *RecalculatedAvailableProductStock `json:"RecalculatedAvailableProductStock"`
}

// 1-1. Product Stock Availability
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
	Batch                        *string `json:"Batch"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}

// 1-2
type ProductStockAvailabilityKeyBylotto struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	Product                      string `json:"Product"`
	Plant                        string `json:"Plant"`
	Batch                        string `json:"Batch"`
	ProductStockAvailabilityDate string `json:"ProductStockAvailabilityDate"`
}

// 2
type ComparisonStock struct {
	CheckedQuantity                 float32 `json:"CheckedQuantity"`
	CheckedDate                     string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
}

// 3
type RecalculatedAvailableProductStock struct {
	RecalculatedAvailableProductStock float32 `json:"RecalculatedAvailableProductStock"`
}
