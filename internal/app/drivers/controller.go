package drivers

import (
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	router         *gin.Engine
	userService    *service.UserService
	messageService *service.MessageService
}

const (
	profileKey = "profile"
	stateKey   = "state"
)

func NewController(
	userService *service.UserService,
	messageService *service.MessageService,
) *Controller {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("nuntius_auth", store))

	controller := &Controller{
		router:         router,
		userService:    userService,
		messageService: messageService,
	}

	router.GET("/login", controller.login)
	router.GET("/login/google", controller.loginWithSSO)
	router.GET("/callback", controller.ssoCallback)

	router.Use(controller.isAuthenticated)
	router.GET("/", controller.index)
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
