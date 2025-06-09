package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type ResepService struct {
	Repo *repository.ResepRepository
}

func NewResepService(repo *repository.ResepRepository) *ResepService {
	return &ResepService{Repo: repo}
}

func (s *ResepService) GetAllResep() ([]models.ResepProduk, error) {
	return s.Repo.GetAllResep()
}
