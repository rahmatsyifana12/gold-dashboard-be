package repositories

import (
	"context"
	"gold-dashboard-be/internal/constants"
	"gold-dashboard-be/internal/models"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type GoldAssetRepository interface {
	CreateGoldAsset(ctx context.Context, goldAsset models.GoldAsset) (err error)
	GetGoldAssetsByUserID(ctx context.Context, userID uint) (goldAssets []models.GoldAsset, err error)
}

type GoldAssetRepositoryImpl struct {
	db *gorm.DB
}

func NewGoldAssetRepository(ioc di.Container) *GoldAssetRepositoryImpl {
	return &GoldAssetRepositoryImpl{
		db: ioc.Get(constants.Postgres).(*gorm.DB),
	}
}

func (r *GoldAssetRepositoryImpl) CreateGoldAsset(ctx context.Context, goldAsset models.GoldAsset) (err error) {
	err = r.db.Create(&goldAsset).WithContext(ctx).Error
	return
}

func (r *GoldAssetRepositoryImpl) GetGoldAssetsByUserID(ctx context.Context, userID uint) (goldAssets []models.GoldAsset, err error) {
	err = r.db.Where("user_id = ?", userID).Find(&goldAssets).WithContext(ctx).Error
	return
}