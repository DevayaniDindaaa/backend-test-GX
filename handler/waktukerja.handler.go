package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

var waktuService = &service.WaktuKerjaService{
	WaktuRepo: &repository.WaktuKerjaRepository{},
}

func GetAllWaktuKerjaHandler(c *fiber.Ctx) error {
	data, err := waktuService.GetAllWaktuKerja()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "waktu kerja shown all",
		"data":    data,
	})
}
