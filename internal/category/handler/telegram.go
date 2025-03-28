package handler

import (
	"context"
	"log/slog"
	
	"github.com/aidosgal/mentor/internal/category/service"
	tele "gopkg.in/telebot.v4"
)

var (
	CategoryButtons = make(map[string]tele.Btn)
	CategoryNames   []string
)

type Handler interface {
	HandleList(c tele.Context) error
	InitializeCategories(ctx context.Context) error
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

func (h *handler) InitializeCategories(ctx context.Context) error {
	categories, err := h.service.List(ctx)
	if err != nil {
		h.log.Error("Failed to retrieve categories", "error", err)
		return err
	}

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	CategoryButtons = make(map[string]tele.Btn)
	CategoryNames = []string{}

	for _, category := range categories {
		btn := menu.Text(category.Name)
		CategoryButtons[category.Name] = btn
		CategoryNames = append(CategoryNames, category.Name)
	}

	return nil
}

func (h *handler) HandleList(c tele.Context) error {
	ctx := context.Background()
	if err := h.InitializeCategories(ctx); err != nil {
		return err
	}
	
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	var buttons []tele.Btn

	for _, name := range CategoryNames {
		buttons = append(buttons, CategoryButtons[name])
	}

	var rows []tele.Row
	for i := 0; i < len(buttons); i += 2 {
		if i+1 < len(buttons) {
			rows = append(rows, menu.Row(buttons[i], buttons[i+1]))
		} else {
			rows = append(rows, menu.Row(buttons[i]))
		}
	}

	menu.Reply(rows...)
	return c.Send("Выберите категорию:", menu)
}
