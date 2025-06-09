package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllTransaksiHandler(c *fiber.Ctx) error {
	repo := &repository.TransaksiRepository{}
	service := service.NewTransaksiService(repo)

	data, err := service.GetAllTransaksi()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaksi shown all",
		"data":    data,
	})
}
