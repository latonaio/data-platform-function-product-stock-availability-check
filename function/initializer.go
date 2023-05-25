package function

import (
	"context"
	dpfm_api_input_reader "data-platform-function-product-stock-availability-check/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "data-platform-function-product-stock-availability-check/DPFM_API_Processing_Data_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"

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

func (f *Function) ProductStockAvailabilityType(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
) (*dpfm_api_processing_data_formatter.ProductStockAvailabilityType, error) {
	isProductStockAvailability := false
	isProductStockAvailabilityByBatch := false
	isProductStockAvailabilityByStorageBin := false
	isProductStockAvailabilityByStorageBinByBatch := false

	if sdc.Header.Product == nil || sdc.Header.BusinessPartner == nil || sdc.Header.Plant == nil || sdc.Header.ProductStockAvailabilityDate == nil {
		return nil, xerrors.New("入力ファイルのProduct,BusinessPartner,PlantまたはProductStockAvailabilityDateがnullです。")
	}
	if len(*sdc.Header.Product) == 0 || len(*sdc.Header.Plant) == 0 || len(*sdc.Header.ProductStockAvailabilityDate) == 0 {
		return nil, xerrors.New("入力ファイルのProduct,PlantまたはProductStockAvailabilityDateが空文字です。")
	}

	if sdc.Header.StorageLocation != nil && sdc.Header.StorageBin != nil && sdc.Header.Batch != nil &&
		len(*sdc.Header.StorageLocation) != 0 && len(*sdc.Header.StorageBin) != 0 && len(*sdc.Header.Batch) != 0 {
		isProductStockAvailabilityByStorageBinByBatch = true
	} else if sdc.Header.StorageLocation != nil && sdc.Header.StorageBin != nil &&
		len(*sdc.Header.StorageLocation) != 0 && len(*sdc.Header.StorageBin) != 0 {
		isProductStockAvailabilityByStorageBin = true
	} else if sdc.Header.Batch != nil && len(*sdc.Header.Batch) != 0 {
		isProductStockAvailabilityByBatch = true
	} else {
		isProductStockAvailability = true
	}

	data := psdc.ConvertToProductStockAvailabilityType(isProductStockAvailability, isProductStockAvailabilityByBatch, isProductStockAvailabilityByStorageBin, isProductStockAvailabilityByStorageBinByBatch)

	return data, nil
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

	psdc.ProductStockAvailabilityType, err = f.ProductStockAvailabilityType(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if psdc.ProductStockAvailabilityType.IsProductStockAvailability {
			// 1-1. Product Stock Availability(通常在庫確認)
			psdc.ProductStockAvailability, e = f.ProductStockAvailability(sdc, psdc)
		} else if psdc.ProductStockAvailabilityType.IsProductStockAvailabilityByBatch {
			// 1-2. Product Stock Availability(ロット在庫確認)
			psdc.ProductStockAvailability, e = f.ProductStockAvailabilityByBatch(sdc, psdc)
		} else if psdc.ProductStockAvailabilityType.IsProductStockAvailabilityByStorageBin {
			// 1-3. Product Stock Availability(棚番通常在庫確認)
			psdc.ProductStockAvailability, e = f.ProductStockAvailabilityByStorageBin(sdc, psdc)
		} else if psdc.ProductStockAvailabilityType.IsProductStockAvailabilityByStorageBinByBatch {
			// 1-4. Product Stock Availability(棚番ロット在庫確認)
			psdc.ProductStockAvailability, e = f.ProductStockAvailabilityByStorageBinByBatch(sdc, psdc)
		}
		if e != nil {
			err = e
			return
		}

		// 2-1 AvailableProductStockの値とRequestedQuantityの値を比較 //1-1~1-4
		psdc.ComparisonStockAndQuantity, e = f.ComparisonStockAndQuantity(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-2,2-3. 利用可能在庫と要求数量の比較 //1-1~1-4,2-1
		if psdc.ComparisonStockAndQuantity.IsAvailableProductStock {
			// 2-2. AvailableProductStock≧RequestedQuantityの場合
			psdc.StockAndQuantity = f.AvailableProductStockProcess(sdc, psdc)
		} else if psdc.ComparisonStockAndQuantity.IsRequestedQuantity {
			// 2-3. AvailableProductStock＜RequestedQuantityの場合
			psdc.StockAndQuantity = f.RequestedQuantityProcess(sdc, psdc)
		}

		//3. 利用可能在庫の再計算 //1-1~1-4,2-2or2-3
		psdc.RecalculatedAvailableProductStock = f.RecalculatedAvailableProductStock(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	f.SetValue(sdc, psdc, osdc)

	return nil
}
