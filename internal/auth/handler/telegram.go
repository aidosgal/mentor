package handler

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/aidosgal/mentor/internal/auth/data"
	"github.com/aidosgal/mentor/internal/auth/service"
	tele "gopkg.in/telebot.v4"
)

type Handler interface {
	HandleStart(c tele.Context) error
	HandleAbout(c tele.Context) error
	HandleWho(c tele.Context) error
	HandleMentor(c tele.Context) error
	HandleListMentor(c tele.Context, category string) error
	HandleHelp(c tele.Context) error
	HandleReview(c tele.Context) error
}

type handler struct {
	log     *slog.Logger
	service service.Service
}

func NewHandler(log *slog.Logger, service service.Service) Handler {
	return &handler{
		log:     log,
		service: service,
	}
}

var (
	menu = &tele.ReplyMarkup{ResizeKeyboard: true}
	inlineMenu = &tele.ReplyMarkup{}

	BtnHelp   = menu.Text("Помощь")
	BtnStart  = menu.Text("Найти Ментора")
	BtnReview = menu.Text("Оставить отзыв ментору")
	
    BtnAbout = inlineMenu.Data("Про нас", "about_us")
    BtnWho = inlineMenu.Data("Кому подходит?", "for_who")
    BtnMentor = inlineMenu.Data("Стать ментором", "become_mentor")
)

func (h *handler) HandleStart(c tele.Context) error {
	menu.Reply(
		menu.Row(BtnStart, BtnHelp),
		menu.Row(BtnReview),
	)
	
	inlineMenu.Inline(
        inlineMenu.Row(BtnAbout, BtnWho),
        inlineMenu.Row(BtnMentor),
    )
	
	userName := c.Message().Chat.Username
	log := h.log.With(
		"method", "HandleStart",
		"userName", userName,
	)
	firstName := c.Message().Chat.FirstNamen
	lastName := c.Message().Chat.LastName
	chatID := c.Message().Chat.ID

	userModel := &data.UserModel{
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		ChatID:    strconv.Itoa(int(chatID)),
		Role:      "mentee",
	}

	ctx := context.Background()

	_, isNewUser, err := h.service.Create(ctx, userModel)
	if err != nil {
		log.Error("Failed to create user", "error", err)
		return err
	}

	var name string
	if firstName == "" || lastName == "" {
		name = fmt.Sprintf("%s%s", firstName, lastName)
	} else if firstName != "" && lastName != "" {
		name = fmt.Sprintf("%s %s", firstName, lastName)
	} else if firstName == "" && lastName == "" {
		name = userName
	}

	var welcomeText string
	if isNewUser {
		welcomeText = fmt.Sprintf("Саламалейкум, %s", name)
	} else {
		welcomeText = fmt.Sprintf("Саламалейкум, %s", name)
	}

	welcomeText = fmt.Sprintf(
        `Саламалейкум, %s! 👋

Представляем вам mitti – ваш персональный инструмент карьерного роста!

🚀 Мы – не просто бот, а стратегический партнер в мире профессиональных возможностей

Мы готовы помочь вам достичь новых профессиональных высот!
        `,
        name,
    )

	heroImage := &tele.Photo{
		File: tele.FromDisk("./public/hero.png"),
		Caption: welcomeText,
	}

	err = c.Send(heroImage, inlineMenu, menu)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) HandleAbout(c tele.Context) error {
	return c.Send("Мы крутые")
}

func (h *handler) HandleWho(c tele.Context) error {
	return c.Send("Всем")
}

func (h *handler) HandleMentor(c tele.Context) error {
	return c.Send("Ага, щас")
}

func (h *handler) HandleListMentor(c tele.Context, category string) error {
	return c.Send("ищи сам")
}

func (h *handler) HandleHelp(c tele.Context) error {
	return c.Send("сам себе помоги")
}

func (h *handler) HandleReview(c tele.Context) error {
	return c.Send("стукач что ли?")
}

