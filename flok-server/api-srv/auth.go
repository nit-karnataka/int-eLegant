package main

import (
	"context"
	authProto "flok-server/auth-srv/authproto"
	userProto "flok-server/user-srv/userproto"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type newuser struct {
	Password string `json:"password"`
}

type newstudent struct {
	*userProto.Student
	Password string `json:"password"`
}

type newpi struct {
	*userProto.PI
	Password string `json:"password"`
}

type newec struct {
	*userProto.EC
	Password string `json:"password"`
}

func (a *app) registerStudent(c echo.Context) error {
	s := &newstudent{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := c.Bind(s); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "profile must be valid",
		})
	}

	_, err := a.userClient.CreateStudent(ctx, &userProto.CreateStudentRequest{
		Student: &userProto.Student{
			Address:       s.Address,
			Doj:           s.Doj,
			Email:         s.Email,
			Gender:        s.Gender,
			Name:          s.Name,
			PhoneNumber:   s.PhoneNumber,
			TypeOfStudent: s.TypeOfStudent,
		},
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	log.Println("Student created")

	_, err = a.authClient.Create(ctx, &authProto.CreateRequest{
		User: &authProto.User{
			Email:       s.Email,
			AccountType: "student",
			Password:    s.Password,
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) registerPI(c echo.Context) error {
	pi := &newpi{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := c.Bind(pi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "profile must be valid",
		})
	}

	_, err := a.userClient.CreatePI(ctx, &userProto.CreatePIRequest{
		Pi: &userProto.PI{
			Address:     pi.Address,
			Email:       pi.Email,
			Gender:      pi.Gender,
			Name:        pi.Name,
			PhoneNumber: pi.PhoneNumber,
		},
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	log.Println("PI created")

	_, err = a.authClient.Create(ctx, &authProto.CreateRequest{
		User: &authProto.User{
			Email:       pi.Email,
			AccountType: "pi",
			Password:    pi.Password,
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) registerEC(c echo.Context) error {
	ec := &newec{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := c.Bind(ec); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "profile must be valid",
		})
	}

	_, err := a.userClient.CreateEC(ctx, &userProto.CreateECRequest{
		Ec: &userProto.EC{
			Affiliation:         ec.Affiliation,
			CurrentOrganization: ec.CurrentOrganization,
			Designation:         ec.Designation,
			Email:               ec.Email,
			Fax:                 ec.Fax,
			Gender:              ec.Gender,
			Name:                ec.Name,
			Position:            ec.Position,
			Qualification:       ec.Qualification,
			Telephone:           ec.Telephone,
		},
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	log.Println("Ec created")

	_, err = a.authClient.Create(ctx, &authProto.CreateRequest{
		User: &authProto.User{
			Email:       ec.Email,
			AccountType: "ec",
			Password:    ec.Password,
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) register(c echo.Context) error {
	p := &newuser{}

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "profile must be valid",
		})
	}

	/* if p.DisplayName == "" {
		p.DisplayName = p.Name
	}

	if p.Address == "" ||
		p.DisplayName == "" ||
		p.Name == "" ||
		!valid.IsEmail(p.Email) ||
		!valid.IsNumeric(p.PhoneNumber) ||
		!valid.IsByteLength(p.PhoneNumber, 10, 10) ||
		!valid.IsByteLength(p.Password, 8, 36) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "all fields must be valid",
		})
	}

	return echo.ErrForbidden */
	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "all fields must be valid",
	})
}

func (a *app) login(c echo.Context) error {
	u := &user{}

	err := c.Bind(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message":  "email and password must",
			"email":    u.Email,
			"password": u.Password,
		})
	}

	if u.Email == "" || u.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message":  "email and password must",
			"email":    u.Email,
			"password": u.Password,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	authRes, err := a.authClient.VerifyUser(ctx, &authProto.VerifyUserRequest{
		Email:    u.Email,
		Password: u.Password,
	})

	if err == nil && authRes != nil {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["role"] = authRes.User.AccountType
		// claims["role"] = "super"
		claims["id"] = u.Email
		// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(a.c.JwtSecret))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
			"role":  authRes.User.AccountType,
		})
	}
	return echo.ErrUnauthorized
}

func getUserFromContext(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
