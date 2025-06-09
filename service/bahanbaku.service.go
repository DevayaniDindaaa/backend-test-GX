package service

import (
	"time"

	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type BahanBakuService struct {
	BahanRepo *repository.BahanBakuRepository
}

func NewBahanBakuService(repo *repository.BahanBakuRepository) *BahanBakuService {
	return &BahanBakuService{BahanRepo: repo}
}

func (s *BahanBakuService) GetAllBahanBaku() ([]models.BahanBaku, error) {
	return s.BahanRepo.GetAllBahanBaku()
}

func (s *BahanBakuService) CreateBahanBaku(bahan *models.BahanBaku) error {
	bahan.CreatedAt = time.Now()
	bahan.UpdatedAt = time.Now()
	return s.BahanRepo.CreateBahanBaku(bahan)
}
