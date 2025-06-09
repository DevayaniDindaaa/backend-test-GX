package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

var transaksiService = &service.TransaksiService{
	TransaksiRepo: &repository.TransaksiRepository{},
}

func GetAllTransaksiHandler(c *fiber.Ctx) error {
	data, err := transaksiService.GetAllTransaksi()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaksi shown all",
		"data":    data,
	})
}
