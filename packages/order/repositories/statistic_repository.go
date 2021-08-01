package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/utils"
	"medilane-api/packages/order/responses"
	"medilane-api/requests"
)

type StatisticRepository struct {
	DB *gorm.DB
}

func NewStatisticRepository(db *gorm.DB) *StatisticRepository {
	return &StatisticRepository{DB: db}
}

func (statRepo *StatisticRepository) StatisticDrugStore(data *[]responses.DrugStoreStatistic, request *requests.DrugStoreStatisticRequest) error {
	sqlRaw := "SELECT FROM_UNIXTIME(ds.created_at/1000, \"%s\") as time, COUNT(*) as count" +
		" FROM drug_store ds" +
		" JOIN address a ON a.id = ds.address_id" +
		" WHERE ds.created_at >= ? AND ds.created_at <= ? AND a.area_id = ? " +
		" GROUP BY FROM_UNIXTIME(ds.created_at/1000, \"%s\")" +
		" ORDER BY time ASC;"

	var sql string
	switch request.Interval {
	case string(utils.Month):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m", "%Y-%m")
	case string(utils.Day):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m-%d", "%Y-%m-%d")
	}

	return statRepo.DB.Raw(sql, request.TimeFrom, request.TimeTo, request.AreaId).Scan(&data).Error
}

func (statRepo *StatisticRepository) StatisticOrderCount(data *[]responses.OrderStatisticCount, request *requests.OrderStatisticCountRequest) error {
	sqlRaw := "SELECT FROM_UNIXTIME(o.created_at/1000, \"%s\") as time, COUNT(*) as count \n" +
		" FROM `order` o \n" +
		" JOIN drug_store ds ON ds.id = o.drug_store_id \n" +
		" JOIN address a ON a.id = ds.address_id \n" +
		" WHERE o.created_at >= ? AND o.created_at <= ? AND a.area_id = ? \n" +
		" GROUP BY FROM_UNIXTIME(o.created_at/1000, \"%s\")\n" +
		" ORDER BY time;"
	var sql string
	switch request.Interval {
	case string(utils.Month):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m", "%Y-%m")
	case string(utils.Day):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m-%d", "%Y-%m-%d")
	}

	return statRepo.DB.Raw(sql, request.TimeFrom, request.TimeTo, request.AreaId).Scan(&data).Error
}

func (statRepo *StatisticRepository) StatisticProductTopCount(data *[]responses.ProductStatisticCount, request *requests.ProductStatisticCountRequest) error {
	sqlRaw := "SELECT date, count, product_name FROM " +
		" (\n\tSELECT date, count, product_name , row_number() over (partition by date order by count DESC) as stt " +
		" FROM (\n\t\tSELECT FROM_UNIXTIME(ds.created_at/1000, \"%s\") as date, " +
		" COUNT(od.quantity) as count , p.name as product_name\n\t\t" +
		" FROM order_detail od  \n\t\t" +
		" JOIN `order` o  ON o.id = od.order_id \n\t\t" +
		" JOIN drug_store ds ON ds.id = o.drug_store_id \n\t\t" +
		" JOIN address a ON a.id = ds.address_id \n\t\t" +
		" JOIN product p ON p.id = od.product_id \n\t\t" +
		" WHERE ds.created_at >= ? AND ds.created_at <= ? " +
		" AND a.area_id = ? \n\t\t" +
		" GROUP BY FROM_UNIXTIME(ds.created_at/1000, \"%s\"), product_name\n\t\t" +
		" ORDER BY date ASC , count DESC\n\t) " +
		" as sub ) " +
		" as sub2\n" +
		" WHERE stt <= ?"

	var sql string
	switch request.Interval {
	case string(utils.Month):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m", "%Y-%m")
	case string(utils.Day):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m-%d", "%Y-%m-%d")
	}
	type resp struct {
		Date        string `json:"date"`
		Count       int64  `json:"count"`
		ProductName string `json:"product_name"`
	}
	var rs []resp
	err := statRepo.DB.Raw(sql, request.TimeFrom, request.TimeTo, request.AreaId, request.Top).Scan(&rs).Error
	if err != nil {
		return err
	}

	var temp = make(map[string][]responses.ProductStatisticCountItem)
	for _, item := range rs {
		if _, ok := temp[item.Date]; !ok {
			temp[item.Date] = []responses.ProductStatisticCountItem{
				{
					ProductName: item.ProductName,
					Count:       item.Count,
				},
			}
		} else {
			temp[item.Date] = append(temp[item.Date], responses.ProductStatisticCountItem{
				ProductName: item.ProductName,
				Count:       item.Count,
			})
		}
	}
	for k, v := range temp {
		*data = append(*data, responses.ProductStatisticCount{
			Time: k,
			Data: v,
		})
	}
	return nil
}
func (statRepo *StatisticRepository) StatisticDrugStoreOrderTopCount(data *[]responses.OrderDrugstoreCount, request *requests.OrderStoreStatisticCountRequest) error {
	sqlRaw := "SELECT date, amount, store_name FROM (" +
		" \n\tSELECT date, amount, store_name , row_number() over (partition by date order by amount DESC) as stt " +
		" FROM (\n\t\t" +
		" SELECT FROM_UNIXTIME(o.created_at/1000, \"%s\") as date, SUM(o.sub_total) as amount, ds.store_name \n\t\t" +
		" FROM `order` o \n\t\t" +
		" JOIN drug_store ds ON ds.id = o.drug_store_id \n\t\t" +
		" JOIN address a ON a.id = ds.address_id \n\t\t" +
		" WHERE o.created_at >= ? AND o.created_at <= ? AND a.area_id = ? \n\t\t" +
		" GROUP BY FROM_UNIXTIME(o.created_at/1000, \"%s\"), ds.store_name \n\t\t" +
		" ORDER BY date ASC , amount DESC\n\t) " +
		" as sub ) " +
		" as sub2\n" +
		" WHERE stt <= ?"

	var sql string
	switch request.Interval {
	case string(utils.Month):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m", "%Y-%m")
	case string(utils.Day):
		sql = fmt.Sprintf(sqlRaw, "%Y-%m-%d", "%Y-%m-%d")
	}
	type resp struct {
		Date      string `json:"date"`
		Amount    int64  `json:"amount"`
		StoreName string `json:"store_name"`
	}
	var rs []resp
	err := statRepo.DB.Raw(sql, request.TimeFrom, request.TimeTo, request.AreaId, request.Top).Scan(&rs).Error
	if err != nil {
		return err
	}

	var temp = make(map[string][]responses.OrderDrugstoreCountItem)
	for _, item := range rs {
		if _, ok := temp[item.Date]; !ok {
			temp[item.Date] = []responses.OrderDrugstoreCountItem{
				{
					StoreName: item.StoreName,
					Amount:    item.Amount,
				},
			}
		} else {
			temp[item.Date] = append(temp[item.Date], responses.OrderDrugstoreCountItem{
				StoreName: item.StoreName,
				Amount:    item.Amount,
			})
		}
	}
	for k, v := range temp {
		*data = append(*data, responses.OrderDrugstoreCount{
			Time: k,
			Data: v,
		})
	}
	return nil
}
