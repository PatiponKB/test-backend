package usecase

func (u *beerUsecase) Delete(beerID uint64) error {
	return u.beerRepository.Delete(beerID)
}