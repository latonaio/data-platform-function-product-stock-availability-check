package api_processing_data_formatter

import (
	"data-platform-function-product-stock-availability-check/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	"database/sql"

	"golang.org/x/xerrors"
)

func (psdc *SDC) ConvertToProductStockAvailabilityType(isProductStockAvailability, isProductStockAvailabilityByBatch, isProductStockAvailabilityByStorageBin, isProductStockAvailabilityByStorageBinByBatch bool) *ProductStockAvailabilityType {
	pm := &requests.ProductStockAvailabilityType{}

	pm.IsProductStockAvailability = isProductStockAvailability
	pm.IsProductStockAvailabilityByBatch = isProductStockAvailabilityByBatch
	pm.IsProductStockAvailabilityByStorageBin = isProductStockAvailabilityByStorageBin
	pm.IsProductStockAvailabilityByStorageBinByBatch = isProductStockAvailabilityByStorageBinByBatch

	data := pm
	res := ProductStockAvailabilityType{
		IsProductStockAvailability:                    data.IsProductStockAvailability,
		IsProductStockAvailabilityByBatch:             data.IsProductStockAvailabilityByBatch,
		IsProductStockAvailabilityByStorageBin:        data.IsProductStockAvailabilityByStorageBin,
		IsProductStockAvailabilityByStorageBinByBatch: data.IsProductStockAvailabilityByStorageBinByBatch,
	}

	return &res
}

// 1. Product Stock Availability
func (psdc *SDC) ConvertToProductStockAvailabilityKey(sdc *dpfm_api_input_reader.SDC) *ProductStockAvailabilityKey {
	pm := &requests.ProductStockAvailabilityKey{
		BusinessPartner:              *sdc.Header.BusinessPartner,
		Product:                      *sdc.Header.Product,
		Plant:                        *sdc.Header.Plant,
		ProductStockAvailabilityDate: *sdc.Header.ProductStockAvailabilityDate,
	}

	data := pm
	res := ProductStockAvailabilityKey{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
	}

	return &res
}

