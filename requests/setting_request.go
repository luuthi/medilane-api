package requests

import validation "github.com/go-ozzo/ozzo-validation"

type SettingRequest struct {
	Key     string `json:"Key"`
	Ios     string `json:"Ios"`
	Android string `json:"Android"`
	Config  string `json:"Config"`
}

func (rr SettingRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Key, validation.Required),
	)
}

type SearchSettingRequest struct {
	Key string `json:"Key"`
}
