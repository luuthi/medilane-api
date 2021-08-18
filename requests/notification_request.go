package requests

import validation "github.com/go-ozzo/ozzo-validation"

type SearchNotificationRequest struct {
	UserId uint `json:"UserId"`
	Limit  int  `json:"limit" example:"10"`
	Offset int  `json:"offset" example:"0"`
}

func (rr SearchNotificationRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type CreateFcmToken struct {
	Token string `json:"Token"`
	User  uint   `json:"User"`
}

func (rr CreateFcmToken) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Token, validation.Required),
		validation.Field(&rr.User, validation.Required),
	)
}
