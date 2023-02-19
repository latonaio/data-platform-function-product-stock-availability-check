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
	Header            Header   `json:"ProductStock"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	ProductStockCode  string   `json:"product_stock_code"`
	Deleted           bool     `json:"deleted"`
}

type Header struct {
	BusinessPartner           int          `json:"BusinessPartner"`
	Product                   string       `json:"Product"`
	Plant                     string       `json:"Plant"`
	StorageLocation           *string      `json:"StorageLocation"`
	Batch                     *string      `json:"Batch"`
	OrderID                   *int         `json:"OrderID"`
	OrderItem                 *int         `json:"OrderItem"`
	Project                   *string      `json:"Project"`
	InventoryStockType        *string      `json:"InventoryStockType"`
	InventorySpecialStockType *string      `json:"InventorySpecialStockType"`
	ProductBaseUnit           *string      `json:"ProductBaseUnit"`
	ProductStock              *string      `json:"ProductStock"`
	Availability              Availability `json:"Availability"`
}

type Availability struct {
	BatchValidityEndDate         *string  `json:"BatchValidityEndDate"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
