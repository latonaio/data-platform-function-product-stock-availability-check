package dpfm_api_input_reader

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	Header            Header   `json:"ProductStockAvailabilityCheck"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	ProductStockCode  string   `json:"product_stock_code"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	Product                         *string  `json:"Product"`
	BusinessPartner                 *int     `json:"BusinessPartner"`
	Plant                           *string  `json:"Plant"`
	StorageLocation                 *string  `json:"StorageLocation"`
	StorageBin                      *string  `json:"StorageBin"`
	Batch                           *string  `json:"Batch"`
	RequestedQuantity               *float32 `json:"RequestedQuantity"`
	ProductStockAvailabilityDate    *string  `json:"ProductStockAvailabilityDate"`
	InventoryStockType              *string  `json:"InventoryStockType"`
	InventorySpecialStockType       *string  `json:"InventorySpecialStockType"`
	AvailableProductStock           *float32 `json:"AvailableProductStock"`
	CheckedQuantity                 *float32 `json:"CheckedQuantity"`
	CheckedDate                     *string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             *bool    `json:"StockIsFullyChecked"`
	Suffix                          *string  `json:"Suffix"`
}
