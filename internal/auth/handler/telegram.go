package handler

import (
	"log/slog"

	"github.com/aidosgal/mentor/internal/auth/service"
	tele "gopkg.in/telebot.v4"
)

type Handler interface {
	HandleStart(c tele.Context) error
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

func (h *handler) HandleStart(c tele.Context) error {
	return c.Send("Саламалейкум")
}
