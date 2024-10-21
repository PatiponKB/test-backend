package handler

import (
	"net/http"
	"strconv"

	"github.com/PatiponKB/backend-test/pkg/beer/model"
	"github.com/PatiponKB/backend-test/pkg/beer/usecase"
	"github.com/labstack/echo/v4"
)

type beerHandler struct {
	beerUsecase usecase.BeerUsecase
}


func NewBeerHandler(beerUsecase usecase.BeerUsecase) BeerHandler {
	return &beerHandler{beerUsecase}
}

func (h *beerHandler) List(ctx echo.Context) error {
	beerFilter := new(model.BeerFilter)

	if err := ctx.Bind(beerFilter); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	beerList, err := h.beerUsecase.List(beerFilter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, beerList)
}

func (h *beerHandler) Create(ctx echo.Context) error {
	beerCreatReq := new(model.BeerCreatRequest)

	if err := ctx.Bind(beerCreatReq); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	beer, err := h.beerUsecase.Create(beerCreatReq)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, beer)
}


func (h *beerHandler) Delete(ctx echo.Context) error {
	beerID, err := h.getID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = h.beerUsecase.Delete(beerID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *beerHandler) getID(ctx echo.Context) (uint64, error) {
	beerID := ctx.Param("beerID")
	beerIDUint64, err := strconv.ParseUint(beerID, 10, 64)
	if err != nil {
		return 0, err
	}

	return beerIDUint64, nil
}

func (h *beerHandler) Update(ctx echo.Context) error {
	beerID, err := h.getID(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	updateReq := new(model.BeerUpdategRequest)
	if err := ctx.Bind(updateReq); err != nil {
		return ctx.String(http.StatusBadRequest,err.Error())
	}

	beer, err := h.beerUsecase.Update(beerID,updateReq)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK,beer)

}

