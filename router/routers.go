package router

import (
	"github.com/DevayaniDindaaa/backend-test-GX/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public Routes
	public := app.Group("/api/v1")
	public.Get("/bahanbaku", handler.GetAllBahanHandler)

	public.Get("/produk", handler.GetAllProdukHandler)

	public.Get("/resep", handler.GetAllResepHandler)

	public.Get("/transaksi", handler.GetAllTransaksiHandler)
}
