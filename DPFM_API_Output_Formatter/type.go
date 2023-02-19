package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey     string      `json:"connection_key"`
	Result            bool        `json:"result"`
	RedisKey          string      `json:"redis_key"`
	Filepath          string      `json:"filepath"`
	APIStatusCode     int         `json:"api_status_code"`
	RuntimeSessionID  string      `json:"runtime_session_id"`
	BusinessPartnerID *int        `json:"business_partner"`
	ServiceLabel      string      `json:"service_label"`
	Message           interface{} `json:"message"`
	APISchema         string      `json:"api_schema"`
	Accepter          []string    `json:"accepter"`
	ProductStockCode  string      `json:"product_stock_code"`
	Deleted           bool        `json:"deleted"`
}

type Message struct {
	ProductStockAvailability *ProductStockAvailability `json:"ProductStockAvailability"`
}

type ProductStockAvailability struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Product                      string  `json:"Product"`
	Plant                        string  `json:"Plant"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}
