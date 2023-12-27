package drivers

import (
	"github.com/a-h/templ"
	"github.com/chizidotdev/nuntius/internal/app/components"
	"github.com/gin-gonic/gin"
)

func renderComponent(component templ.Component, ctx *gin.Context) {
	err := component.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (c *Controller) Index(ctx *gin.Context) {
	component := components.Index()
	renderComponent(component, ctx)
}

func (c *Controller) Login(ctx *gin.Context) {
	component := components.Login()
	renderComponent(component, ctx)
}
