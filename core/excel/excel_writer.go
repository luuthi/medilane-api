package excelWriter

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"medilane-api/core/funcHelpers"
	"medilane-api/models"
	"os"
	"path/filepath"
	time "time"
)

type Column struct {
	Width float64 `json:"width"`
	Value string  `json:"value"`
	Name  string  `json:"name"`
}

type ExcelWriter struct {
	Path         string
	File         *excelize.File
	CurrentSheet string
	Headers      []string
	Cols         map[string]Column
	rowIdx       int
}

func NewExcelWriter(fileName string, headers []string, columns []Column) (*ExcelWriter, error) {
	ew := &ExcelWriter{}
	err := ew.Init(fileName, headers, columns)
	if err != nil {
		return nil, err
	}
	return ew, nil
}

func (e *ExcelWriter) PrepareFolder(output string) (string, error) {
	workingDir, _ := os.Getwd()

	absPath := filepath.Join(workingDir, output)
	dirPath := filepath.Dir(absPath)

	// create new if not exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return absPath, nil
}

func (e *ExcelWriter) Init(fileName string, headers []string, columns []Column) error {
	fPath, err := e.PrepareFolder(fileName)
	if err != nil {
		return err
	}

	// new file excel
	f := excelize.NewFile()
	e.File = f

	// set headers
	e.Headers = headers

	e.Path = fPath

	e.Cols = make(map[string]Column)
	for _, v := range columns {
		e.Cols[v.Value] = v
	}

	return nil
}

func (e *ExcelWriter) SetSheetActive(sheetName string) {
	e.CurrentSheet = sheetName
	i := e.File.NewSheet(e.CurrentSheet)
	e.File.SetActiveSheet(i)
}

func (e *ExcelWriter) WriteHeader() error {
	//set header height
	err := e.File.SetRowHeight(e.CurrentSheet, e.rowIdx, 20)
	if err != nil {
		return err
	}

	for i := 1; i <= len(e.Headers); i++ {
		cellName, _ := excelize.CoordinatesToCellName(i, e.rowIdx)
		colName := cellName[:1]
		if err := e.File.SetColWidth(e.CurrentSheet, colName, colName, e.Cols[e.Headers[i-1]].Width); err != nil {
			return err
		}
		if err := e.File.SetCellValue(e.CurrentSheet, cellName, e.Cols[e.Headers[i-1]].Name); err != nil {
			return err
		}
	}
	styleId, err := e.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}, "font":{"bold":true,"italic":false,"family":"Arial","size":14,"color":"#000000"},"fill":{"type":"pattern","color":["#d4ecff"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}

	cellStart, _ := excelize.CoordinatesToCellName(1, e.rowIdx)
	cellEnd, _ := excelize.CoordinatesToCellName(len(e.Headers), e.rowIdx)
	err = e.File.SetCellStyle(e.CurrentSheet, cellStart, cellEnd, styleId)
	if err != nil {
		return err
	}
	e.rowIdx++
	return nil
}

