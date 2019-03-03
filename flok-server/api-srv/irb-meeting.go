package main

import (
	"context"
	fileProto "flok-server/file-srv/fileproto"
	meetingProto "flok-server/meeting-srv/meetingproto"
	projectProto "flok-server/project-srv/projectproto"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func (a *app) testFileUpload(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	fileName := c.QueryParam("fileName")
	ext := c.QueryParam("ext")

	file, err := c.FormFile("test")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// sID := ""
	// sID := claims["id"]

	fileRes, err := a.fileClient.Create(ctx, &fileProto.CreateRequest{
		Data:      b,
		Name:      fileName,
		Extension: ext,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": fileRes.Id,
	})
}

func (a *app) getIRBMeeting(c echo.Context) error {
	// claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// sID := ""
	// sID := claims["id"]

	meetingRes, err := a.meetingClient.View(ctx, &meetingProto.ViewRequest{
		Id: c.Param("irbmeeting"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, meetingRes.Meeting)
}

func (a *app) addIRBMeetingInitailMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.AddInitialMembers(ctx, &meetingProto.AddInitialMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) removeIRBMeetingInitailMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.RemoveInitialMembers(ctx, &meetingProto.RemoveInitialMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) setIRBMeetingAgenda(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a sec is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	fileName := c.QueryParam("fileName")
	ext := c.QueryParam("ext")

	file, err := c.FormFile("agenda")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// sID := ""
	// sID := claims["id"]

	fileRes, err := a.fileClient.Create(ctx, &fileProto.CreateRequest{
		Data:      b,
		Name:      fileName,
		Extension: ext,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = a.meetingClient.SetAgenda(ctx, &meetingProto.SetAgendaRequest{
		Doc: fileRes.Id,
		Id:  c.Param("irbmeeting"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) setIRBMeetingMinutes(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a sec is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	fileName := c.QueryParam("fileName")
	ext := c.QueryParam("ext")

	file, err := c.FormFile("minutes")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// sID := ""
	// sID := claims["id"]

	fileRes, err := a.fileClient.Create(ctx, &fileProto.CreateRequest{
		Data:      b,
		Name:      fileName,
		Extension: ext,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = a.meetingClient.SetMinutes(ctx, &meetingProto.SetMinutesRequest{
		Doc: fileRes.Id,
		Id:  c.Param("irbmeeting"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) addIRBMeetingPresentMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.AddPresentMembers(ctx, &meetingProto.AddPresentMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) removeIRBMeetingPresentMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.RemovePresentMembers(ctx, &meetingProto.RemovePresentMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) addIRBMeetingQLMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.AddQLMembers(ctx, &meetingProto.AddQLMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) removeIRBMeetingQLMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.RemoveQLMembers(ctx, &meetingProto.RemoveQLMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) approveIRBMeetingAgenda(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "dor" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "only account with DOR role allowed",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := a.meetingClient.ApproveAgenda(ctx, &meetingProto.ApproveAgendaRequest{
		Approve: isApproved,
		Id:      c.Param("irbmeeting"),
		Comment: c.QueryParam("rejectionComment"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"msg": "Ok"})
}

func (a *app) approveIRBMeetingMinute(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "dor" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "only account with DOR role allowed",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := a.meetingClient.ApproveMinute(ctx, &meetingProto.ApproveMinuteRequest{
		Approve: isApproved,
		Id:      c.Param("irbmeeting"),
		Comment: c.QueryParam("rejectionComment"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"msg": "Ok"})
}

func (a *app) addIRBMeetingFAMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.AddFAMembers(ctx, &meetingProto.AddFAMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) removeIRBMeetingFAMembers(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	members := []string{}

	err := c.Bind(members)
	if err != nil || len(members) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err = a.meetingClient.RemoveFAMembers(ctx, &meetingProto.RemoveFAMembersRequest{
		Id:      c.Param("irbmeeting"),
		Members: members,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) setIRBMeetingChair(c echo.Context) error {
	claims := getUserFromContext(c)
	// uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "not enough permissions",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := a.meetingClient.SetChairPerson(ctx, &meetingProto.SetChairPersonRequest{
		Id:   c.Param("irbmeeting"),
		User: c.QueryParam("chairPerson"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (a *app) addIRBMeetingQLR(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	projectRes, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if projectRes.Project.Student != claims["id"].(string) {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "no permissions",
		})
	}

	fileName := c.QueryParam("fileName")
	ext := c.QueryParam("ext")

	file, err := c.FormFile("qlr")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// sID := ""
	// sID := claims["id"]

	fileRes, err := a.fileClient.Create(ctx, &fileProto.CreateRequest{
		Data:      b,
		Name:      fileName,
		Extension: ext,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = a.meetingClient.AddQLR(ctx, &meetingProto.AddQLRRequest{
		Id:  c.Param("irbmeeting"),
		Qlr: fileRes.Id,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) removeIRBMeetingQLR(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	projectRes, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if projectRes.Project.Student != claims["id"].(string) {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "no permissions",
		})
	}

	// sID := ""
	// sID := claims["id"]

	_, err = a.meetingClient.RemoveQLR(ctx, &meetingProto.RemoveQLRRequest{
		Id:  c.Param("irbmeeting"),
		Qlr: c.Param("qlr"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) aprroveIRBMeetingQLRPI(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "pi" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a PI is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	projectRes, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if projectRes.Project.PrincipalInvestigator != claims["id"].(string) {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "no permissions",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	_, err = a.meetingClient.ApproveQLRPI(ctx, &meetingProto.ApproveQLRPIRequest{
		Id:         c.Param("irbmeeting"),
		IsApproved: isApproved,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) aprroveIRBMeetingQLRSec(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a Sec is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	_, err := a.meetingClient.ApproveQLRSec(ctx, &meetingProto.ApproveQLRSecRequest{
		Id:         c.Param("irbmeeting"),
		IsApproved: isApproved,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) setIRBMeetingTranslation(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a Sec is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	_, err := a.meetingClient.SetTranslation(ctx, &meetingProto.SetTranslationRequest{
		Id: c.Param("irbmeeting"),
		T:  isApproved,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) addIRBMeetingTranslation(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	projectRes, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if projectRes.Project.Student != claims["id"].(string) {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "no permissions",
		})
	}

	fileName := c.QueryParam("fileName")
	ext := c.QueryParam("ext")

	file, err := c.FormFile("translation")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// sID := ""
	// sID := claims["id"]

	fileRes, err := a.fileClient.Create(ctx, &fileProto.CreateRequest{
		Data:      b,
		Name:      fileName,
		Extension: ext,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = a.meetingClient.AddTranslation(ctx, &meetingProto.AddTranslationRequest{
		Id:          c.Param("irbmeeting"),
		Translation: fileRes.Id,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) removeIRBMeetingTranslation(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	projectRes, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if projectRes.Project.Student != claims["id"].(string) {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "no permissions",
		})
	}

	// sID := ""
	// sID := claims["id"]

	_, err = a.meetingClient.RemoveTranslation(ctx, &meetingProto.RemoveTranslationRequest{
		Id:          c.Param("irbmeeting"),
		Translation: c.Param("translation"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) aprroveIRBMeetingTranslationPI(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "pi" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a PI is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	projectRes, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if projectRes.Project.PrincipalInvestigator != claims["id"].(string) {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "no permissions",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	_, err = a.meetingClient.ApproveTranslationPI(ctx, &meetingProto.ApproveTranslationPIRequest{
		Id:         c.Param("irbmeeting"),
		IsApproved: isApproved,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) aprroveIRBMeetingTranslationSec(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a Sec is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	_, err := a.meetingClient.ApproveTranslationSec(ctx, &meetingProto.ApproveTranslationSecRequest{
		Id:         c.Param("irbmeeting"),
		IsApproved: isApproved,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) setIRBMeetingQLReview(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "ec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a EC is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	status, _ := strconv.ParseBool(c.QueryParam("status"))

	_, err := a.meetingClient.SetQLReview(ctx, &meetingProto.SetQLReviewRequest{
		Id:               c.Param("irbmeeting"),
		Member:           claims["id"].(string),
		RejectionComment: c.QueryParam("rejectionComment"),
		Status:           status,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) setIRBMeetingTranslationReview(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "ec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a EC is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	status, _ := strconv.ParseBool(c.QueryParam("status"))

	_, err := a.meetingClient.SetTranslationReview(ctx, &meetingProto.SetTranslationReviewRequest{
		Id:               c.Param("irbmeeting"),
		Member:           claims["id"].(string),
		RejectionComment: c.QueryParam("rejectionComment"),
		Status:           status,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

func (a *app) setIRBMeetingFA(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "ec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a EC is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	status, _ := strconv.ParseBool(c.QueryParam("status"))

	r, err := a.meetingClient.SetFA(ctx, &meetingProto.SetFARequest{
		Id:     c.Param("irbmeeting"),
		Member: claims["id"].(string),
		Status: status,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	if r.IsApproved {
		_, err = a.projectClient.SetIsApproved(ctx, &projectProto.SetIsApprovedRequest{
			Id:     c.Param("id"),
			Status: true,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}
