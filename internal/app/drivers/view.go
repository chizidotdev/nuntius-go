package drivers

import (
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

func (c *Controller) index(ctx *gin.Context) {
	user := c.getAuthenticatedUser(ctx)
	if user.Username == "" {
		ctx.Redirect(http.StatusFound, "/settings")
		return
	}
	component := components.Index()
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
