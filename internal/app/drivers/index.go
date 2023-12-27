package drivers

import "github.com/gin-gonic/gin"

func (c *Controller) Index(ctx *gin.Context) {
	component := index()
	err := component.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
