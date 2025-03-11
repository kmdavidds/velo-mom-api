package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/valentinusdelvin/velo-mom-api/entity"
	"github.com/valentinusdelvin/velo-mom-api/models"
	"gorm.io/gorm"
)

type InterWebinarRepository interface {
	CreateWebinar(Webinar entity.Webinar) (entity.Webinar, error)
	GetWebinars() ([]models.GetWebinars, error)
	GetWebinarByID(id string) (entity.Webinar, error)
	CreateWebinarAttendee(tx *gorm.DB, attendee entity.WebinarAttendee) error
	UpdateWebinarInfo(tx *gorm.DB, webinarID uuid.UUID) error
}

type WebinarRepository struct {
	db *gorm.DB
}

func NewWebinarRepository(db *gorm.DB) InterWebinarRepository {
	return &WebinarRepository{
		db: db,
	}
}

func (wr *WebinarRepository) CreateWebinar(webinar entity.Webinar) (entity.Webinar, error) {
	err := wr.db.Create(&webinar).Error
	if err != nil {
		return webinar, err
	}
	return webinar, nil
}

func (wr *WebinarRepository) GetWebinars() ([]models.GetWebinars, error) {
	var webinars []models.GetWebinars

	err := wr.db.Table("webinars").Find(&webinars).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(webinars)
	return webinars, nil
}

func (wr *WebinarRepository) GetWebinarByID(id string) (entity.Webinar, error) {
	var webinar entity.Webinar

	err := wr.db.Table("webinars").Where("id = ?", id).Find(&webinar).Error
	if err != nil {
		return webinar, err
	}

	return webinar, nil
}

func (wr *WebinarRepository) CreateWebinarAttendee(tx *gorm.DB, attendee entity.WebinarAttendee) error {
	return tx.Create(&attendee).Error
}

func (wr *WebinarRepository) UpdateWebinarInfo(tx *gorm.DB, webinarID uuid.UUID) error {
	return tx.Model(&entity.Webinar{}).
		Where("id = ? AND quota > 0", webinarID).
		Updates(map[string]interface{}{
			"quota":     gorm.Expr("quota - 1"),
			"is_bought": true,
		}).Error
}
