package home

import (
	"github.com/johnazedo/playmovie-bff/src/hermes"
	. "github.com/johnazedo/playmovie-bff/src/home/domain"
	. "github.com/johnazedo/playmovie-bff/src/home/infra"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type HomeController struct {
	GetHomeUseCase
}

func NewHomeController() HomeController {
	return HomeController{
		GetHomeUseCase{
			MovieRepositoryImpl{Service: hermes.NewClient()},
			CatalogRepositoryImpl{},
		},
	}
}

func (ctrl *HomeController) GetHome(c echo.Context) error {
	log.Println("Get home")
	trails, err := ctrl.GetHomeUseCase.Execute()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, trails)
}
