package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

var tenagaService = &service.TenagaKerjaService{
	TenagaRepo: &repository.TenagaKerjaRepository{},
}

func GetAllTenagaKerjaHandler(c *fiber.Ctx) error {
	data, err := tenagaService.GetAllTenagaKerja()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "tenaga kerja shown all",
		"data":    data,
	})
}

func CreateTenagaKerjaHandler(c *fiber.Ctx) error {
	var tenaga models.TenagaKerja
	if err := c.BodyParser(&tenaga); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	err := tenagaService.CreateTenagaKerja(&tenaga)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success created tenaga",
		"data":    tenaga,
	})
}
