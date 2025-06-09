package handler

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
	"github.com/DevayaniDindaaa/backend-test-GX/service"
	"github.com/gofiber/fiber/v2"
)

var bahanBakuService = &service.BahanBakuService{
	BahanRepo: &repository.BahanBakuRepository{},
}

func GetAllBahanHandler(c *fiber.Ctx) error {
	data, err := bahanBakuService.GetAllBahanBaku()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "bahan baku shown all",
		"data":    data,
	})
}

func CreateBahanBakuHandler(c *fiber.Ctx) error {
	var bahan models.BahanBaku
	if err := c.BodyParser(&bahan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	err := bahanBakuService.CreateBahanBaku(&bahan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success created bahan baku",
		"data":    bahan,
	})
}
