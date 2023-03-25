package api_processing_data_formatter

import (
	"data-platform-function-product-stock-availability-check/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

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
	pm := &requests.ProductStockAvailability{}

	for i := 0; true; i++ {

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_stock_availability_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := &ProductStockAvailability{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductStockAvailabilityKeyBylotto(sdc *dpfm_api_input_reader.SDC) *ProductStockAvailabilityKeyBylotto {
	pm := &requests.ProductStockAvailabilityKeyBylotto{
		BusinessPartner:              *sdc.Header.BusinessPartner,
		Product:                      *sdc.Header.Product,
		Plant:                        *sdc.Header.Plant,
		Batch:                        *sdc.Header.Batch,
		ProductStockAvailabilityDate: *sdc.Header.ProductStockAvailabilityDate,
	}

	data := pm
	res := ProductStockAvailabilityKeyBylotto{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
	}

	return &res
}

func (psdc *SDC) ConvertToProductStockAvailabilityBylotto(rows *sql.Rows) (*ProductStockAvailability, error) {
	pm := &requests.ProductStockAvailability{}

	for i := 0; true; i++ {

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_stock_availability_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			&pm.Batch,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := &ProductStockAvailability{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res, nil
}

// 2
func (psdc *SDC) ConvertToComparisonAvailableStock(sdc *dpfm_api_input_reader.SDC) *ComparisonStock {
	pm := &requests.ComparisonStock{
		CheckedQuantity:                 *sdc.Header.RequestedQuantity,
		CheckedDate:                     psdc.ProductStockAvailability.ProductStockAvailabilityDate,
		OpenConfirmedQuantityInBaseUnit: 0,
		StockIsFullyChecked:             true,
	}

	data := pm
	res := ComparisonStock{
		CheckedQuantity:                 data.CheckedQuantity,
		CheckedDate:                     data.CheckedDate,
		OpenConfirmedQuantityInBaseUnit: data.OpenConfirmedQuantityInBaseUnit,
		StockIsFullyChecked:             data.StockIsFullyChecked,
	}

	return &res
}

func (psdc *SDC) ConvertToComparisonRequestedStock(sdc *dpfm_api_input_reader.SDC, difference float32) *ComparisonStock {
	pm := &requests.ComparisonStock{
		CheckedQuantity:                 psdc.ProductStockAvailability.AvailableProductStock,
		CheckedDate:                     psdc.ProductStockAvailability.ProductStockAvailabilityDate,
		OpenConfirmedQuantityInBaseUnit: difference,
		StockIsFullyChecked:             false,
	}

	data := pm
	res := ComparisonStock{
		CheckedQuantity:                 data.CheckedQuantity,
		CheckedDate:                     data.CheckedDate,
		OpenConfirmedQuantityInBaseUnit: data.OpenConfirmedQuantityInBaseUnit,
		StockIsFullyChecked:             data.StockIsFullyChecked,
	}

	return &res
}

func (psdc *SDC) ConvertToRecalculatedAvailableProductStock(sdc *dpfm_api_input_reader.SDC, difference float32) *RecalculatedAvailableProductStock {
	pm := &requests.RecalculatedAvailableProductStock{
		RecalculatedAvailableProductStock: difference,
	}

	data := pm
	res := RecalculatedAvailableProductStock{
		RecalculatedAvailableProductStock: data.RecalculatedAvailableProductStock,
	}

	return &res
}
