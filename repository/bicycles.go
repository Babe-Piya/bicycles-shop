package repository

import (
	"bicycles-shop/model"
	"context"
	"gorm.io/gorm"
)

type BicyclesRepository interface {
	GetListBicycles(ctx context.Context) ([]model.Bicycles, error)
	UpdateBuyStatusBicycle(ctx context.Context, id int) error
	GetBicycleCanBuy(ctx context.Context, id int) (model.Bicycles, error)
	InsertBuyer(ctx context.Context, buyers model.Buyers) error
	InsertBicycle(bicycle model.Bicycles) (model.Bicycles, error)
}

type bicyclesRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) BicyclesRepository {
	return &bicyclesRepository{
		db: db,
	}
}

func (r *bicyclesRepository) GetListBicycles(ctx context.Context) ([]model.Bicycles, error) {
	var listBicycles []model.Bicycles
	result := r.db.WithContext(ctx).Table("bicycles").Find(&listBicycles)
	if result.Error != nil {
		return []model.Bicycles{}, result.Error
	}
	return listBicycles, nil
}

func (r *bicyclesRepository) UpdateBuyStatusBicycle(ctx context.Context, id int) error {
	var buyBicycle model.Bicycles
	err := r.db.WithContext(ctx).Model(&buyBicycle).Where("id = ?", id).Update("status", model.BUY).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bicyclesRepository) GetBicycleCanBuy(ctx context.Context, id int) (model.Bicycles, error) {
	var bicycle model.Bicycles
	result := r.db.WithContext(ctx).Table("bicycles").Where("id = ? AND status = ?", id, model.SELL).First(&bicycle)
	if result.Error != nil {
		return model.Bicycles{}, result.Error
	}

	return bicycle, nil
}

func (r *bicyclesRepository) InsertBuyer(ctx context.Context, buyers model.Buyers) error {
	err := r.db.WithContext(ctx).Create(&buyers).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *bicyclesRepository) InsertBicycle(bicycle model.Bicycles) (model.Bicycles, error) {
	err := r.db.Create(&bicycle).Error
	if err != nil {
		return model.Bicycles{}, err
	}

	return bicycle, nil
}
