package main

import (
	"context"
	formProto "flok-server/form-srv/formproto"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (a *app) getForm(c echo.Context) error {
	// claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, e := a.formClient.View(ctx, &formProto.ViewRequest{
		Id: c.Param("id"),
	})

	log.Printf("FORM: %+v", res.Form)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": e.Error(),
		})
	}

	if res.Form.ProtocolForm != nil {
		return c.JSON(http.StatusOK, res.Form.ProtocolForm)

	}

	return c.NoContent(http.StatusOK)
}

func (a *app) updateForm(c echo.Context) error {
	// claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	formRes, err := a.formClient.View(ctx, &formProto.ViewRequest{
		Id: c.Param("id"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	if formRes.Form.ProtocolForm != nil {
		log.Printf("FORM: %+v", formRes.Form)
		err = c.Bind(formRes.Form.ProtocolForm)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
	}

	_, err = a.formClient.Update(ctx, &formProto.UpdateRequest{
		Form: formRes.Form,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "OK",
	})
}
