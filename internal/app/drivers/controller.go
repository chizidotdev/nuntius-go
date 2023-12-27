package drivers

import (
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-gonic/gin"
)

type (
	Controller struct {
		router         *gin.Engine
		userService    *service.UserService
		messageService *service.MessageService
	}

	Link struct {
		Uri   string
		Title string
	}
)

var (
	AuthenticatedLinks = []Link{
		{
			Uri:   "/",
			Title: "Home",
		},
		{
			Uri:   "/messages",
			Title: "Messages",
		},
		{
			Uri:   "/settings",
			Title: "Settings",
		},
	}

	UnAuthenticatedLinks = []Link{
		{
			Uri:   "/login",
			Title: "Login",
		},
	}
)

func NewController(
	userService *service.UserService,
	messageService *service.MessageService,
) *Controller {
	router := gin.Default()

	controller := &Controller{
		router:         router,
		userService:    userService,
		messageService: messageService,
	}

	router.GET("/", controller.Index)
	//router.GET("/login", controller.Login)
	//router.POST("/login", controller.LoginPost)
	//router.GET("/register", controller.Register)
	//router.POST("/register", controller.RegisterPost)
	//router.GET("/logout", controller.Logout)
	//
	//router.GET("/settings", controller.Settings)
	//router.POST("/settings", controller.SettingsPost)
	//
	//router.GET("/messages", controller.Messages)
	//router.GET("/messages/:id", controller.Message)
	//router.POST("/messages/:id", controller.MessagePost)

	return controller
}

func (c *Controller) Start(address string) error {
	return c.router.Run(address)
}
