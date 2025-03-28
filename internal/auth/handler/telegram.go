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
    BtnReturn = menu.Text("Изменить область")
	BtnReview = menu.Text("Оставить отзыв ментору")
	
    BtnAbout = inlineMenu.Data("Про нас", "about_us")
    BtnWho = inlineMenu.Data("Кому подходит?", "for_who")
    BtnMentor = inlineMenu.Data("Стать ментором", "become_mentor")
	BtnNext = inlineMenu.Data("Следующий", "next")
	BtnPrev = inlineMenu.Data("Предыдущий", "prev")
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
	firstName := c.Message().Chat.FirstName
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
        `Добро пожаловать, %s! 👋

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

	err = c.Send(heroImage, inlineMenu)
	if err != nil {
		return err
	}
	err = c.Send(".", menu)
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
	profileText := `🏆 Галимжан Айдос - Senior Software Engineer

Профессиональный опыт:

• Google (2018-2023): Senior Backend Engineer
  - Разработка масштабируемых облачных сервисов
  - Работа над ключевыми проектами Google Cloud Platform

Стоимость услуг:
• Разовая консультация (1 час): 15 000 ₸
• Пакет из 5 консультаций: 65 000 ₸
• Подготовка к интервью (2 недели): 30 000 ₸

Связь: @bizzarchikk
`

	profileImage := &tele.Photo{
		File: tele.FromURL("https://img.freepik.com/free-photo/close-up-upset-american-black-person_23-2148749582.jpg"),
		Caption: profileText,
	}

	inlineMenu.Inline(
		inlineMenu.Row(BtnPrev, BtnNext),
	)
		
	menu.Reply(
		menu.Row(BtnReturn, BtnHelp),
		menu.Row(BtnReview),
	)
	
	err := c.Send(profileImage, inlineMenu)
	if err != nil {
		return err
	}
	err = c.Send(".", menu)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) HandleHelp(c tele.Context) error {
	return c.Send("сам себе помоги")
}

func (h *handler) HandleReview(c tele.Context) error {
	return c.Send("стукач что ли?")
}

