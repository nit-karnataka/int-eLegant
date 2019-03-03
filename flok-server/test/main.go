package main

import (
	"context"
	"flok-server/lib"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/thoas/go-funk"

	"github.com/globalsign/mgo/bson"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var collectionName = "testing"

type Approval struct {
	Member           string `json:"member,omitempty"`
	Status           bool   `json:"status,omitempty"`
	RejectionComment string `json:"rejectionComment,omitempty"`
}

type Step struct {
	ID                int        `json:"id,omitempty"`
	Name              string     `json:"name,omitempty"`
	TypeOf            int        `json:"typeOf,omitempty"`
	Status            bool       `json:"status,omitempty"`
	Doc               string     `json:"doc,omitempty"`
	Forms             []string   `json:"forms,omitempty"`
	RejectStep        int        `json:"rejectStep,omitempty"`
	AcceptNext        int        `json:"acceptNext,omitempty"`
	AcceptNextExtra   int        `json:"acceptNextExtra,omitempty"`
	RejectionComment  string     `json:"rejectionComment,omitempty"`
	RejectionComments []string   `json:"rejectionComments,omitempty"`
	Members           int        `json:"members,omitempty"`
	What              int        `json:"what,omitempty"`
	Approv            []Approval `json:"approv,omitempty"`
	OwnerTypes        []string   `json:"ownerTypes,omitempty"`
	Who               string     `json:"who,omitempty"`
	WhoID             int        `json:"whoID,omitempty"`
	PrefilledMembers  []string   `json:"prefilledMembers,omitempty"`
	GeneratorQuery    string     `json:"generatorQuery,omitempty"`
	AcceptCnt         int        `json:"acceptCnt,omitempty"`
}

type Process struct {
	ID       bson.ObjectId    `json:"id,omitempty" bson:"_id"`
	State    map[int][]string `json:"state,omitempty"`
	Steps    []Step           `json:"steps,omitempty"`
	CurrStep int              `json:"currStep,omitempty"`
}

type app struct {
	s *lib.Store
	e *echo.Echo
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
		if step.Who == userID || userID == step.Who || (step.WhoID == -1 && funk.ContainsString(step.PrefilledMembers, userID)) || funk.ContainsString(p.State[step.WhoID-1], userID) || funk.ContainsString(step.PrefilledMembers, userID) {
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
				err = c.Bind(&members)
				if err != nil {
					log.Printf("ERR: %+v", err)
				}
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
				i := funk.IndexOfString(mems, c.QueryParam("member"))
				if i != -1 {
					mems = append(mems[:i], mems[i+1:]...)
					if p.State == nil {
						p.State = map[int][]string{}
					}
					p.State[step.ID-1] = mems
					for idx, mem := range step.Approv {
						if mem.Member == c.QueryParam("member") {
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
		fmt.Println(p.CurrStep, step.PrefilledMembers, step.WhoID, p.State[step.WhoID-1], userID)
		if userID == step.Who || (step.WhoID == -1 && funk.ContainsString(step.PrefilledMembers, userID)) || (step.WhoID != -1 && funk.ContainsString(p.State[step.WhoID-1], userID)) || funk.ContainsString(step.PrefilledMembers, userID) {
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

func (a *app) closeApp() {
	log.Println("Closing app")
	if a.e != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		log.Println("Closing echo server")
		if err := a.e.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}
}

func (a *app) get(c echo.Context) error {
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

	return c.JSON(http.StatusOK, p)
}

func (a *app) initApp() error {
	err := a.s.Connect("localhost:27017", "", "", "test")
	if err != nil {
		return err
	}
	a.e.HideBanner = true
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())
	a.e.Use(middleware.CORS())

	a.e.POST("/:id", a.process)
	a.e.GET("/:id", a.get)

	go func() {
		if err := a.e.Start(":8081"); err != nil {
			a.e.Logger.Info("shutting down the server")
		}
	}()

	return nil
}

func main() {
	/* var typeOf = map[int]string{
		1: "UPLOAD",
		2: "APPROVE",
		3: "GRP APPROVE",
		4: "SELECT",
	} */

	/* p := &Process{
		ID: bson.NewObjectId().Hex(),
		Steps: []Step{
			Step{
				ID:         1,
				AcceptNext: 2,
				TypeOf:     1,
			},
			Step{
				ID:         2,
				TypeOf:     2,
				AcceptNext: -1,
				What:       1,
				Who:        "jkjjkjjhjk",
				RejectStep: 1,
			},
		},
	}

	s := lib.Store{}
	err := s.Connect("localhost:27017", "", "", "test")
	if err != nil {
		panic(err)
	}

	defer s.Close()

	ss, err := s.GetMongoSession()
	if err != nil {
		panic(err)
	}

	err = ss.DB("").C("testing").Insert(p)
	if err != nil {
		panic(err)
	} */

	a := &app{
		e: echo.New(),
		s: &lib.Store{},
	}

	if err := a.initApp(); err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Exiting")

	a.closeApp()
}
