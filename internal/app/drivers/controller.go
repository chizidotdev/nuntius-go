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

	// Public routes
	router.GET("/message/:username", controller.message)
	router.POST("/message/:username", controller.sendMessage)

	// Auth routes
	router.GET("/login", controller.login)
	router.GET("/login/google", controller.loginWithSSO)
	router.GET("/callback", controller.ssoCallback)

	// Protected routes
	router.Use(controller.isAuthenticated)
	router.GET("/", controller.index)
	router.GET("/logout", controller.logout)
	router.GET("/settings", controller.settings)
	router.POST("/settings", controller.saveSettings)
	router.GET("/messages", controller.messages)

	return controller
}

func (c *Controller) Start(address string) error {
	return c.router.Run(address)
}
