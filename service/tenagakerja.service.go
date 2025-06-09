package service

import (
	"time"

	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type TenagaKerjaService struct {
	TenagaRepo *repository.TenagaKerjaRepository
}

func NewTenagaKerjaService(repo *repository.TenagaKerjaRepository) *TenagaKerjaService {
	return &TenagaKerjaService{TenagaRepo: repo}
}

func (s *TenagaKerjaService) GetAllTenagaKerja() ([]models.TenagaKerja, error) {
	return s.TenagaRepo.GetAllTenagaKerja()
}

func (s *TenagaKerjaService) CreateTenagaKerja(tenagaKerja *models.TenagaKerja) error {
	tenagaKerja.CreatedAt = time.Now()
	tenagaKerja.UpdatedAt = time.Now()
	return s.TenagaRepo.CreateTenagaKerja(tenagaKerja)
}
