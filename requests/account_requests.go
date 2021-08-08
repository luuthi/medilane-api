package requests

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	utils2 "medilane-api/core/utils"
)

type SearchAccountRequest struct {
	Username string     `json:"username"  example:"admin"`
	FullName string     `json:"full_name"  example:"admin"`
	Email    string     `json:"email" example:"admin@gmail.com"`
	Status   string     `json:"status" example:"true"`
	Type     []string   `json:"type" example:"staff/user/supplier/manufacturer"`
	IsAdmin  *bool      `json:"is_admin" example:"true"`
	Limit    int        `json:"limit" example:"10"`
	Offset   int        `json:"offset" example:"0"`
	Sort     SortOption `json:"sort"`
	TimeFrom *int64     `json:"time_from"`
	TimeTo   *int64     `json:"time_to"`
}

type SortOption struct {
	SortField     string `json:"sort_field"`
	SortDirection string `json:"sort_direction"`
}

func (rr SearchAccountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
		validation.Field(&rr.TimeFrom, validation.Min(0)),
		validation.Field(&rr.TimeTo, validation.Min(0)),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(rr.TimeFrom, rr.TimeTo))),
	)
}

type EditAccountRequest struct {
	FullName *string   `json:"full_name"  example:"admin"`
	Email    *string   `json:"email" example:"admin@gmail.com"`
	Status   *bool     `json:"status" example:"true"`
	IsAdmin  *bool     `json:"is_admin" example:"true"`
	Roles    *[]string `json:"roles"`
}

func (rr EditAccountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Email, validation.Required, is.Email),
		validation.Field(&rr.FullName, validation.Required),
		validation.Field(&rr.Status, validation.NotNil),
		validation.Field(&rr.IsAdmin, validation.NotNil))
}

type CreateAccountRequest struct {
	Email       string   `json:"email" validate:"required" example:"john.doe@gmail.com"`
	Username    string   `json:"username" validate:"required" example:"JohnDoe"`
	Password    string   `json:"password"  validate:"required" example:"123qweA@"`
	FullName    string   `json:"Name" validate:"required" example:"John Doe"`
	IsAdmin     *bool    `json:"IsAdmin" validate:"required" example:"true" `
	Type        string   `json:"Type"  validate:"required" example:"super_admin/staff/user/supplier/manufacturer"`
	DrugStoreID *uint    `json:"DrugStoreID"`
	PartnerID   *uint    `json:"PartnerID"`
	Roles       []string `json:"Roles"`
}

func (rr CreateAccountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Email, validation.Required, is.Email),
		validation.Field(&rr.Username, validation.Required, validation.Length(3, 32)),
		validation.Field(&rr.Password, validation.Required, validation.Length(6, 32)),
		validation.Field(&rr.FullName, validation.Required),
		validation.Field(&rr.IsAdmin, validation.NotNil),
		validation.Field(&rr.Type, validation.In(string(utils2.STAFF), string(utils2.USER), string(utils2.SUPPLIER), string(utils2.MANUFACTURER)),
			validation.By(checkRequireByType(rr.Type, rr.DrugStoreID, rr.PartnerID))),
	)
}
func checkRequireByType(_type string, DrugStoreID *uint, PartnerID *uint) validation.RuleFunc {
	return func(value interface{}) error {
		if _type == string(utils2.USER) {
			if DrugStoreID == nil {
				return errors.New("type is user require drugstore ID")
			}
		} else if _type == string(utils2.SUPPLIER) || _type == string(utils2.MANUFACTURER) {
			if PartnerID == nil {
				return errors.New("type is partner require partner ID")
			}
		}
		return nil
	}
}

type StaffRelationship struct {
	DrugStoreId  uint   `json:"DrugStoresId"`
	Relationship string `json:"Relationship"`
}

type AssignStaffRequest struct {
	DrugStoresIdLst []uint `json:"DrugStoresIdLst"`
}

func (rr AssignStaffRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.DrugStoresIdLst, validation.Required, validation.Length(1, 1000)))
}
