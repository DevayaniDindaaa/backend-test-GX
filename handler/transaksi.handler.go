package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
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

func CreateTransaksiHandler(c *fiber.Ctx) error {
	var transaksi models.TransaksiPenjualan
	if err := c.BodyParser(&transaksi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	err := transaksiService.CreateTransaksi(&transaksi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "transaksi inserted",
		"data":    transaksi,
	})
}
