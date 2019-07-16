package user

import (
	"github.com/jinzhu/gorm"
	"github.com/yaumianwar/sample-golang-api/shared"
	"github.com/yaumianwar/sample-golang-api/user/form"
	"github.com/yaumianwar/sample-golang-api/user/model"
)

// GetAllUsers : to get all users data
func (svc *Service) GetAllUsers() ([]model.User, error) {
	var mUsers []model.User

	// database query, use Find to get all data
	if err := svc.db.Table(model.User{}.TableName()).
		Find(&mUsers).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = shared.ErrDataNotFound
		}
		return mUsers, err
	}

	return mUsers, nil
}

// GetSingleUser : to get single user data by id
func (svc *Service) GetSingleUser(ID uint64) (model.User, error) {
	var mUser model.User

	// database query, use First to get single data. First equal ASC LIMIT 1
	if err := svc.db.Table(mUser.TableName()).
		Where("id = ?", ID).
		First(&mUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = shared.ErrDataNotFound
		}
		return mUser, err
	}

	return mUser, nil
}

// CreateUser : to create new user data
func (svc *Service) CreateUser(data form.UserForm) (model.User, error) {
	var mUser model.User

	// assign data from form to model
	mUser = model.User{
		Name:     data.Name,
		SkinType: data.SkinType,
		Age:      data.Age,
	}

	// database query, insert data to database
	if err := svc.db.Table(mUser.TableName()).
		Create(&mUser).Error; err != nil {
		return mUser, err
	}

	return mUser, nil
}

// UpdateUser : to update existing user data
func (svc *Service) UpdateUser(mUser model.User, data form.UserForm) (model.User, error) {
	// assign data from form to model
	mUser.Name = data.Name
	mUser.SkinType = data.SkinType
	mUser.Age = data.Age

	// database query, update existing data
	if err := svc.db.Table(mUser.TableName()).
		Save(&mUser).Error; err != nil {
		return mUser, err
	}

	return mUser, nil
}

// IsSkincareExist : to check if user have skincare data
func (svc *Service) IsSkincareExist(userID uint64) bool {
	var count uint64
	if err := svc.db.Table(model.Skincare{}.TableName()).
		Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return true
	}
	return count > 0
}

// DeleteUser : to delete existing user data
func (svc *Service) DeleteUser(mUser model.User) error {
	// check existing user skincare before delete user data
	if svc.IsSkincareExist(mUser.ID) {
		return shared.ErrDeleteSkincareBefore
	}

	// database query, delete existing data
	if err := svc.db.Table(mUser.TableName()).
		Delete(&mUser).Error; err != nil {
		return err
	}

	return nil
}
