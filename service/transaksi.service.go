package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type TransaksiService struct {
	TransaksiRepo *repository.TransaksiRepository
}

func NewTransaksiService(repo *repository.TransaksiRepository) *TransaksiService {
	return &TransaksiService{TransaksiRepo: repo}
}

func (s *TransaksiService) GetAllTransaksi() ([]models.TransaksiPenjualan, error) {
	return s.TransaksiRepo.GetAllTransaksi()
}
