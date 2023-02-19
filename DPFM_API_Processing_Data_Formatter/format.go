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
		BusinessPartner:              sdc.Header.BusinessPartner,
		Product:                      sdc.Header.Product,
		Plant:                        sdc.Header.Plant,
		ProductStockAvailabilityDate: sdc.Header.Availability.ProductStockAvailabilityDate,
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
