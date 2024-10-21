package repository

import (
	"github.com/PatiponKB/backend-test/entities"
	_beerModel "github.com/PatiponKB/backend-test/pkg/beer/model"
	"gorm.io/gorm"
)

type BeerRepository interface {
	Create(beerItem *entities.Beer) (*entities.Beer, error)
	List(beerFilter *_beerModel.BeerFilter) ([]*entities.Beer, error)
	Count(beerFilter *_beerModel.BeerFilter) (int64, error)
	Delete(beerID uint64) error
	FindByID(beerID uint64) (*entities.Beer, error)
	Update(beerID uint64, beerUpdategRequest *_beerModel.BeerUpdategRequest) (uint64, error)
}

type beerRepository struct {
	db *gorm.DB
}

func NewBeerRepository(db *gorm.DB) BeerRepository {
	return &beerRepository{db}
}

func (r *beerRepository) Create(beerItem *entities.Beer) (*entities.Beer, error) {
	beer := new(entities.Beer)
	if err := r.db.Create(beerItem).Scan(beer).Error; err != nil {
		return nil, err
	}
	return beer, nil
}

func (r *beerRepository) List(beerFilter *_beerModel.BeerFilter) ([]*entities.Beer, error) {
	beerList := make([]*entities.Beer, 0)
	query := r.db.Model(&entities.Beer{})
	if beerFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+beerFilter.Name+"%")
	}
	if beerFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+beerFilter.Description+"%")
	}
	offset := int((beerFilter.Page - 1) * beerFilter.Size)
	size := int(beerFilter.Size)
	query.Offset(offset).Limit(size).Find(&beerList)
	return beerList, nil
}

func (r *beerRepository) Count(beerFilter *_beerModel.BeerFilter) (int64, error) {
	query := r.db.Model(&entities.Beer{})
	if beerFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+beerFilter.Name+"%")
	}
	if beerFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+beerFilter.Description+"%")
	}
	var count int64
	query.Count(&count)

	return count, nil
}

func (r *beerRepository) Delete(beerID uint64) error {
	if err := r.db.Table("beers").Where("id = ?", beerID).Delete(&entities.Beer{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *beerRepository) FindByID(beerID uint64) (*entities.Beer, error) {
	beer := new(entities.Beer)

	if err := r.db.First(beer, beerID).Error; err != nil {
		return nil, err
	}

	return beer, nil
}


func (r *beerRepository) Update(beerID uint64, beerUpdategRequest *_beerModel.BeerUpdategRequest) (uint64, error) {
	if err := r.db.Model(&entities.Beer{}).Where("id = ?",beerID).Updates(beerUpdategRequest).Error;err != nil {
		return 0, err
	}
	return beerID, nil
}