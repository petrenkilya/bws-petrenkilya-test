package handlers

import (
	"bws_test/app/usecase"
	"github.com/gofiber/fiber"
	"log"
)

const internalErrorStatus = 500

type HackerHandler struct {
	useCase usecase.Hackers
	logger  *log.Logger
}

func CreateHackersHandler(useCase usecase.Hackers, logger *log.Logger) *HackerHandler {
	return &HackerHandler{useCase: useCase, logger: logger}
}

func (h *HackerHandler) Get(c *fiber.Ctx) {
	result, err := h.useCase.Get()
	if err != nil {
		h.logger.Print(err)
		c.SendStatus(internalErrorStatus)
		return
	}
	err = c.JSON(result)
	if err != nil {
		h.logger.Print(err)
		c.SendStatus(internalErrorStatus)
		return
	}
}
