package responses

//DrugStoreStatistic drugstore count
type DrugStoreStatistic struct {
	Time  string `json:"time"`
	Count int64  `json:"count"`
}
type DrugStoreStatisticResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    []DrugStoreStatistic `json:"data"`
}

//ProductStatisticCount product statistic
type ProductStatisticCount struct {
	Time string                      `json:"time"`
	Data []ProductStatisticCountItem `json:"data"`
}
type ProductStatisticCountItem struct {
	ProductName string `json:"product_name"`
	Count       int64  `json:"count"`
}
type ProductStatisticCountResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    []ProductStatisticCount `json:"data"`
}

//OrderStatisticCount order count
type OrderStatisticCount struct {
	Time  string `json:"time"`
	Count int64  `json:"count"`
}

type OrderStatisticCountResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []OrderStatisticCount `json:"data"`
}

//OrderDrugstoreCount  statistic sum amount order by drugstore
type OrderDrugstoreCount struct {
	Time string                    `json:"time"`
	Data []OrderDrugstoreCountItem `json:"data"`
}
type OrderDrugstoreCountItem struct {
	StoreName string `json:"store_name"`
	Amount    int64  `json:"amount"`
}
type OrderDrugstoreCountResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []OrderDrugstoreCount `json:"data"`
}
