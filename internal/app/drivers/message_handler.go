package drivers

import (
	"github.com/chizidotdev/nuntius/internal/app/components"
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-gonic/gin"
)

type SendMessageReq struct {
	Content string `form:"content" binding:"required"`
}

func (c *Controller) sendMessage(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		renderComponent(components.MessageError("Username is required"), ctx)
		return
	}

	var req SendMessageReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		renderComponent(components.MessageError(err.Error()), ctx)
		return
	}

	err = c.messageService.CreateMessage(ctx, service.CreateMessageReq{
		Content:  req.Content,
		Username: username,
	})
	if err != nil {
		renderComponent(components.MessageError(err.Error()), ctx)
		return
	}

	renderComponent(components.MessageResponse(), ctx)
}
