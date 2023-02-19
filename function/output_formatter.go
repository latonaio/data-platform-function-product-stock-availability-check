package function

import (
	dpfm_api_output_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
)

func (f *Function) SetValue(
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) {
	productStockAvailability := dpfm_api_output_formatter.ConvertToProductStockAvailability(psdc)

	osdc.Message = dpfm_api_output_formatter.Message{
		ProductStockAvailability: productStockAvailability,
	}
}
