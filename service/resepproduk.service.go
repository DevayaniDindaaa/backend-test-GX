package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type ResepService struct {
	ResepRepo *repository.ResepRepository
}

func NewResepService(repo *repository.ResepRepository) *ResepService {
	return &ResepService{ResepRepo: repo}
}

func (s *ResepService) GetAllResep() ([]models.ResepProduk, error) {
	return s.ResepRepo.GetAllResep()
}
