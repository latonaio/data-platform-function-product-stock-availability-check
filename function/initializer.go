package function

import (
	"context"
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type Function struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *Function {
	return &Function{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *Function) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1. Product Stock Availability
		psdc.ProductStockAvailability, e = f.ProductStockAvailability(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	f.SetValue(psdc, osdc)

	return nil
}
