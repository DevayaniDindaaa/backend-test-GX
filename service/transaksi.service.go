package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type TransaksiService struct {
	Repo *repository.TransaksiRepository
}

func NewTransaksiService(repo *repository.TransaksiRepository) *TransaksiService {
	return &TransaksiService{Repo: repo}
}

func (s *TransaksiService) GetAllTransaksi() ([]models.TransaksiPenjualan, error) {
	return s.Repo.GetAllTransaksi()
}