func (e *ExcelWriter) WriteOrderHeader(order *models.Order) (err error) {
	//set first row to start
	e.rowIdx = 2
	//set header height
	err = e.File.SetRowHeight(e.CurrentSheet, e.rowIdx, 40)
	if err != nil {
		return err
	}
	styleHeaderId, err := e.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}, "font":{"bold":true,"italic":false,"family":"Arial","size":14,"color":"#FFFFFF"},"fill":{"type":"pattern","color":["#4c9cc4"],"pattern":1}}`)
	if err != nil {
		return err
	}
	styleHeaderValueId, err := e.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}, "font":{"bold":true,"italic":false,"family":"Arial","size":12,"color":"#000000"}}`)
	if err != nil {
		return err
	}
	styleDateTimeId, err := e.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}, "font":{"bold":true,"italic":false,"family":"Arial","size":12,"color":"#000000"}, "number_format": 22}`)
	if err != nil {
		return err
	}
	// col nhà thuốc
	cellStartDrugstore, _ := excelize.CoordinatesToCellName(1, e.rowIdx)
	cellEndDrugstore, _ := excelize.CoordinatesToCellName(2, e.rowIdx)
	err = e.File.MergeCell(e.CurrentSheet, cellStartDrugstore, cellEndDrugstore)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartDrugstore, "Nhà thuốc"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartDrugstore, cellStartDrugstore, styleHeaderId); err != nil {
		return err
	}
	cellStartDrugstoreValue, _ := excelize.CoordinatesToCellName(1, e.rowIdx+1)
	cellEndDrugstoreValue, _ := excelize.CoordinatesToCellName(2, e.rowIdx+1)
	err = e.File.MergeCell(e.CurrentSheet, cellStartDrugstoreValue, cellEndDrugstoreValue)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartDrugstoreValue, order.Drugstore.StoreName); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartDrugstoreValue, cellStartDrugstoreValue, styleHeaderValueId); err != nil {
		return err
	}

	// col địa chỉ
	cellStartAddress, _ := excelize.CoordinatesToCellName(3, e.rowIdx)
	cellEndAddress, _ := excelize.CoordinatesToCellName(5, e.rowIdx)
	err = e.File.MergeCell(e.CurrentSheet, cellStartAddress, cellEndAddress)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartAddress, "Địa chỉ"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartAddress, cellStartAddress, styleHeaderId); err != nil {
		return err
	}
	cellStartAddressValue, _ := excelize.CoordinatesToCellName(3, e.rowIdx+1)
	cellEndAddressValue, _ := excelize.CoordinatesToCellName(5, e.rowIdx+1)
	err = e.File.MergeCell(e.CurrentSheet, cellStartAddressValue, cellEndAddressValue)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartAddressValue, fmt.Sprintf("%s-%s-%s,%s", order.Address.Street, order.Address.Ward, order.Address.District, order.Address.Province)); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartAddressValue, cellStartAddressValue, styleHeaderValueId); err != nil {
		return err
	}

	// col ngày mua
	cellDate, _ := excelize.CoordinatesToCellName(6, e.rowIdx)
	if err := e.File.SetCellValue(e.CurrentSheet, cellDate, "Ngày mua"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellDate, cellDate, styleHeaderId); err != nil {
		return err
	}
	cellDateValue, _ := excelize.CoordinatesToCellName(6, e.rowIdx+1)
	if err := e.File.SetCellValue(e.CurrentSheet, cellDateValue, time.Unix(order.CreatedAt/1000, 0)); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellDateValue, cellDateValue, styleDateTimeId); err != nil {
		return err
	}

	// col mã đơn hàng
	cellCode, _ := excelize.CoordinatesToCellName(7, e.rowIdx)
	if err := e.File.SetCellValue(e.CurrentSheet, cellCode, "Mã đơn hàng"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellCode, cellCode, styleHeaderId); err != nil {
		return err
	}
	cellCodeValue, _ := excelize.CoordinatesToCellName(7, e.rowIdx+1)
	if err := e.File.SetCellValue(e.CurrentSheet, cellCodeValue, order.OrderCode); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellCodeValue, cellCodeValue, styleHeaderValueId); err != nil {
		return err
	}

	// col trạng thái
	cellStatus, _ := excelize.CoordinatesToCellName(8, e.rowIdx)
	if err := e.File.SetCellValue(e.CurrentSheet, cellStatus, "Trạng thái"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStatus, cellStatus, styleHeaderId); err != nil {
		return err
	}
	cellStatusValue, _ := excelize.CoordinatesToCellName(8, e.rowIdx+1)
	if err := e.File.SetCellValue(e.CurrentSheet, cellStatusValue, funcHelpers.VietnameseStatus(order.Status)); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStatusValue, cellStatusValue, styleHeaderValueId); err != nil {
		return err
	}
	e.rowIdx += 6
	return
}

func (e *ExcelWriter) WriteOrderFooter(order *models.Order) (err error) {
	styleNumberSummaryId, err := e.File.NewStyle(`{"alignment":{"horizontal":"right","vertical":"center","wrap_text":true}, "font":{"bold":true,"italic":false,"family":"Arial","size":12,"color":"#000000"}, "number_format": 3}`)
	if err != nil {
		return err
	}
	// row subTotal
	cellStartSubTotal, _ := excelize.CoordinatesToCellName(1, e.rowIdx)
	cellEndSubTotal, _ := excelize.CoordinatesToCellName(len(e.Headers)-1, e.rowIdx)
	cellValueSubTotal, _ := excelize.CoordinatesToCellName(len(e.Headers), e.rowIdx)
	err = e.File.MergeCell(e.CurrentSheet, cellStartSubTotal, cellEndSubTotal)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartSubTotal, "Tạm tính"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartSubTotal, cellStartSubTotal, styleNumberSummaryId); err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellValueSubTotal, order.SubTotal); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellValueSubTotal, cellValueSubTotal, styleNumberSummaryId); err != nil {
		return err
	}
	e.rowIdx++
	// row discount
	cellStartDiscount, _ := excelize.CoordinatesToCellName(1, e.rowIdx)
	cellEndSubDiscount, _ := excelize.CoordinatesToCellName(len(e.Headers)-1, e.rowIdx)
	cellValueDiscount, _ := excelize.CoordinatesToCellName(len(e.Headers), e.rowIdx)
	err = e.File.MergeCell(e.CurrentSheet, cellStartDiscount, cellEndSubDiscount)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartDiscount, "Giảm giá"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartDiscount, cellStartDiscount, styleNumberSummaryId); err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellValueDiscount, order.Discount); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellValueDiscount, cellValueDiscount, styleNumberSummaryId); err != nil {
		return err
	}
	e.rowIdx++

	// row shipping
	cellStartShipping, _ := excelize.CoordinatesToCellName(1, e.rowIdx)
	cellEndShipping, _ := excelize.CoordinatesToCellName(len(e.Headers)-1, e.rowIdx)
	cellValueShipping, _ := excelize.CoordinatesToCellName(len(e.Headers), e.rowIdx)
	err = e.File.MergeCell(e.CurrentSheet, cellStartShipping, cellEndShipping)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartShipping, "Phí giao hàng"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartShipping, cellStartShipping, styleNumberSummaryId); err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellValueShipping, order.ShippingFee); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellValueShipping, cellValueShipping, styleNumberSummaryId); err != nil {
		return err
	}
	e.rowIdx++

	// row total
	cellStartTotal, _ := excelize.CoordinatesToCellName(1, e.rowIdx)
	cellEndTotal, _ := excelize.CoordinatesToCellName(len(e.Headers)-1, e.rowIdx)
	cellValueTotal, _ := excelize.CoordinatesToCellName(len(e.Headers), e.rowIdx)
	err = e.File.MergeCell(e.CurrentSheet, cellStartTotal, cellEndTotal)
	if err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellStartTotal, "Tổng"); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellStartTotal, cellStartTotal, styleNumberSummaryId); err != nil {
		return err
	}
	if err := e.File.SetCellValue(e.CurrentSheet, cellValueTotal, order.Total); err != nil {
		return err
	}
	if err := e.File.SetCellStyle(e.CurrentSheet, cellValueTotal, cellValueTotal, styleNumberSummaryId); err != nil {
		return err
	}
	e.rowIdx++
	return
}

func (e *ExcelWriter) WriteOrderBody(order *models.Order) (err error) {
	err = e.WriteHeader()
	if err != nil {
		return err
	}

	styleNumberId, err := e.File.NewStyle(`{"alignment":{"horizontal":"right","vertical":"center","wrap_text":true}, "font":{"bold":false,"italic":false,"family":"Arial","size":12,"color":"#000000"}, "number_format": 3}`)
	if err != nil {
		return err
	}

	styleStringId, err := e.File.NewStyle(`{"alignment":{"horizontal":"left","vertical":"center","wrap_text":true}, "font":{"bold":false,"italic":false,"family":"Arial","size":12,"color":"#000000"}}`)
	if err != nil {
		return err
	}

	styleStringCenterId, err := e.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}, "font":{"bold":false,"italic":false,"family":"Arial","size":12,"color":"#000000"}}`)
	if err != nil {
		return err
	}

	styleCenterId, err := e.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center","wrap_text":true}, "font":{"bold":false,"italic":false,"family":"Arial","size":12,"color":"#000000"}}`)
	if err != nil {
		return err
	}

	for rIdx, r := range order.OrderDetails {
		for cIdx, c := range e.Headers {
			cellName, _ := excelize.CoordinatesToCellName(cIdx+1, e.rowIdx+rIdx)
			var cellStyleId int
			var cellData interface{}
			switch c {
			case "Cost":
				cellData = r.Cost
				cellStyleId = styleNumberId
			case "Quantity":
				cellData = r.Quantity
				cellStyleId = styleNumberId
			case "Discount":
				cellData = r.Discount
				cellStyleId = styleNumberId
			case "ProductName":
				cellData = r.Product.Name
				cellStyleId = styleStringId
			case "Unit":
				cellData = r.Variant.Name
				cellStyleId = styleStringCenterId
			case "No":
				cellData = rIdx + 1
				cellStyleId = styleCenterId
			case "SubTotal":
				cellData = float64(r.Quantity) * r.Cost
				cellStyleId = styleNumberId
			case "Total":
				cellData = float64(r.Quantity)*r.Cost - (float64(r.Quantity) * r.Cost * r.Discount / 100)
				cellStyleId = styleNumberId
			}
			if err := e.File.SetCellValue(e.CurrentSheet, cellName, cellData); err != nil {
				return err
			}
			if err := e.File.SetCellStyle(e.CurrentSheet, cellName, cellName, cellStyleId); err != nil {
				return err
			}
		}
	}

	e.rowIdx += len(order.OrderDetails) + 1

	return
}
