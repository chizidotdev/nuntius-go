package drivers

import (
	"fmt"
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *Controller) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, "/login")
}

func (c *Controller) loginWithSSO(ctx *gin.Context) {
	errRedirectURL := "" //config.EnvVars.AuthDomain + "/u/login/errors"
	state, err := c.userService.GenerateAuthState()
	if err != nil {
		ctx.Redirect(http.StatusFound, errRedirectURL)
		return
	}
	session := sessions.Default(ctx)
	session.Set(stateKey, state)
	if err := session.Save(); err != nil {
		ctx.Redirect(http.StatusFound, errRedirectURL)
		return
	}

	googleAuthConfig := c.userService.GetGoogleAuthConfig()
	url := googleAuthConfig.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *Controller) ssoCallback(ctx *gin.Context) {
	errRedirectURL := "/auth/error"

	session := sessions.Default(ctx)
	if ctx.Query(stateKey) != session.Get(stateKey) {
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s?errors=invalid_state", errRedirectURL))
		return
	}

	code := ctx.Query("code")
	userProfile, err := c.userService.GoogleCallback(ctx, code)
	if err != nil {
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s?errors=failed_to_exchange", errRedirectURL))
		return
	}

	session.Set(profileKey, userProfile)
	if err := session.Save(); err != nil {
		ctx.Redirect(http.StatusFound, errRedirectURL)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (c *Controller) saveSettings(ctx *gin.Context) {
	var req service.SaveSettingsReq
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := c.getAuthenticatedUser(ctx)
	log.Println(user)
	req.Email = user.Email
	user, err = c.userService.SaveSettings(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(ctx)
	session.Set(profileKey, user)
	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}
