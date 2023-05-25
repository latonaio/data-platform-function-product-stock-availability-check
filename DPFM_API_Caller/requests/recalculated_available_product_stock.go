package requests

type RecalculatedAvailableProductStock struct {
	AvailableProductStock             float32 `json:"AvailableProductStock"`
	CheckedQuantity                   float32 `json:"CheckedQuantity"`
	RecalculatedAvailableProductStock float32 `json:"RecalculatedAvailableProductStock"`
}
