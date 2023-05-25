package api_processing_data_formatter

type SDC struct {
	ProductStockAvailabilityType      *ProductStockAvailabilityType      `json:"ProductStockAvailabilityType"`
	ProductStockAvailability          *ProductStockAvailability          `json:"ProductStockAvailability"`
	ComparisonStockAndQuantity        *ComparisonStockAndQuantity        `json:"ComparisonStockAndQuantity"`
	StockAndQuantity                  *StockAndQuantity                  `json:"StockAndQuantity"`
	RecalculatedAvailableProductStock *RecalculatedAvailableProductStock `json:"RecalculatedAvailableProductStock"`
}

type ProductStockAvailabilityType struct {
	IsProductStockAvailability                    bool `json:"IsProductStockAvailability"`
	IsProductStockAvailabilityByBatch             bool `json:"IsProductStockAvailabilityByBatch"`
	IsProductStockAvailabilityByStorageBin        bool `json:"IsProductStockAvailabilityByStorageBin"`
	IsProductStockAvailabilityByStorageBinByBatch bool `json:"IsProductStockAvailabilityByStorageBinByBatch"`
}

// 1. Product Stock Availability
type ProductStockAvailabilityKey struct {
	Product                      string `json:"Product"`
	BusinessPartner              int    `json:"BusinessPartner"`
	Plant                        string `json:"Plant"`
	StorageLocation              string `json:"StorageLocation"`
	StorageBin                   string `json:"StorageBin"`
	Batch                        string `json:"Batch"`
	ProductStockAvailabilityDate string `json:"ProductStockAvailabilityDate"`
}

type ProductStockAvailability struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Product                      string  `json:"Product"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageBin                   string  `json:"StorageBin"`
	Batch                        string  `json:"Batch"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}

// 2. 利用可能在庫と要求数量の比較
type ComparisonStockAndQuantity struct {
	AvailableProductStock   float32 `json:"AvailableProductStock"`
	RequestedQuantity       float32 `json:"RequestedQuantity"`
	IsAvailableProductStock bool    `json:"IsAvailableProductStock"`
	IsRequestedQuantity     bool    `json:"IsRequestedQuantity"`
}

type StockAndQuantity struct {
	CheckedQuantity                 float32 `json:"CheckedQuantity"`
	CheckedDate                     string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
}

// 3. 利用可能在庫の再計算
type RecalculatedAvailableProductStock struct {
	AvailableProductStock             float32 `json:"AvailableProductStock"`
	CheckedQuantity                   float32 `json:"CheckedQuantity"`
	RecalculatedAvailableProductStock float32 `json:"RecalculatedAvailableProductStock"`
}
