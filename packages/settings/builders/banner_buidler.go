package builders

import "medilane-api/models"

type BannerBuilder struct {
	url        string
	starTime   int64
	expireTime int64
	visible    *bool
	id         uint
}

func NewBannerBuilder() *BannerBuilder {
	return &BannerBuilder{}
}

func (bannerBuilder *BannerBuilder) SetURL(url string) *BannerBuilder {
	bannerBuilder.url = url
	return bannerBuilder
}

func (bannerBuilder *BannerBuilder) SetVisible(visible *bool) *BannerBuilder {
	bannerBuilder.visible = visible
	return bannerBuilder
}

func (bannerBuilder *BannerBuilder) SetStartTime(s int64) *BannerBuilder {
	bannerBuilder.starTime = s
	return bannerBuilder
}

func (bannerBuilder *BannerBuilder) SetExpireTime(e int64) *BannerBuilder {
	bannerBuilder.expireTime = e
	return bannerBuilder
}

func (bannerBuilder *BannerBuilder) SetId(id uint) *BannerBuilder {
	bannerBuilder.id = id
	return bannerBuilder
}

func (bannerBuilder *BannerBuilder) Build() models.Banner {
	common := models.CommonModelFields{
		ID: bannerBuilder.id,
	}
	appSetting := models.Banner{
		CommonModelFields: common,
		Url:               bannerBuilder.url,
		StartTime:         bannerBuilder.starTime,
		ExpireTime:        bannerBuilder.expireTime,
		Visible:           bannerBuilder.visible,
	}

	return appSetting
}
