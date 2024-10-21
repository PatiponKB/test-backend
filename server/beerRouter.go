package server

import (
	"github.com/PatiponKB/backend-test/pkg/beer/handler"
	"github.com/PatiponKB/backend-test/pkg/beer/usecase"
	"github.com/PatiponKB/backend-test/repository"
)

func (s *echoServer) initBeerRouter() {
	router := s.app.Group("/v1/beer")

	beerRepository := repository.NewBeerRepository(s.db)
	beerUsecase := usecase.NewBeerUsecase(beerRepository)
	beerHandler := handler.NewBeerHandler(beerUsecase)

	router.GET("", beerHandler.List)
	router.POST("", beerHandler.Create)
	router.PUT("/:beerID", beerHandler.Update)
	router.DELETE("/:beerID", beerHandler.Delete)
}
