package drivers

import (
	"github.com/a-h/templ"
	"github.com/chizidotdev/nuntius/config"
	"github.com/chizidotdev/nuntius/internal/app/components"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func renderComponent(component templ.Component, ctx *gin.Context) {
	err := component.Render(ctx, ctx.Writer)
	if err != nil {
		log.Fatal(err)
		return
	}
}

const whatsAppBaseUrl = "https://api.whatsapp.com/send?text=%F0%9F%92%80Hey%21+Write+a+%2Asecret+anonymous+message%2A+for+me..+%F0%9F%98%89+I+%2Awon%27t+know%2A+who+wrote+it..+%F0%9F%92%80%F0%9F%A4%8C+%F0%9F%91%89+"

func (c *Controller) index(ctx *gin.Context) {
	user := c.getAuthenticatedUser(ctx)
	if user.Username == "" {
		ctx.Redirect(http.StatusFound, "/settings")
		return
	}

	baseUrl := config.EnvVars.BaseUrl
	profileLink := baseUrl + "/message/" + user.Username
	whatsAppUrl := templ.SafeURL(whatsAppBaseUrl + profileLink)

	component := components.Index(user, profileLink, whatsAppUrl)
	renderComponent(component, ctx)
}
