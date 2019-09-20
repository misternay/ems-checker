package handlers

import (
	"context"
	"ems-checker/thai-post"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type HTTPCallBackHanlder struct {
	Bot *linebot.Client
}

// NewServiceHTTPHandler provide the init set up handlers path to handle request
func EchoHandler(e *echo.Echo, linebot *linebot.Client) {

	hanlders := &HTTPCallBackHanlder{Bot: linebot}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "it's so good.")
	})

	e.POST("/callback", hanlders.Callback)
}

// Callback for handle request from line
func (handler *HTTPCallBackHanlder) Callback(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	events, err := handler.Bot.ParseRequest(c.Request())

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.String(400, linebot.ErrInvalidSignature.Error())
		} else {
			c.String(500, "internal")
		}
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				messageFromPing := thai_post.GetEmsInformation(message.Text)
				if _, err = handler.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(messageFromPing)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}

	return c.JSON(200, "")
}
