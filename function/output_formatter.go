package function

import (
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
)

func (f *Function) SetValue(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	header, err := dpfm_api_output_formatter.ConvertToHeader(sdc, psdc)
	if err != nil {
		return err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Header: *header,
	}

	return err
}
