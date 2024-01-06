package drivers

import (
	"github.com/chizidotdev/nuntius/internal/core/service"
	"github.com/gin-gonic/gin"
	"log"
)

type SendMessageReq struct {
	Content string `form:"content" binding:"required"`
}

func (c *Controller) sendMessage(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(400, gin.H{
			"message": "username is required",
		})
		return
	}

	var req SendMessageReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println("username", username)
	err = c.messageService.CreateMessage(ctx, service.CreateMessageReq{
		Content:  req.Content,
		Username: username,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	responseHTML := `
		<html>
		<p>Your message has been sent. Now it's your turn to dare your friends to send you a message!</p>
		<a href="login">Click here to Login</a>
		</html>
	`
	ctx.JSON(200, responseHTML)

	//c.message(ctx)
}
