package funcHelpers

import "medilane-api/core/utils/order"

func StringContain(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func UintContains(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func VietnameseStatus(statusOrder string) string {
	switch statusOrder {
	case order.Confirm.String():
		return "Chưa xác nhận"
	case order.Confirmed.String():
		return "Đã xác nhận"
	case order.Delivered.String():
		return "Đã giao hàng"
	case order.Delivery.String():
		return "Đang giao hàng"
	case order.Received.String():
		return "Khách đã nhận hàng"
	case order.Processing.String():
		return "Đã xử lý"
	case order.Packaging.String():
		return "Đã đóng gói"
	default:
		return ""
	}
}
