package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllProdukHandler(c *fiber.Ctx) error {
	repo := &repository.ProdukRepository{}
	service := service.NewProdukService(repo)

	data, err := service.GetAllProduk()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "produk shown all",
		"data":    data,
	})
}
