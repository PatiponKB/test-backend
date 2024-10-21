package usecase

import (
	"github.com/PatiponKB/backend-test/entities"
	_beerModel "github.com/PatiponKB/backend-test/pkg/beer/model"
	"github.com/PatiponKB/backend-test/repository"
)

type beerUsecase struct {
	beerRepository repository.BeerRepository
}


func NewBeerUsecase(beerRepository repository.BeerRepository) BeerUsecase {
	return &beerUsecase{beerRepository}
}

func (u *beerUsecase) List(beerFilter *_beerModel.BeerFilter) (*_beerModel.Result, error) {
	beerList, err := u.beerRepository.List(beerFilter)
	if err != nil {
		return nil, err
	}

	totalCounting, err := u.beerRepository.Count(beerFilter)
	if err != nil {
		return nil, err
	}

	totalPage := u.totalPageCal(totalCounting, beerFilter.Size)
	result := u.toResult(beerList, beerFilter.Page, totalPage)

	return result, nil
}

func (u *beerUsecase) totalPageCal(totalBeer, size int64) int64 {
	totalPage := totalBeer / size

	if totalPage%size != 0 {
		totalPage++
	}

	return totalPage
}

func (u *beerUsecase) toResult(beerEntitiesList []*entities.Beer, page, totalPage int64) *_beerModel.Result {
	beerModelList := make([]*_beerModel.Beer, 0)

	for _, beer := range beerEntitiesList {
		beerModelList = append(beerModelList, beer.ToModel())
	}

	return &_beerModel.Result{
		Beer: beerModelList,
		Paginate: _beerModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
