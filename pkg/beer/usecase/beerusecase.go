package usecase

import (
	_beerModel "github.com/PatiponKB/backend-test/pkg/beer/model"
)

type BeerUsecase interface {
	Create(beerItem *_beerModel.BeerCreatRequest)(*_beerModel.Beer, error)
	List(beerFilter *_beerModel.BeerFilter) (*_beerModel.Result, error)
	Delete(beerID uint64) error
	Update(beerID uint64, beerUpdategRequest *_beerModel.BeerUpdategRequest) (*_beerModel.Beer, error)
}