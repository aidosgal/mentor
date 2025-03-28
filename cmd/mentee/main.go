package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	authhandler "github.com/aidosgal/mentor/internal/auth/handler"
	authrepository "github.com/aidosgal/mentor/internal/auth/repository"
	authservice "github.com/aidosgal/mentor/internal/auth/service"
	categoryrepository "github.com/aidosgal/mentor/internal/category/repository"
	categoryservice "github.com/aidosgal/mentor/internal/category/service"
	categoryhandler "github.com/aidosgal/mentor/internal/category/handler"
	"github.com/aidosgal/mentor/internal/config"

	_ "github.com/lib/pq"
	tele "gopkg.in/telebot.v4"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)
	
	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		panic(err)
	}
	
	authRepository := authrepository.NewRepository(db, log)
	authService := authservice.NewService(log, authRepository)
	authHandler := authhandler.NewHandler(log , authService)

	categoryRepository := categoryrepository.NewRepository(log, db)
	categoryService := categoryservice.NewService(log, categoryRepository)
	categoryHandler := categoryhandler.NewHandler(log, categoryService)

	pref := tele.Settings{
		Token:  cfg.Telegram.Api,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle("/start", authHandler.HandleStart)
	b.Handle(&authhandler.BtnAbout, authHandler.HandleAbout)
	b.Handle(&authhandler.BtnWho, authHandler.HandleWho)
	b.Handle(&authhandler.BtnMentor, authHandler.HandleMentor)
	b.Handle(&authhandler.BtnReview, authHandler.HandleReview)
	b.Handle(&authhandler.BtnHelp, authHandler.HandleHelp)

	b.Handle(&authhandler.BtnStart, categoryHandler.HandleList)

	b.Start()
}
