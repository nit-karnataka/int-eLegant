package main

import (
	"context"
	authProto "flok-server/auth-srv/authproto"
	chatProto "flok-server/chat-srv/chatproto"
	fileProto "flok-server/file-srv/fileproto"
	formProto "flok-server/form-srv/formproto"
	"flok-server/lib"
	meetingProto "flok-server/meeting-srv/meetingproto"
	portalProto "flok-server/portal-srv/portalproto"
	projectProto "flok-server/project-srv/projectproto"
	userProto "flok-server/user-srv/userproto"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"

	"google.golang.org/grpc"

	_ "github.com/crgimenes/goconfig/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type app struct {
	c       config
	e       *echo.Echo
	address string
	s       *lib.Store

	authClient authProto.AuthServiceClient
	authCon    *grpc.ClientConn
	authAddr   string

	chatClient chatProto.ChatServiceClient
	chatCon    *grpc.ClientConn
	chatAddr   string

	userClient userProto.UserServiceClient
	userCon    *grpc.ClientConn
	userAddr   string

	portalClient portalProto.PortalServiceClient
	portalCon    *grpc.ClientConn
	portalAddr   string

	fileClient fileProto.FileServiceClient
	fileCon    *grpc.ClientConn
	fileAddr   string

	meetingClient meetingProto.MeetingServiceClient
	meetingCon    *grpc.ClientConn
	meetingAddr   string

	projectClient projectProto.ProjectServiceClient
	projectCon    *grpc.ClientConn
	projectAddr   string

	formClient formProto.FormServiceClient
	formCon    *grpc.ClientConn
	formAddr   string
}

type config struct {
	JwtSecret string `json:"jwtSecret"`
}

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

func newApp(address, configFile, authAddr, userAddr, projectAddr, chatAddr, portalAddr, meetingAddr, fileAddr, formAddr string) *app {
	/* if _, err := os.Stat(configfile); os.IsNotExist(err) {
	}

	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		return nil
	}

	log.Printf("Read %s", string(data)) */
	cfg := getConfig(configFile)

	log.Printf("Config %+v", cfg)

	return &app{
		c:           *cfg,
		e:           echo.New(),
		address:     address,
		authAddr:    authAddr,
		userAddr:    userAddr,
		projectAddr: projectAddr,
		fileAddr:    fileAddr,
		chatAddr:    chatAddr,
		portalAddr:  portalAddr,
		meetingAddr: meetingAddr,
		formAddr:    formAddr,
		s:           &lib.Store{},
	}
}

