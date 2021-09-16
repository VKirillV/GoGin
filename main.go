package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type Test struct {
	id   int
	name string
	time string
}

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
	r := gin.Default()
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3307)/test")
	if err != nil {
		panic(err.Error())
	}

	res, err := db.Query("SELECT * FROM test")

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
	var data Test
	for res.Next() {
		err := res.Scan(&data.id, &data.name, &data.time)
		if err != nil {
			panic(err.Error())
		}
	}
	r.GET("/table/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Time": &data.time,
			"Name": &data.name,
			"id":   &data.id,
		})

	})

	r.GET("/hello/:article", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": c.Param("article"),
		})

		var msg string = c.Param("article")
		a.client.SendMessage(m.Chat.ID, msg)

		insert, err := db.Query("INSERT INTO test(name) VALUES(?)", msg)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		fmt.Println("Succesfully")
	})
	r.Run(":8080")
}
