package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TOKEN")
}

func main() {
	token = os.Getenv("TOKEN")
	fmt.Println(token)
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	log.Fatal(bot.Start())
}

func (a *application) startHandler(m *tbot.Message) {
	msg := "Server is working!!!"
	a.client.SendMessage(m.Chat.ID, msg)
	variable := "world"
	r := gin.Default()
	r.GET(variable)
	r.GET("/hello/:variable", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})
	
	r.Run(":8080")

}
