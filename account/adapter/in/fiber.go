package in

import (
	"strconv"

	"github.com/Terence1105/hexagonal/account/application/port/in"
	"github.com/gofiber/fiber/v2"
)

type FiberAdapter struct {
	transferUseCase in.ITransferUseCase
}

func (f *FiberAdapter) Transfer(c *fiber.Ctx) error {
	fromAccountName := c.Params("fromAccountName")
	toAccountName := c.Params("toAccountName")
	amount := c.Params("amount")

	a, err := strconv.Atoi(amount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	transferCommand, err := in.NewTransferCommand(fromAccountName, toAccountName, a)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = f.transferUseCase.TransferUseCase(transferCommand)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (f *FiberAdapter) Run() {
	app := fiber.New()
	app.Get("/transfer/:fromAccountName/:toAccountName/:amount", f.Transfer)
	app.Listen(":8080")
}

func NewFiberAdapter(transferUseCase in.ITransferUseCase) FiberAdapter {
	return FiberAdapter{
		transferUseCase: transferUseCase,
	}
}
