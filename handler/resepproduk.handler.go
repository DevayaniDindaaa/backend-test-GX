package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

var resepService = &service.ResepService{
	ResepRepo: &repository.ResepRepository{},
}

func GetAllResepHandler(c *fiber.Ctx) error {
	data, err := resepService.GetAllResep()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "resep produk shown all",
		"data":    data,
	})
}
