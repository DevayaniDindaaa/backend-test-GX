package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

var produkService = &service.ProdukService{
	ProdukRepo: &repository.ProdukRepository{},
}

func GetAllProdukHandler(c *fiber.Ctx) error {
	data, err := produkService.GetAllProduk()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "produk shown all",
		"data":    data,
	})
}

func CreateProduk(c *fiber.Ctx) error {
	var input models.ProdukInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	resp, err := produkService.CreateProdukWithResep(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "product created",
		"data":    resp,
	})
}
