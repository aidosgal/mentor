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
	"github.com/aidosgal/mentor/internal/config"
	
	tele "gopkg.in/telebot.v4"
	_ "github.com/lib/pq"
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

	pref := tele.Settings{
		Token:  cfg.Telegram.Api,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle("/start", authHandler.HandleStart)

	b.Start()
}
