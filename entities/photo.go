package entities

type PhotoEntityGetterInterface interface {
	GetTier() int
	GetSubject() string
	GetTabs() []string
}

type PhotoEntity struct {
	tier    int
	subject string
	tags    []string
}

func (photo PhotoEntity) Clone(i PhotoEntityGetterInterface) PhotoEntity {
	return PhotoEntity{
		tier:    i.GetTier(),
		subject: i.GetSubject(),
		tags:    i.GetTabs(),
	}
}

func (photo PhotoEntity) GetTier() int {
	return photo.tier
}

func (photo PhotoEntity) GetSubject() string {
	return photo.subject
}

func (photo PhotoEntity) GetTags() []string {
	return photo.tags
}
