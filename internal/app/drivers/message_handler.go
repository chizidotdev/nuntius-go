package drivers

import (
	"fmt"
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
