package user

import (
	"github.com/jinzhu/gorm"
	"github.com/yaumianwar/sample-golang-api/shared"
	"github.com/yaumianwar/sample-golang-api/user/form"
	"github.com/yaumianwar/sample-golang-api/user/model"
)

// GetAllSkincareByUser : to get all skincare data by user
func (svc *Service) GetAllSkincareByUser(mUser model.User) ([]model.Skincare, error) {
	var mSkincares []model.Skincare

	// database query, use Find to get all data
	if err := svc.db.Table(model.Skincare{}.TableName()).
		Where("user_id = ?", mUser.ID).
		Find(&mSkincares).Error; err != nil {
		return mSkincares, err
	}

	if len(mSkincares) <= 0 {
		return mSkincares, shared.ErrDataNotFound
	}

	return mSkincares, nil
}

// GetSingleSkincare : to get single skincare data by id and user id
func (svc *Service) GetSingleSkincare(userID, ID uint64) (model.Skincare, error) {
	var mSkincare model.Skincare

	// database query, use First to get single data. First equal ASC LIMIT 1
	if err := svc.db.Table(mSkincare.TableName()).
		Where("id = ? and user_id = ?", ID, userID).
		First(&mSkincare).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = shared.ErrDataNotFound
		}
		return mSkincare, err
	}

	return mSkincare, nil
}

// CreateSkincare : to create skincare data
func (svc *Service) CreateSkincare(mUser model.User, data form.SkincareForm) (model.Skincare, error) {
	var mSkincare model.Skincare

	// assign data from form to model
	mSkincare = model.Skincare{
		UserID:      mUser.ID,
		Brand:       data.Brand,
		Name:        data.Name,
		ProductType: data.ProductType,
		SkinConcern: data.SkinConcern,
	}

	// database query, insert data to database
	if err := svc.db.Table(mSkincare.TableName()).
		Create(&mSkincare).Error; err != nil {
		return mSkincare, err
	}

	return mSkincare, nil
}

// UpdateSkincare : to update existing skincare data
func (svc *Service) UpdateSkincare(mSkincare model.Skincare, data form.SkincareForm) (model.Skincare, error) {
	mSkincare.Brand = data.Brand
	mSkincare.Name = data.Name
	mSkincare.ProductType = data.ProductType
	mSkincare.SkinConcern = data.SkinConcern

	// database query, update existing data
	if err := svc.db.Table(mSkincare.TableName()).
		Save(&mSkincare).Error; err != nil {
		return mSkincare, err
	}

	return mSkincare, nil
}

// DeleteSkincare : to delete existing skincare data
func (svc *Service) DeleteSkincare(mSkincare model.Skincare) error {

	// database query, delete existing data
	if err := svc.db.Table(mSkincare.TableName()).
		Delete(&mSkincare).Error; err != nil {
		return err
	}

	return nil
}
