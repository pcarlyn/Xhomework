package cmd

import "github.com/labstack/echo"

func Route(group *echo.Group) {
	group.GET("item/:caption", ReadItem)
	group.POST("item", WriteItem)
}
