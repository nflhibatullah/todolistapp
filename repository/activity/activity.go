package activity

import (
	"gorm.io/gorm"
	"todo/entity"
)

type (
	Activity struct {
		db *gorm.DB
	}

	ActivityRepository interface {
		CreateActivity(activity *entity.Activity) (*entity.Activity, error)
		GetActivityByID(id int64) (*entity.Activity, error)
		GetAllActivity() ([]*entity.Activity, error)
		UpdateActivity(activity *entity.Activity) (*entity.Activity, error)
		DeleteActivity(id int64) error
	}
)

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &Activity{
		db: db,
	}
}

func (a *Activity) CreateActivity(activity *entity.Activity) (*entity.Activity, error) {

	tx := a.db.Create(activity)

	if tx.Error != nil {
		return activity, tx.Error
	}

	return activity, nil
}

func (a *Activity) GetActivityByID(id int64) (*entity.Activity, error) {
	var activity entity.Activity
	tx := a.db.First(&activity, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &activity, nil
}

func (a *Activity) GetAllActivity() ([]*entity.Activity, error) {
	var activities []*entity.Activity
	tx := a.db.Find(&activities)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return activities, nil
}

func (a *Activity) UpdateActivity(activity *entity.Activity) (*entity.Activity, error) {

	var result = &entity.Activity{}

	tx := a.db.First(&result, activity.ID).Updates(activity)

	if tx.Error != nil {
		return nil, tx.Error
	}

	result.UpdatedAt = activity.UpdatedAt
	result.Title = activity.Title

	return result, nil
}

func (a *Activity) DeleteActivity(id int64) error {
	var activity = &entity.Activity{}
	tx := a.db.First(activity, id).Delete(activity, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
