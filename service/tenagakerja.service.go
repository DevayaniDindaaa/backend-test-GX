package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type TenagaKerjaService struct {
	Repo *repository.TenagaKerjaRepository
}

func NewTenagaKerjaService(repo *repository.TenagaKerjaRepository) *TenagaKerjaService {
	return &TenagaKerjaService{Repo: repo}
}

func (s *TenagaKerjaService) GetAllTenagaKerja() ([]models.TenagaKerja, error) {
	return s.Repo.GetAllTenagaKerja()
}
