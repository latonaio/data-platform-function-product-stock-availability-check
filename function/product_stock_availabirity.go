package function

import (
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
	"time"
)

func (f *Function) ProductStockAvailability(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailability, error) {

	dataKey := psdc.ConvertToProductStockAvailabilityKey(sdc)

	rows, err := f.db.Query(
		`SELECT BusinessPartner, Product, Plant, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_data
		WHERE (BusinessPartner, Product, Plant, ProductStockAvailabilityDate) = (?, ?, ?, ?);`, dataKey.BusinessPartner, dataKey.Product, dataKey.Plant, dataKey.ProductStockAvailabilityDate,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToProductStockAvailability(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func getSystemDate() string {
	day := time.Now()
	return day.Format("2006-01-02")
}
