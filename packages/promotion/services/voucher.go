package services

import (
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/promotion/builders"
	"medilane-api/requests"
)

func (promoService *Service) CreateVoucher(request *requests.VoucherRequest) (error, *models.Voucher) {
	voucher := builders.NewVoucherBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		SetValue(request.Value).
		SetMaxValue(request.MaxValue).
		SetUnit(request.Unit).
		SetType(request.Type).
		Builder()

	err := promoService.DB.Table(utils.TblVoucher).Create(&voucher).Error
	if err != nil {
		return err, nil
	}
	return nil, &voucher
}

func (promoService *Service) EditVoucher(request *requests.VoucherRequest, id uint) (error, *models.Voucher) {
	voucher := builders.NewVoucherBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		SetValue(request.Value).
		SetType(request.Type).
		SetMaxValue(request.MaxValue).
		SetUnit(request.Unit).
		SetID(id).
		Builder()
	err := promoService.DB.Table(utils.TblVoucher).Updates(&voucher).Error
	if err != nil {
		return err, nil
	}
	return nil, &voucher
}

func (promoService *Service) DeleteVoucher(id uint) error {
	voucher := builders.NewVoucherBuilder().
		SetID(id).
		SetDeleted(true).
		Builder()
	err := promoService.DB.Table(utils.TblVoucher).Updates(&voucher).Error
	if err != nil {
		return err
	}
	return nil
}
