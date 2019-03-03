package main

import (
	"context"
	chatProto "flok-server/chat-srv/chatproto"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (a *app) getChat(c echo.Context) error {
	claims := getUserFromContext(c)
	uID := claims["id"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, e := a.chatClient.View(ctx, &chatProto.ViewRequest{
		ChatID: c.Param("id"),
		UserID: uID,
	})

	if e != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, res.Chat)
}

func (a *app) addChat(c echo.Context) error {
	claims := getUserFromContext(c)
	uID := claims["id"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, e := a.chatClient.InsertComment(ctx, &chatProto.InsertCommentRequest{
		ChatID:  c.Param("id"),
		Owner:   uID,
		Content: c.QueryParam("content"),
	})

	if e != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "OK",
	})
}
