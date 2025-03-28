package handler

import (
	"context"
	"log/slog"

	"github.com/aidosgal/mentor/internal/category/service"
	tele "gopkg.in/telebot.v4"
)

type Handler interface {
	HandleList(c tele.Context) error
}

type handler struct {
	log *slog.Logger
	service service.Service
}

func NewHandler(log *slog.Logger, service service.Service) Handler {
	return &handler{
		log: log,
		service: service,
	}
}

func (h *handler) HandleList(c tele.Context) error {
    ctx := context.Background()
    categories, err := h.service.List(ctx)
    if err != nil {
        h.log.Error("Failed to retrieve categories", "error", err)
        return err
    }

    menu := &tele.ReplyMarkup{ResizeKeyboard: true}
    
    var buttons []tele.Btn

    for _, category := range categories {
        btn := menu.Text(category.Name)
        buttons = append(buttons, btn)
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
