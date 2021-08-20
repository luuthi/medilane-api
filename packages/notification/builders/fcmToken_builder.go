package builders

import "medilane-api/models"

type FcmTokenBuilder struct {
	token string
	user  uint
}

func NewFcmTokenBuilder() *FcmTokenBuilder {
	return &FcmTokenBuilder{}
}

func (fcmTokenBuilder *FcmTokenBuilder) SetToken(token string) (z *FcmTokenBuilder) {
	fcmTokenBuilder.token = token
	return fcmTokenBuilder
}

func (fcmTokenBuilder *FcmTokenBuilder) SetUser(user uint) (z *FcmTokenBuilder) {
	fcmTokenBuilder.user = user
	return fcmTokenBuilder
}

func (fcmTokenBuilder *FcmTokenBuilder) Build() models.FcmToken {
	area := models.FcmToken{
		Token: fcmTokenBuilder.token,
		User:  fcmTokenBuilder.user,
	}

	return area
}
