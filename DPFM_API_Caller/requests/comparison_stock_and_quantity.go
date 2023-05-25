package requests

type ComparisonStockAndQuantity struct {
	AvailableProductStock   float32 `json:"AvailableProductStock"`
	RequestedQuantity       float32 `json:"RequestedQuantity"`
	IsAvailableProductStock bool    `json:"IsAvailableProductStock"`
	IsRequestedQuantity     bool    `json:"IsRequestedQuantity"`
}
