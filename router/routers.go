package router

import (
	"github.com/DevayaniDindaaa/backend-test-GX/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public Routes
	public := app.Group("/api/v1")
	public.Get("/bahanbaku", handler.GetAllBahanHandler)

	public.Get("/tenagakerja", handler.GetAllTenagaKerjaHandler)

	public.Get("/waktupembuatan", handler.GetAllWaktuKerjaHandler)

	public.Get("/produk", handler.GetAllProdukHandler)
	public.Post("/produk", handler.CreateProduk)

	public.Get("/resep", handler.GetAllResepHandler)

	public.Get("/transaksi", handler.GetAllTransaksiHandler)
}
