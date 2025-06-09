package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type WaktuKerjaService struct {
	Repo *repository.WaktuKerjaRepository
}

func NewWaktuKerjaService(repo *repository.WaktuKerjaRepository) *WaktuKerjaService {
	return &WaktuKerjaService{Repo: repo}
}

func (s *WaktuKerjaService) GetAllWaktuKerja() ([]models.WaktuPembuatan, error) {
	return s.Repo.GetAllWaktuKerja()
}
