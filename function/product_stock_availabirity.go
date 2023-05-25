package function

import (
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"

	"golang.org/x/xerrors"
)

// 1. Product Stock Availability
func (f *Function) ProductStockAvailability(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailability, error) {

	dataKey := psdc.ConvertToProductStockAvailabilityKey(sdc)

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_availability_data
		WHERE (Product, BusinessPartner, Plant, ProductStockAvailabilityDate) = (?, ?, ?, ?);`, dataKey.Product, dataKey.BusinessPartner, dataKey.Plant, dataKey.ProductStockAvailabilityDate,
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

func (f *Function) ProductStockAvailabilityByBatch(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailability, error) {

	dataKey := psdc.ConvertToProductStockAvailabilityByBatchKey(sdc)

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, Batch, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_avail_by_btch
		WHERE (Product, BusinessPartner, Plant, Batch, ProductStockAvailabilityDate) = (?, ?, ?, ?, ?);`, dataKey.Product, dataKey.BusinessPartner, dataKey.Plant, dataKey.Batch, dataKey.ProductStockAvailabilityDate,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToProductStockAvailabilityByBatch(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *Function) ProductStockAvailabilityByStorageBin(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailability, error) {

	dataKey := psdc.ConvertToProductStockAvailabilityByStorageBinKey(sdc)

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, StorageLocation, StorageBin, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_avail_by_strg_bin
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin, ProductStockAvailabilityDate) = (?, ?, ?, ?, ?, ?);`, dataKey.Product, dataKey.BusinessPartner, dataKey.Plant, dataKey.StorageLocation, dataKey.StorageBin, dataKey.ProductStockAvailabilityDate,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToProductStockAvailabilityByStorageBin(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *Function) ProductStockAvailabilityByStorageBinByBatch(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailability, error) {

	dataKey := psdc.ConvertToProductStockAvailabilityByStorageBinByBatchKey(sdc)

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, StorageLocation, StorageBin, Batch, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_avail_by_strg_bin_btch
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin, Batch, ProductStockAvailabilityDate) = (?, ?, ?, ?, ?, ?, ?);`, dataKey.Product, dataKey.BusinessPartner, dataKey.Plant, dataKey.StorageLocation, dataKey.StorageBin, dataKey.Batch, dataKey.ProductStockAvailabilityDate,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToProductStockAvailabilityByStorageBinByBatch(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

// 2. 利用可能在庫と要求数量の比較
func (f *Function) ComparisonStockAndQuantity(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ComparisonStockAndQuantity, error) {
	availableProductStock := psdc.ProductStockAvailability.AvailableProductStock

	if sdc.Header.RequestedQuantity == nil {
		return nil, xerrors.New("入力ファイルのRequestedQuantityがnullです。")
	}
	requestedQuantity := *sdc.Header.RequestedQuantity

	isAvailableProductStock := false
	isRequestedQuantity := false

	if availableProductStock >= requestedQuantity {
		isAvailableProductStock = true
	} else {
		isRequestedQuantity = true
	}

	data := psdc.ConvertToComparisonStockAndQuantity(availableProductStock, requestedQuantity, isAvailableProductStock, isRequestedQuantity)

	return data, nil
}

func (f *Function) AvailableProductStockProcess(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.StockAndQuantity {
	checkedQuantity := *sdc.Header.RequestedQuantity
	checkedDate := psdc.ProductStockAvailability.ProductStockAvailabilityDate
	openConfirmedQuantityInBaseUnit := float32(0)
	stockIsFullyChecked := true

	data := psdc.ConvertToStockAndQuantity(checkedQuantity, checkedDate, openConfirmedQuantityInBaseUnit, stockIsFullyChecked)

	return data
}

// 3. 利用可能在庫の再計算
func (f *Function) RequestedQuantityProcess(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.StockAndQuantity {
	checkedQuantity := psdc.ProductStockAvailability.AvailableProductStock
	checkedDate := psdc.ProductStockAvailability.ProductStockAvailabilityDate
	openConfirmedQuantityInBaseUnit := psdc.ProductStockAvailability.AvailableProductStock - *sdc.Header.RequestedQuantity
	stockIsFullyChecked := false

	data := psdc.ConvertToStockAndQuantity(checkedQuantity, checkedDate, openConfirmedQuantityInBaseUnit, stockIsFullyChecked)

	return data
}

func (f *Function) RecalculatedAvailableProductStock(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) *dpfm_api_processing_data_formatter.RecalculatedAvailableProductStock {
	availableProductStock := psdc.ProductStockAvailability.AvailableProductStock
	checkedQuantity := psdc.StockAndQuantity.CheckedQuantity
	recalculatedAvailableProductStock := availableProductStock - checkedQuantity

	data := psdc.ConvertToRecalculatedAvailableProductStock(availableProductStock, checkedQuantity, recalculatedAvailableProductStock)

	return data
}
