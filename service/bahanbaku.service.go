package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type BahanBakuService struct {
	Repo *repository.BahanBakuRepository
}

func NewBahanBakuService(repo *repository.BahanBakuRepository) *BahanBakuService {
	return &BahanBakuService{Repo: repo}
}

func (s *BahanBakuService) GetAllBahanBaku() ([]models.BahanBaku, error) {
	return s.Repo.GetAllBahanBaku()
}
