package drivers

import (
	"github.com/chizidotdev/nuntius/internal/core/domain"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *Controller) isAuthenticated(ctx *gin.Context) {
	user := c.getAuthenticatedUser(ctx)
	if user.Email == "" {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	ctx.Next()
}

func (c *Controller) getAuthenticatedUser(ctx *gin.Context) *domain.User {
	session := sessions.Default(ctx)
	profile := session.Get(profileKey)
	user, ok := profile.(domain.User)
	if !ok {
		return &domain.User{}
	}

	log.Println("user >>>>>>>>>", user)
	log.Println("username .........", user.Username)
	return &user
}
