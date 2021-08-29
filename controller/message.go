package controller

import (
	"os"

	"github.com/mahaupt/bg-back/models"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

type MessageController struct{}

func (m *MessageController) SendMessage(c *gin.Context) {
	var msg models.Message
	err := c.Bind(&msg)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	text := "Kontaktformular\nName: " + msg.Name + "\n" + "Email: " + msg.Email + "\n" + "Message: " + msg.Message
	whmsg := slack.WebhookMessage{
		Username: "Bitgladiator Kontakt",
		Text:     text,
	}
	err = slack.PostWebhook(os.Getenv("SLACK_WEBHOOK_URL"), &whmsg)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"name":    msg.Name,
		"email":   msg.Email,
		"message": msg.Message,
	})
}

func (m *MessageController) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"health": "ok",
	})
}
