package models

type CommonModelFields struct {
	ID        uint  `json:"-" gorm:"primary_key"`
	FakeId    *UID  `json:"id" gorm:"-"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (m *CommonModelFields) GenUID(dbType int) {
	uid := NewUID(uint32(m.ID), dbType, 1)
	m.FakeId = &uid
}
