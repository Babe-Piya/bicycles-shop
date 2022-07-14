package repository

import (
	"bicycles-shop/model"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type BicyclesRepository interface {
	GetListBicycles(ctx context.Context) ([]model.Bicycles, error)
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
	err := r.db.WithContext(ctx).Table("bicycles").Find(&listBicycles).Error
	fmt.Println(listBicycles)
	if err != nil {
		return []model.Bicycles{}, err
	}
	return listBicycles, nil
}