func (psdc *SDC) ConvertToProductStockAvailability(rows *sql.Rows) (*ProductStockAvailability, error) {
	defer rows.Close()
	pm := &requests.ProductStockAvailability{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_stock_product_stock_availability_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := &ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductStockAvailabilityByBatchKey(sdc *dpfm_api_input_reader.SDC) *ProductStockAvailability {
	pm := &requests.ProductStockAvailability{
		Product:                      *sdc.Header.Product,
		BusinessPartner:              *sdc.Header.BusinessPartner,
		Plant:                        *sdc.Header.Plant,
		Batch:                        *sdc.Header.Batch,
		ProductStockAvailabilityDate: *sdc.Header.ProductStockAvailabilityDate,
	}

	data := pm
	res := ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
	}

	return &res
}

func (psdc *SDC) ConvertToProductStockAvailabilityByBatch(rows *sql.Rows) (*ProductStockAvailability, error) {
	defer rows.Close()
	pm := &requests.ProductStockAvailability{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.Batch,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_stock_product_stock_availability_by_batch_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := &ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductStockAvailabilityByStorageBinKey(sdc *dpfm_api_input_reader.SDC) *ProductStockAvailability {
	pm := &requests.ProductStockAvailability{
		Product:                      *sdc.Header.Product,
		BusinessPartner:              *sdc.Header.BusinessPartner,
		Plant:                        *sdc.Header.Plant,
		StorageLocation:              *sdc.Header.StorageLocation,
		StorageBin:                   *sdc.Header.StorageBin,
		ProductStockAvailabilityDate: *sdc.Header.ProductStockAvailabilityDate,
	}

	data := pm
	res := ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		StorageLocation:              data.StorageLocation,
		StorageBin:                   data.StorageBin,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
	}

	return &res
}

func (psdc *SDC) ConvertToProductStockAvailabilityByStorageBin(rows *sql.Rows) (*ProductStockAvailability, error) {
	defer rows.Close()
	pm := &requests.ProductStockAvailability{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageBin,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_stock_product_stock_availability_by_storage_bin_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := &ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		StorageLocation:              data.StorageLocation,
		StorageBin:                   data.StorageBin,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductStockAvailabilityByStorageBinByBatchKey(sdc *dpfm_api_input_reader.SDC) *ProductStockAvailability {
	pm := &requests.ProductStockAvailability{
		Product:                      *sdc.Header.Product,
		BusinessPartner:              *sdc.Header.BusinessPartner,
		Plant:                        *sdc.Header.Plant,
		StorageLocation:              *sdc.Header.StorageLocation,
		StorageBin:                   *sdc.Header.StorageBin,
		Batch:                        *sdc.Header.Batch,
		ProductStockAvailabilityDate: *sdc.Header.ProductStockAvailabilityDate,
	}

	data := pm
	res := ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		StorageLocation:              data.StorageLocation,
		StorageBin:                   data.StorageBin,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
	}

	return &res
}

func (psdc *SDC) ConvertToProductStockAvailabilityByStorageBinByBatch(rows *sql.Rows) (*ProductStockAvailability, error) {
	defer rows.Close()
	pm := &requests.ProductStockAvailability{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageBin,
			&pm.Batch,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_stock_product_stock_availability_by_storage_bin_by_batch_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := &ProductStockAvailability{
		Product:                      data.Product,
		BusinessPartner:              data.BusinessPartner,
		Plant:                        data.Plant,
		StorageLocation:              data.StorageLocation,
		StorageBin:                   data.StorageBin,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res, nil
}

// 2. 利用可能在庫と要求数量の比較
func (psdc *SDC) ConvertToComparisonStockAndQuantity(availableProductStock, requestedQuantity float32, isAvailableProductStock, isRequestedQuantity bool) *ComparisonStockAndQuantity {
	pm := &requests.ComparisonStockAndQuantity{}

	pm.AvailableProductStock = availableProductStock
	pm.RequestedQuantity = requestedQuantity
	pm.IsAvailableProductStock = isAvailableProductStock
	pm.IsRequestedQuantity = isRequestedQuantity

	data := pm
	res := ComparisonStockAndQuantity{
		AvailableProductStock:   data.AvailableProductStock,
		RequestedQuantity:       data.RequestedQuantity,
		IsAvailableProductStock: data.IsAvailableProductStock,
		IsRequestedQuantity:     data.IsRequestedQuantity,
	}

	return &res
}

func (psdc *SDC) ConvertToStockAndQuantity(checkedQuantity float32, checkedDate string, openConfirmedQuantityInBaseUnit float32, stockIsFullyChecked bool) *StockAndQuantity {
	pm := &requests.StockAndQuantity{}

	pm.CheckedQuantity = checkedQuantity
	pm.CheckedDate = checkedDate
	pm.OpenConfirmedQuantityInBaseUnit = openConfirmedQuantityInBaseUnit
	pm.StockIsFullyChecked = stockIsFullyChecked

	data := pm
	res := StockAndQuantity{
		CheckedQuantity:                 data.CheckedQuantity,
		CheckedDate:                     data.CheckedDate,
		OpenConfirmedQuantityInBaseUnit: data.OpenConfirmedQuantityInBaseUnit,
		StockIsFullyChecked:             data.StockIsFullyChecked,
	}

	return &res
}

// 3. 利用可能在庫の再計算
func (psdc *SDC) ConvertToRecalculatedAvailableProductStock(availableProductStock, checkedQuantity, recalculatedAvailableProductStock float32) *RecalculatedAvailableProductStock {
	pm := &requests.RecalculatedAvailableProductStock{}

	pm.AvailableProductStock = availableProductStock
	pm.CheckedQuantity = checkedQuantity
	pm.RecalculatedAvailableProductStock = recalculatedAvailableProductStock

	data := pm
	res := RecalculatedAvailableProductStock{
		AvailableProductStock:             data.AvailableProductStock,
		CheckedQuantity:                   data.CheckedQuantity,
		RecalculatedAvailableProductStock: data.RecalculatedAvailableProductStock,
	}

	return &res
}
