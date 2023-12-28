package drivers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) loginWithSSO(ctx *gin.Context) {
	errRedirectURL := "" //config.EnvVars.AuthDomain + "/u/login/errors"
	state, err := c.userService.GenerateAuthState()
	if err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, errRedirectURL)
		return
	}
	session := sessions.Default(ctx)
	session.Set(stateKey, state)
	if err := session.Save(); err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, errRedirectURL)
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
		ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s?errors=invalid_state", errRedirectURL))
		return
	}

	code := ctx.Query("code")
	userProfile, err := c.userService.GoogleCallback(ctx, code)
	if err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("%s?errors=failed_to_exchange", errRedirectURL))
		return
	}

	session.Set(profileKey, userProfile)
	if err := session.Save(); err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, errRedirectURL)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, "/")
}
