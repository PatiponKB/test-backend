package usecase

import (
	"github.com/PatiponKB/backend-test/entities"
	_beerModel "github.com/PatiponKB/backend-test/pkg/beer/model"
)

func (u *beerUsecase) Create(beerRequest *_beerModel.BeerCreatRequest) (*_beerModel.Beer, error) {

	beer := &entities.Beer{
		Name:        beerRequest.Name,
		Type:        beerRequest.Type,
		Description: beerRequest.Description,
		Picture:     beerRequest.Picture,
	}
	res, err := u.beerRepository.Create(beer)
	if err != err {
		return nil, err
	}
	return res.ToModel(), nil

}