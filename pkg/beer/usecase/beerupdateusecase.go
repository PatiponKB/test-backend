package usecase

import _beerModel "github.com/PatiponKB/backend-test/pkg/beer/model"

func (u *beerUsecase) Update(beerID uint64, beerUpdategRequest *_beerModel.BeerUpdategRequest) (*_beerModel.Beer, error) {
	_ ,err := u.beerRepository.Update(beerID, beerUpdategRequest)
	if err != nil {
		return nil, err
	}

	beerUpdategResult ,err := u.beerRepository.FindByID(beerID)
	if err != nil {
		return nil, err
	}

	return beerUpdategResult.ToModel(), nil
}
