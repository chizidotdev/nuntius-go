package drivers

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/chizidotdev/nuntius/internal/app/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

func renderComponent(component templ.Component, ctx *gin.Context) {
	err := component.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

const whatsAppBaseUrl = "https://api.whatsapp.com/send?text=%F0%9F%92%80Hey%21+Write+a+%2Asecret+anonymous+message%2A+for+me..+%F0%9F%98%89+I+%2Awon%27t+know%2A+who+wrote+it..+%F0%9F%92%80%F0%9F%A4%8C+%F0%9F%91%89+https://nuntius.aidmedium.com/message/"
const baseUrl = "https://nuntius.aidmedium.com"

func (c *Controller) index(ctx *gin.Context) {
	user := c.getAuthenticatedUser(ctx)
	if user.Username == "" {
		ctx.Redirect(http.StatusFound, "/settings")
		return
	}

	whatsAppUrl := templ.SafeURL(whatsAppBaseUrl + user.Username)
	profileLink := baseUrl + "/message/" + user.Username
	component := components.Index(user, profileLink, whatsAppUrl)
	renderComponent(component, ctx)
}

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

func (c *Controller) messages(ctx *gin.Context) {
	user := c.getAuthenticatedUser(ctx)
	msgs, err := c.messageService.ListMessages(ctx, user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	component := components.Messages(msgs)
	renderComponent(component, ctx)
}

func (c *Controller) message(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.userService.GetByUsername(ctx, username)
	if err != nil {
		errMsg := fmt.Sprintf("User with username '%s' not found", username)
		component := components.Message("", errMsg)
		renderComponent(component, ctx)
		return
	}
	component := components.Message(user.Username, "")
	renderComponent(component, ctx)
}
