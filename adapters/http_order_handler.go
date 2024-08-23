package adapters

import (
	"clean_architecture/entities"
	"clean_architecture/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
	orderUserCase usecases.OrderUseCase
}

func NewHttpOrderHandler(useCase usecases.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUserCase: useCase}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Invalid"})
	}

	if err := h.orderUserCase.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "something went wrong"})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}
