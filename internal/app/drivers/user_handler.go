package drivers

import (
	"fmt"
	"github.com/chizidotdev/nuntius/internal/app/components"
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) login(ctx *gin.Context) {
	user := c.getAuthenticatedUser(ctx)
	if user.Email != "" {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	component := components.Login()
	renderComponent(component, ctx)
}

func (c *Controller) settings(ctx *gin.Context) {
	component := components.Settings()
	renderComponent(component, ctx)
}

func (c *Controller) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	if err := session.Save(); err != nil {
		renderComponent(components.LogoutError(), ctx)
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
		renderComponent(components.SettingsError("Invalid username input"), ctx)
		return
	}

	user := c.getAuthenticatedUser(ctx)
	req.Email = user.Email
	user, err = c.userService.SaveSettings(ctx, req)
	if err != nil {
		renderComponent(components.SettingsError(err.Error()), ctx)
		return
	}

	session := sessions.Default(ctx)
	session.Set(profileKey, user)
	if err := session.Save(); err != nil {
		renderComponent(components.SettingsError(err.Error()), ctx)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}
