package dpfm_api_output_formatter

import (
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
)

func ConvertToProductStockAvailability(psdc *dpfm_api_processing_data_formatter.SDC) *ProductStockAvailability {
	data := psdc.ProductStockAvailability

	res := &ProductStockAvailability{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return res
}
