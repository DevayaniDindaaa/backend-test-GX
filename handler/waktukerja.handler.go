package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllWaktuKerjaHandler(c *fiber.Ctx) error {
	repo := &repository.WaktuKerjaRepository{}
	service := service.NewWaktuKerjaService(repo)

	data, err := service.GetAllWaktuKerja()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "waktu kerja shown all",
		"data":    data,
	})
}
