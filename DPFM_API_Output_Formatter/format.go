package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeader(
	sdc *dpfm_api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*Header, error) {
	var err error

	header := &Header{}
	inputHeader := sdc.Header

	// 入力ファイル
	header, err = jsonTypeConversion(header, inputHeader)
	if err != nil {
		return nil, err
	}

	header.CheckedQuantity = psdc.StockAndQuantity.CheckedQuantity
	header.CheckedDate = psdc.StockAndQuantity.CheckedDate
	header.OpenConfirmedQuantityInBaseUnit = psdc.StockAndQuantity.OpenConfirmedQuantityInBaseUnit
	header.StockIsFullyChecked = psdc.StockAndQuantity.StockIsFullyChecked
	header.AvailableProductStock = psdc.RecalculatedAvailableProductStock.RecalculatedAvailableProductStock

	return header, nil
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
