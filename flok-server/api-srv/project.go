package main

import (
	"context"
	chatProto "flok-server/chat-srv/chatproto"
	meetingProto "flok-server/meeting-srv/meetingproto"
	portalProto "flok-server/portal-srv/portalproto"
	projectProto "flok-server/project-srv/projectproto"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/thoas/go-funk"

	"github.com/labstack/echo"
)

func (a *app) getAllProjects(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := a.projectClient.GetAll(ctx, &projectProto.GetAllRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res.Projects)
}

func (a *app) getAllStudentProjects(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	log.Printf("claims: %+v", claims)

	// sID := ""
	sID := claims["id"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := a.projectClient.GetByStudent(ctx, &projectProto.GetByStudentRequest{
		Student: sID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res.Projects)
}

func (a *app) getProject(c echo.Context) error {
	claims := getUserFromContext(c)
	log.Printf("claims: %+v", claims)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	// sID := ""
	// sID := claims["id"]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := a.projectClient.Read(ctx, &projectProto.ReadRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res.Project)
}

func (a *app) createProject(c echo.Context) error {
	project := &projectProto.Project{}

	err := c.Bind(project)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = time.ParseDuration(project.Period)
	if err != nil || project.Period == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "project period wrong",
		})
	}

	if project.Title == "" || project.PrincipalInvestigator == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "project tille and PI required",
		})
	}

	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	// sID := ""
	sID := claims["id"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	secID := ""
	dor := ""

	owners := []string{
		// secID,
		// dor,
		sID,
		project.PrincipalInvestigator,
	}

	chatRes, err := a.chatClient.Create(ctx, &chatProto.CreateRequest{
		Owner:    append(owners, project.CoInvestigators...),
		IsFreeze: false,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	project.PiChatID = chatRes.Id

	owners = append(
		append(
			owners, project.CoInvestigators...,
		), secID,
	)

	chatRes, err = a.chatClient.Create(ctx, &chatProto.CreateRequest{
		Owner:    owners,
		IsFreeze: false,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	project.SecChatID = chatRes.Id

	chatRes, err = a.chatClient.Create(ctx, &chatProto.CreateRequest{
		Owner: append(
			owners, dor,
		),
		IsFreeze: false,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	project.DorChatID = chatRes.Id

	portalRes, err := a.portalClient.Create(ctx, &portalProto.CreateRequest{
		IsFreeze:  false,
		PiID:      project.PrincipalInvestigator,
		StudentID: sID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	meetingRes, err := a.meetingClient.Create(ctx, &meetingProto.CreateRequest{
		ChairPerson: "",
	})
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": err.Error(),
		})
	}

	project.IrbMeetingID = ""
	project.ThesisPortalID = portalRes.Id
	project.IrbMeetingID = meetingRes.Id

	project.Progress = int32(0)
	project.Sec = secID
	project.Dor = dor
	project.CreatedAt = time.Now().Format(time.RFC3339)
	project.Documents = []string{}
	project.Dotc = ""
	project.Dots = ""
	project.Doa = ""
	project.IsApproved = false
	project.IsIhesisDone = false
	project.IsIRBMeetingDone = false
	project.IsCompleted = false
	project.Student = sID

	projectRes, err := a.projectClient.Create(ctx, &projectProto.CreateRequest{
		Project: project,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id": projectRes.Id,
	})
}

/* func (a *app) uploadThesis(c echo.Context) error {
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

	fileName := c.QueryParam("fileName")
	ext := c.QueryParam("ext")

	file, err := c.FormFile("thesis")
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

	sID := ""
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

	_, err = a.portalClient.AddDocument(ctx, &portalProto.AddDocumentRequest{
		DocumentID: fileRes.Id,
		StudentID:  sID,
		PortalID:   projectRes.Project.ThesisPortalID,
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

func (a *app) removeThesis(c echo.Context) error {
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

	sID := ""
	// sID := claims["id"]

	_, err = a.portalClient.RemoveDocument(ctx, &portalProto.RemoveDocumentRequest{
		DocumentID: c.Param("thesis"),
		StudentID:  sID,
		PortalID:   projectRes.Project.ThesisPortalID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
} */

func (a *app) setThesisSubmitted(c echo.Context) error {
	claims := getUserFromContext(c)
	if claims["role"].(string) != "student" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "only a student is allowed to access this API",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	isSubmitted, _ := strconv.ParseBool(c.QueryParam("isSubmitted"))

	_, err := a.portalClient.SetSubmitted(ctx, &portalProto.SetSubmittedRequest{
		Id:          c.Param("thesis"),
		User:        claims["id"].(string),
		IsSubmitted: isSubmitted,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": "OK",
	})
}

/* func (a *app) getThesis(c echo.Context) error {
	uID := getUserFromContext(c)["id"].(string)

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

	// sID := ""
	// sID := claims["id"]

	portalRes, err := a.portalClient.View(ctx, &portalProto.ViewRequest{
		UserID:   uID,
		PortalID: projectRes.Project.ThesisPortalID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, portalRes.Portal)
} */

func (a *app) getThesis(c echo.Context) error {
	claims := getUserFromContext(c)
	uID := claims["id"].(string)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// sID := ""
	// sID := claims["id"]

	portalRes, err := a.portalClient.View(ctx, &portalProto.ViewRequest{
		UserID:   uID,
		PortalID: c.Param("thesis"),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, portalRes.Portal)
}

func (a *app) approveThesisPI(c echo.Context) error {
	claims := getUserFromContext(c)
	uID := claims["id"].(string)

	if claims["role"].(string) != "pi" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "only account with PI role allowed",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	portalRes, err := a.portalClient.SetApprove(ctx, &portalProto.SetApproveRequest{
		IsApproved:       isApproved,
		PortalID:         c.Param("thesis"),
		RejectionComment: c.QueryParam("rejectionComment"),
		UserID:           uID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	if !portalRes.Success {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "insufficient permissions",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"msg": "Ok"})
}

func (a *app) approveThesisSec(c echo.Context) error {
	claims := getUserFromContext(c)
	uID := claims["id"].(string)

	if claims["role"].(string) != "sec" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "only account with sec role allowed",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	portalRes, err := a.portalClient.SetApprove(ctx, &portalProto.SetApproveRequest{
		IsApproved:       isApproved,
		PortalID:         c.Param("thesis"),
		RejectionComment: c.QueryParam("rejectionComment"),
		UserID:           uID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	if !portalRes.Success {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "insufficient permissions",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"msg": "Ok"})
}

func (a *app) approveThesisDOR(c echo.Context) error {
	claims := getUserFromContext(c)
	uID := claims["id"].(string)

	if claims["role"].(string) != "dor" && claims["role"].(string) != "super" {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "only account with DOR role allowed",
		})
	}

	isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	portalRes, err := a.portalClient.SetApprove(ctx, &portalProto.SetApproveRequest{
		IsApproved:       isApproved,
		PortalID:         c.Param("thesis"),
		RejectionComment: c.QueryParam("rejectionComment"),
		UserID:           uID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	if !portalRes.Success {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "insufficient permissions",
		})
	}

	if isApproved {
		_, err = a.projectClient.SetThesisDone(ctx, &projectProto.SetThesisDoneRequest{
			Id: c.Param("id"),
		})
		if err != nil {
			return c.JSON(http.StatusForbidden, echo.Map{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"msg": "Ok"})
}

func (a *app) process(c echo.Context) error {
	db, err := a.s.GetMongoSession()
	if err != nil {
		return err
	}
	defer db.Close()

	id := c.Param("id")
	p := &Process{}

	fmt.Println(id)

	err = db.DB("").C(collectionName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(p)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	userID := c.QueryParam("u")

	step := &p.Steps[p.CurrStep-1]
	fmt.Printf("%+v\n%+v\n", step, p.CurrStep-1)
	if p.State == nil {
		p.State = map[int][]string{}
	}
	switch step.TypeOf {
	case 1:
		if step.Who == userID || userID == step.Who || (step.WhoID == -1 && funk.ContainsString(step.PrefilledMembers, userID)) || funk.ContainsString(p.State[step.WhoID-1], userID) {
			fmt.Printf("%+v\n", "called")
			step.Doc = c.QueryParam("doc")
			p.CurrStep = step.AcceptNext
			step.RejectionComments = []string{}
			step.RejectionComment = ""
			fmt.Printf("%+v\n", p.Steps[0])
		} else {
			fmt.Println("GO TO HELL", step.Who, step.PrefilledMembers, userID)
		}
	case 2:
		if step.Who == userID {
			isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))
			rejectionComment := c.QueryParam("comment")

			if isApproved {
				p.CurrStep = step.AcceptNext
			} else {
				p.Steps[step.RejectStep-1].RejectionComment = rejectionComment
				p.CurrStep = step.RejectStep
			}
		} else if funk.ContainsString(step.PrefilledMembers, userID) {
			isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))
			rejectionComment := c.QueryParam("comment")

			// member := ""
			cnt := 0

			hasAllSubmitted := false

			allMem := step.PrefilledMembers

			if step.WhoID != -1 {
				allMem = p.State[step.WhoID-1]
			}

			fmt.Println("lemsadasd", len(step.Approv), allMem)

			for idx, mem := range allMem {
				if len(step.Approv) < idx+1 {
					fmt.Println("ADDING", mem)
					step.Approv = append(step.Approv, Approval{
						Member:           mem,
						RejectionComment: "",
						Status:           false,
					})
				}
			}

			appr := &step.Approv

			for idx, mem := range *appr {
				if mem.Member == userID {
					fmt.Println(rejectionComment)
					mem.Status = isApproved
					mem.RejectionComment = rejectionComment
					fmt.Println(mem, isApproved)
				}
				if !mem.Status && len(mem.RejectionComment) <= 0 {
					// fmt.Println(mem.Status, r)
					// if cnt
					// p.CurrStep = step.RejectStep
					// p.Steps[step.RejectStep-1].RejectionComments = append(p.Steps[step.RejectStep-1].RejectionComments, rejectionComment)
					hasAllSubmitted = false
				} else {
					hasAllSubmitted = true
				}
				if mem.Status {
					cnt++
				}
				step.Approv[idx] = mem
			}

			// fmt.Println((*appr)[0])

			if hasAllSubmitted {
				flag := true
				for _, mem := range *appr {
					if !mem.Status {
						flag = false
						break
					}
				}
				if flag { //step.AcceptCnt {
					p.CurrStep = step.AcceptNext
				} else {
					fmt.Println("RECJ", len(step.PrefilledMembers), cnt)
					step.Approv = []Approval{}
					p.CurrStep = step.RejectStep
					p.Steps[step.RejectStep-1].RejectionComments = append(p.Steps[step.RejectStep-1].RejectionComments, rejectionComment)
				}
			} else {
				fmt.Println("ALL NOTfghnm")
			}
		} else {
			fmt.Println("GO TO HELL")
		}
	case 3:
		fmt.Println(step.Who == userID || (userID == step.Who || (step.WhoID != -1 && funk.ContainsString(p.State[step.WhoID-1], userID))))
		if c.QueryParam("accept") == "yes" && (step.Who == userID || (userID == step.Who || (step.WhoID != -1 && funk.ContainsString(p.State[step.WhoID-1], userID)))) {
			fmt.Printf("%+v\n", "called sdhfhs")
			p.CurrStep = step.AcceptNextExtra
		} else if step.Who == userID || (userID == step.Who || (step.WhoID != -1 && funk.ContainsString(p.State[step.WhoID-1], userID))) {
			fmt.Printf("%+v\n", "called")
			fmt.Println("CALLED\n\n\n\n")
			isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))
			rejectionComment := c.QueryParam("comment")

			// member := ""
			cnt := 0

			hasAllSubmitted := false

			allMem := step.PrefilledMembers

			if step.WhoID != -1 {
				allMem = p.State[step.WhoID-1]
			}

			fmt.Println("ALL MEMBERS", allMem)

			for idx, mem := range allMem {
				if len(step.Approv) < idx+1 {
					fmt.Println("ADDING", mem)
					step.Approv = append(step.Approv, Approval{
						Member:           mem,
						RejectionComment: "",
						Status:           false,
					})
				}
			}

			appr := &step.Approv

			for idx, mem := range *appr {
				if mem.Member == userID {
					fmt.Println(rejectionComment)
					mem.Status = isApproved
					mem.RejectionComment = rejectionComment
					fmt.Println(mem, isApproved)
				}
				if !mem.Status && len(mem.RejectionComment) <= 0 {
					// fmt.Println(mem.Status, r)
					// if cnt
					// p.CurrStep = step.RejectStep
					// p.Steps[step.RejectStep-1].RejectionComments = append(p.Steps[step.RejectStep-1].RejectionComments, rejectionComment)
					hasAllSubmitted = false
				} else {
					hasAllSubmitted = true
				}
				if mem.Status {
					cnt++
				}
				step.Approv[idx] = mem
			}

			// fmt.Println((*appr)[0])
			if hasAllSubmitted {
				flag := true
				for _, mem := range *appr {
					fmt.Println("CHECKING", mem)
					if !mem.Status {
						flag = false
						break
					}
				}
				if flag { //step.AcceptCnt {
					p.CurrStep = step.AcceptNext
				} else {
					fmt.Println("RECJ", len(step.PrefilledMembers), cnt)
					step.Approv = []Approval{}
					p.CurrStep = step.RejectStep
					if step.RejectStep == -1 {
						break
					}
					p.Steps[step.RejectStep-1].RejectionComments = append(p.Steps[step.RejectStep-1].RejectionComments, rejectionComment)
				}
			} else {
				fmt.Println("ALL NOTfghnm")
			}
		}
	case 4:
		if c.QueryParam("mem") == "yes" {
			if c.QueryParam("action") == "add" {
				fmt.Println("ADDING")
				members := []Approval{}
				c.Bind(&members)
				mems := []string{}
				for _, mem := range members {
					mems = append(mems, mem.Member)
				}
				if p.State == nil {
					p.State = map[int][]string{}
				}
				fmt.Println(mems)
				p.State[step.ID-1] = mems
				p.Steps[step.ID-1].Approv = members
			} else {
				mems := p.State[step.ID-1]
				i := funk.IndexOfString(mems, c.QueryParam("id"))
				if i != -1 {
					mems = append(mems[:i], mems[i+1:]...)
					if p.State == nil {
						p.State = map[int][]string{}
					}
					p.State[step.ID-1] = mems
					for idx, mem := range step.Approv {
						if mem.Member == c.QueryParam("id") {
							i = idx
							break
						}
					}
					step.Approv = append(step.Approv[:i], step.Approv[i+1:]...)
				}
			}
		} else if c.QueryParam("accept") == "yes" && (userID == step.Who || (step.WhoID != -1 && funk.ContainsString(step.PrefilledMembers, userID))) {
			p.CurrStep = step.AcceptNext
		} else {
			isApproved, _ := strconv.ParseBool(c.QueryParam("isApproved"))
			rejectionComment := c.QueryParam("comment")

			member := userID

			for idx, mem := range step.Approv {
				if mem.Member == member {
					mem.Status = isApproved
					mem.RejectionComment = rejectionComment
					step.Approv[idx] = mem
					break
				}
			}
		}
	case 5:
		fmt.Println(p.CurrStep, step.PrefilledMembers, step.WhoID, p.State[step.WhoID-1])
		if userID == step.Who || (step.WhoID == -1 && funk.ContainsString(step.PrefilledMembers, userID)) || (step.WhoID != -1 && funk.ContainsString(p.State[step.WhoID-1], userID)) {
			members := []string{}
			c.Bind(&members)
			fmt.Println("sadas", members)
			if p.State == nil {
				p.State = map[int][]string{}
			}
			p.State[step.ID-1] = members
			p.CurrStep = step.AcceptNext
		} else {
			fmt.Println("GET THE HELL")
		}
	case 6:
		p.CurrStep = step.AcceptNext
	}

	err = db.DB("").C("testing").Update(bson.M{"_id": bson.ObjectIdHex(id)}, p)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.NoContent(200)
}

func (a *app) getProcess(c echo.Context) error {
	db, err := a.s.GetMongoSession()
	if err != nil {
		return err
	}
	defer db.Close()

	id := c.Param("id")
	p := &Process{}
	err = db.DB("").C(collectionName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(p)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, p)
}