func (a *app) initApp() error {

	err := a.s.Connect("localhost:27017", "", "", "test")
	if err != nil {
		return err
	}
	log.Println("Connecting to auth service...", a.authAddr)

	authCon, err := grpc.Dial(a.authAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	log.Println("Error while connecting to auth service", err)

	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to auth service %v", err)
		a.closeApp()
		return err
	}
	a.authCon = authCon
	authSrvClient := authProto.NewAuthServiceClient(authCon)
	a.authClient = authSrvClient
	userCon, err := grpc.Dial(a.userAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		log.Printf("error while connecting to user service %v", err)
		a.closeApp()
		// lis.Close()
		return err
	}
	a.userCon = userCon
	userSrvClient := userProto.NewUserServiceClient(userCon)
	a.userClient = userSrvClient

	projectCon, err := grpc.Dial(a.projectAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to project service %v", err)
		a.closeApp()
		return err
	}
	a.projectCon = projectCon
	projectSrvClient := projectProto.NewProjectServiceClient(projectCon)
	a.projectClient = projectSrvClient

	portalCon, err := grpc.Dial(a.portalAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to portal service %v", err)
		a.closeApp()
		return err
	}
	a.portalCon = portalCon
	portalSrvClient := portalProto.NewPortalServiceClient(portalCon)
	a.portalClient = portalSrvClient

	chatCon, err := grpc.Dial(a.chatAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to chat service %v", err)
		a.closeApp()
		return err
	}
	a.chatCon = chatCon
	chatSrvClient := chatProto.NewChatServiceClient(chatCon)
	a.chatClient = chatSrvClient

	log.Printf("meeting addr: %s", a.meetingAddr)

	meetingCon, err := grpc.Dial(a.meetingAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to meeting service %v", err)
		a.closeApp()
		return err
	}
	a.meetingCon = meetingCon
	meetingSrvClient := meetingProto.NewMeetingServiceClient(meetingCon)
	a.meetingClient = meetingSrvClient

	fileCon, err := grpc.Dial(a.fileAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to file service %v", err)
		a.closeApp()
		return err
	}
	a.fileCon = fileCon
	fileSrvClient := fileProto.NewFileServiceClient(fileCon)
	a.fileClient = fileSrvClient

	formCon, err := grpc.Dial(a.formAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*3))
	if err != nil {
		// a.eb.Close()
		// lis.Close()
		log.Printf("error while connecting to file service %v", err)
		a.closeApp()
		return err
	}
	a.formCon = formCon
	formSrvClient := formProto.NewFormServiceClient(formCon)
	a.formClient = formSrvClient

	a.e.HideBanner = true
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())
	a.e.Use(middleware.CORS())

	// a.e.Static("/", "public")

	a.e.POST("/test-file", a.testFileUpload)

	a.e.POST("/login", a.login)
	a.e.POST("/register/student", a.registerStudent)
	a.e.POST("/register/pi", a.registerPI)
	a.e.POST("/register/ec", a.registerEC)

	api := a.e.Group("/api")

	api.Use(middleware.JWT([]byte(a.c.JwtSecret)))

	// register := api.Group("/register")

	project := api.Group("/project")
	user := api.Group("/user")
	form := api.Group("/form")

	project.GET("", a.getAllProjects)
	project.GET("/", a.getAllProjects)
	project.GET("/student", a.getAllStudentProjects)
	project.GET("/student/", a.getAllStudentProjects)
	project.POST("", a.createProject)
	project.POST("/", a.createProject)
	project.GET("/:id", a.getProject)
	project.GET("/:id/thesis/:thesis", a.getThesis)
	project.PATCH("/:id/thesis/:thesis/submitted", a.setThesisSubmitted)
	project.PATCH("/:id/thesis/:thesis/approvepi", a.approveThesisPI)
	project.PATCH("/:id/thesis/:thesis/approvesec", a.approveThesisSec)
	project.PATCH("/:id/thesis/:thesis/approvedor", a.approveThesisDOR)

	project.GET("/:id/irbmeeting/:irbmeeting", a.getIRBMeeting)
	project.POST("/:id/irbmeeting/:irbmeeting/initaialmembers", a.addIRBMeetingInitailMembers)
	project.DELETE("/:id/irbmeeting/:irbmeeting/initaialmembers", a.removeIRBMeetingInitailMembers)
	project.POST("/:id/irbmeeting/:irbmeeting/agenda", a.setIRBMeetingAgenda)
	project.POST("/:id/irbmeeting/:irbmeeting/minutes", a.setIRBMeetingMinutes)
	project.POST("/:id/irbmeeting/:irbmeeting/initaialmembers", a.addIRBMeetingPresentMembers)
	project.DELETE("/:id/irbmeeting/:irbmeeting/initaialmembers", a.removeIRBMeetingPresentMembers)
	project.POST("/:id/irbmeeting/:irbmeeting/qlmembers", a.addIRBMeetingQLMembers)
	project.DELETE("/:id/irbmeeting/:irbmeeting/qlmembers", a.removeIRBMeetingQLMembers)
	project.PATCH("/:id/irbmeeting/:irbmeeting/agenda/approve", a.approveIRBMeetingAgenda)
	project.PATCH("/:id/irbmeeting/:irbmeeting/minutes/approve", a.approveIRBMeetingMinute)
	project.POST("/:id/irbmeeting/:irbmeeting/qlmembers", a.addIRBMeetingFAMembers)
	project.DELETE("/:id/irbmeeting/:irbmeeting/qlmembers", a.removeIRBMeetingFAMembers)
	project.PATCH("/:id/irbmeeting/:irbmeeting/chair", a.setIRBMeetingChair)
	project.POST("/:id/irbmeeting/:irbmeeting/qlr", a.addIRBMeetingQLR)
	project.DELETE("/:id/irbmeeting/:irbmeeting/qlr", a.removeIRBMeetingQLR)
	project.PATCH("/:id/irbmeeting/:irbmeeting/qlrpi/approve", a.aprroveIRBMeetingQLRPI)
	project.PATCH("/:id/irbmeeting/:irbmeeting/qlrsec/approve", a.aprroveIRBMeetingQLRSec)
	project.PATCH("/:id/irbmeeting/:irbmeeting/translation", a.setIRBMeetingTranslation)
	project.POST("/:id/irbmeeting/:irbmeeting/translation", a.addIRBMeetingTranslation)
	project.DELETE("/:id/irbmeeting/:irbmeeting/translation", a.removeIRBMeetingTranslation)
	project.PATCH("/:id/irbmeeting/:irbmeeting/translationpi/approve", a.aprroveIRBMeetingTranslationPI)
	project.PATCH("/:id/irbmeeting/:irbmeeting/translationsec/approve", a.aprroveIRBMeetingTranslationSec)
	project.PATCH("/:id/irbmeeting/:irbmeeting/qlreview", a.setIRBMeetingQLReview)
	project.PATCH("/:id/irbmeeting/:irbmeeting/translationreview", a.setIRBMeetingTranslationReview)
	project.PATCH("/:id/irbmeeting/:irbmeeting/fa", a.setIRBMeetingFA)

	project.POST("/process/:id", a.process)
	project.GET("/process/:id", a.process)

	user.GET("/ec", a.getAllECMembers)
	user.GET("/ec/", a.getAllECMembers)

	chat := api.Group("/chat")

	chat.GET("/:id", a.getChat)
	chat.POST("/:id", a.addChat)

	form.GET("/:id", a.getForm)
	form.POST("/:id", a.updateForm)

	// Start server
	go func() {
		if err := a.e.Start(a.address); err != nil {
			a.e.Logger.Info("shutting down the server")
		}
	}()

	log.Println("init complete")

	return nil
}

func (a *app) closeApp() {
	a.s.Close()
	log.Println("Closing app")
	if a.e != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		log.Println("Closing echo server")
		if err := a.e.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}
	log.Println("Closed echo server")

	if a.authCon != nil {
		a.authCon.Close()
	}
	log.Println("auth grpc conn closed")

	if a.userCon != nil {
		a.userCon.Close()
	}
	log.Println("user grpc conn closed")

	if a.projectCon != nil {
		a.projectCon.Close()
	}
	log.Println("hub-con grpc conn closed")
}
