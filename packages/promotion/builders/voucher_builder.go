package builders

import "medilane-api/models"

type VoucherBuilder struct {
	name     string
	_type    string
	value    float32
	note     string
	maxValue float32
	unit     string
	id       uint
	deleted  bool
}

func NewVoucherBuilder() *VoucherBuilder {
	return &VoucherBuilder{}
}

func (voucherBuilder *VoucherBuilder) SetName(name string) *VoucherBuilder {
	voucherBuilder.name = name
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetType(_type string) *VoucherBuilder {
	voucherBuilder._type = _type
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetUnit(unit string) *VoucherBuilder {
	voucherBuilder.unit = unit
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetDeleted(isDeleted bool) *VoucherBuilder {
	voucherBuilder.deleted = isDeleted
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetValue(value float32) *VoucherBuilder {
	voucherBuilder.value = value
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetMaxValue(maxValue float32) *VoucherBuilder {
	voucherBuilder.maxValue = maxValue
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetNote(note string) *VoucherBuilder {
	voucherBuilder.note = note
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) SetID(id uint) *VoucherBuilder {
	voucherBuilder.id = id
	return voucherBuilder
}

func (voucherBuilder *VoucherBuilder) Builder() models.Voucher {
	common := models.CommonModelFields{
		ID: voucherBuilder.id,
	}

	voucher := models.Voucher{
		CommonModelFields: common,
		Name:              voucherBuilder.name,
		Type:              voucherBuilder._type,
		Value:             voucherBuilder.value,
		MaxValue:          voucherBuilder.maxValue,
		Unit:              voucherBuilder.unit,
		Note:              voucherBuilder.note,
		Deleted:           voucherBuilder.deleted,
	}
	return voucher
}
