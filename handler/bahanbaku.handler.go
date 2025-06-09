package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllBahanHandler(c *fiber.Ctx) error {
	repo := &repository.BahanBakuRepository{}
	service := service.NewBahanBakuService(repo)

	data, err := service.GetAllBahanBaku()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "bahan baku shown all",
		"data":    data,
	})
}
