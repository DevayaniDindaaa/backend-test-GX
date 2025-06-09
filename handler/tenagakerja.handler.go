package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllTenagaKerjaHandler(c *fiber.Ctx) error {
	repo := &repository.TenagaKerjaRepository{}
	service := service.NewTenagaKerjaService(repo)

	data, err := service.GetAllTenagaKerja()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "tenaga kerja shown all",
		"data":    data,
	})
}
