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

func (f *Function) ProductStockAvailabilityBylotto(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailability, error) {

	dataKey := psdc.ConvertToProductStockAvailabilityKeyBylotto(sdc)

	rows, err := f.db.Query(
		`SELECT BusinessPartner, Product, Plant, Batch, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_data
		WHERE (BusinessPartner, Product, Plant, Batch, ProductStockAvailabilityDate) = (?, ?, ?, ?, ?);`, dataKey.BusinessPartner, dataKey.Product, dataKey.Plant, dataKey.Batch, dataKey.ProductStockAvailabilityDate,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToProductStockAvailabilityBylotto(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *Function) ComparisonAvailableStock(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.ComparisonStock {

	data := psdc.ConvertToComparisonAvailableStock(sdc)

	return data
}

func (f *Function) ComparisonRequestedStock(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.ComparisonStock {

	difference := psdc.ProductStockAvailability.AvailableProductStock - *sdc.Header.RequestedQuantity

	data := psdc.ConvertToComparisonRequestedStock(sdc, difference)

	return data
}

func (f *Function) RecalculatedAvailableProductStock(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.RecalculatedAvailableProductStock {

	difference := psdc.ProductStockAvailability.AvailableProductStock - psdc.ComparisonStock.CheckedQuantity

	data := psdc.ConvertToRecalculatedAvailableProductStock(sdc, difference)

	return data
}

func getSystemDate() string {
	day := time.Now()
	return day.Format("2006-01-02")
}
