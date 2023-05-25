package requests

type ProductStockAvailabilityType struct {
	IsProductStockAvailability                    bool `json:"IsProductStockAvailability"`
	IsProductStockAvailabilityByBatch             bool `json:"IsProductStockAvailabilityByBatch"`
	IsProductStockAvailabilityByStorageBin        bool `json:"IsProductStockAvailabilityByStorageBin"`
	IsProductStockAvailabilityByStorageBinByBatch bool `json:"IsProductStockAvailabilityByStorageBinByBatch"`
}
