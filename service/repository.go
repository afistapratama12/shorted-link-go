package service

import "gorm.io/gorm"

type Repository interface {
	Add(shortLink ShortedLink) (ShortedLink, error)
	FindLink(shortLink string) (ShortedLink, error)
	FindById(id string) (ShortedLink, error)
	UpdateShortLink(id string, dataUpdate map[string]interface{}) (ShortedLink, error)
}

type repository struct {
	db *gorm.DB
}

// func to intialize new repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Add(shortLink ShortedLink) (ShortedLink, error) {
	if err := r.db.Create(&shortLink).Error; err != nil {
		return shortLink, err
	}

	return shortLink, nil
}

func (r *repository) FindLink(shortLink string) (ShortedLink, error) {
	var shortData ShortedLink

	if err := r.db.Where("short_link = ?", shortLink).Find(&shortData).Error; err != nil {
		return shortData, err
	}

	return shortData, nil

}

func (r *repository) FindById(id string) (ShortedLink, error) {
	var shortData ShortedLink

	if err := r.db.Where("id = ?", id).Find(&shortData).Error; err != nil {
		return shortData, err
	}

	return shortData, nil
}

func (r *repository) UpdateShortLink(id string, dataUpdate map[string]interface{}) (ShortedLink, error) {
	var shortData ShortedLink

	if err := r.db.Model(shortData).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return shortData, err
	}

	if err := r.db.Where("id = ?", id).Find(&shortData).Error; err != nil {
		return shortData, err
	}

	return shortData, nil
}
