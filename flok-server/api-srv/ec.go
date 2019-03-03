package main

import (
	"context"
	userProto "flok-server/user-srv/userproto"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (a *app) getAllECMembers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := a.userClient.GetAllEC(ctx, &userProto.GetAllRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res.Ecs)
}
