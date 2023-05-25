package requests

type ProductStockAvailabilityKey struct {
	Product                      string `json:"Product"`
	BusinessPartner              int    `json:"BusinessPartner"`
	Plant                        string `json:"Plant"`
	StorageLocation              string `json:"StorageLocation"`
	StorageBin                   string `json:"StorageBin"`
	Batch                        string `json:"Batch"`
	ProductStockAvailabilityDate string `json:"ProductStockAvailabilityDate"`
}
