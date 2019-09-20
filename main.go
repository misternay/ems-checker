package main

import (
	"ems-checker/handlers"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	defer recoverError()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.EchoHandler(e, connectToLine())

	e.Logger.Fatal(e.Start(":6000"))
}

func connectToLine() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Printf("%+v\n", err)
	}

	return bot
}

func recoverError() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
