package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type ProdukService struct {
	Repo *repository.ProdukRepository
}

func NewProdukService(repo *repository.ProdukRepository) *ProdukService {
	return &ProdukService{Repo: repo}
}

func (s *ProdukService) GetAllProduk() ([]models.Produk, error) {
	return s.Repo.GetAllProduk()
}
