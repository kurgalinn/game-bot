package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/kurgalinn/game-bot/internal/repository/sqlite"
	"github.com/kurgalinn/game-bot/internal/service"
	"github.com/kurgalinn/game-bot/internal/usecase"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	repo, err := sqlite.NewRepository(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal(err)
	}

	lt, err := layout.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	bot, err := tele.NewBot(lt.Settings())
	if err != nil {
		log.Fatal(err)
	}

	h := usecase.NewHandler(repo, lt, service.NewSetsPool(rc, time.Minute*10))

	if cmd := lt.Commands(); cmd != nil {
		if err := bot.SetCommands(cmd); err != nil {
			log.Fatal(err)
		}
	}

	bot.OnError = h.OnError
	bot.Use(middleware.AutoRespond())
	bot.Use(lt.Middleware(os.Getenv("DEFAULT_LOCALE")))

	bot.Handle("/start", h.MainMenu)
	bot.Handle(lt.Callback("back"), h.MainMenu)
	bot.Handle(lt.Callback("profile"), h.Profile)
	bot.Handle(lt.Callback("edit_nickname"), h.EditName)
	bot.Handle(lt.Callback("help"), h.Help)

	bot.Handle(lt.Callback("movie_game"), h.MovieGameStart)
	bot.Handle(tele.OnPollAnswer, h.MovieNextLevel)
	bot.Handle(lt.Callback("scoreboard"), h.MovieGameScoreboard)

	bot.Handle(tele.OnText, h.UserInput)

	bot.Start()
}
