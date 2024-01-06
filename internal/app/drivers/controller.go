package drivers

import (
	"github.com/chizidotdev/nuntius/config"
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
	db *gorm.DB,
	userService *service.UserService,
	messageService *service.MessageService,
) *Controller {
	router := gin.Default()

	store := gormsessions.NewStore(db, true, []byte(config.EnvVars.AuthSecret))
	store.Options(sessions.Options{
		MaxAge:   86400 * 7, // 7 days
		Secure:   false,
		HttpOnly: true,
		Domain:   config.EnvVars.CookieDomain,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
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
