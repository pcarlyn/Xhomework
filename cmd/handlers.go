package cmd

import (
	"encoding/json"
	"homework/web/internal/models"
	"homework/web/internal/utils"
	"net/http"

	"github.com/labstack/echo"
)

var items []models.Item = utils.NewItems()

func WriteItem(c echo.Context) error {
	response := new(models.Item)
	err := json.NewDecoder(c.Request().Body).Decode(response)
	if err != nil {
		return err
	}
	utils.AddItem(&items, response)
	return c.JSON(http.StatusOK, "Saved")
}

func ReadItem(c echo.Context) error {
	response := c.Param("caption")
	item := utils.SearchByCaption(items, response)
	if item == nil {
		return c.JSON(http.StatusOK, "Not Found")
	}
	return c.JSON(http.StatusOK, item)
}
