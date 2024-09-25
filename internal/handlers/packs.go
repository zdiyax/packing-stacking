package handlers

import (
	"net/http"
	"packing-stacking/internal/domain"
	"strconv"

	"packing-stacking/internal/service"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Add(c echo.Context) error
	List(c echo.Context) error
	Remove(c echo.Context) error
	Calculate(c echo.Context) error
	Index(c echo.Context) error
}

type PacksHandler struct {
	svc service.PacksService
}

func NewPacksHandler() Handler {
	return PacksHandler{
		svc: service.NewPacksService(),
	}
}

// @Summary Add pack
// @Description Add a new pack with unique quantity
// @ID add-pack
// @Param quantity path int true "Quantity"
// @Success 200
// @Router /packs/add/{quantity} [get]
func (h PacksHandler) Add(c echo.Context) error {
	quantity := c.QueryParam("quantity")

	i, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}

	err = h.svc.Add(domain.Pack{Quantity: i})
	if err != nil {
		return err
	}

	packs, err := h.svc.List()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "index.html", packs)
}

// @Summary List packs
// @Description List all available packs
// @ID list-packs
// @Success 200 {array} []domain.Pack
// @Router /packs/list [get]
func (h PacksHandler) List(c echo.Context) error {
	packs, err := h.svc.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &packs)
}

// @Summary Remove pack
// @Description Removes a pack with a certain quantity
// @ID remove-pack
// @Param quantity path int true "Quantity"
// @Success 200
// @Router /packs/remove/{quantity} [get]
func (h PacksHandler) Remove(c echo.Context) error {
	quantity := c.QueryParam("quantity")

	i, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}

	err = h.svc.Remove(domain.Pack{Quantity: i})
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, quantity)

}

// @Summary Calculate packs
// @Description Calculates the needed packs combination for a given quantity
// @ID calculate-packs
// @Success 200 {object} map[domain.Pack]int
// @Router /packs/calculate/{quantity} [get]
func (h PacksHandler) Calculate(c echo.Context) error {
	quantityParam := c.QueryParam("quantity")

	quantity, err := strconv.Atoi(quantityParam)
	if err != nil {
		return err
	}

	calculation, err := h.svc.Calculate(quantity)
	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	}

	return c.Render(http.StatusOK, "calculation.html", calculation)
}

func (h PacksHandler) Index(c echo.Context) error {
	packs, err := h.svc.List()
	if err != nil {
		return c.Render(http.StatusConflict, "index.html", nil)
	}

	return c.Render(http.StatusOK, "index.html", packs)
}
