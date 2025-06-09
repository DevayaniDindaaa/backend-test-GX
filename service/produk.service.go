package service

import (
	"github.com/DevayaniDindaaa/backend-test-GX/models"
	"github.com/DevayaniDindaaa/backend-test-GX/repository"
)

type ProdukService struct {
	ProdukRepo *repository.ProdukRepository
}

func NewProdukService(repo *repository.ProdukRepository) *ProdukService {
	return &ProdukService{ProdukRepo: repo}
}

func (s *ProdukService) GetAllProduk() ([]models.Produk, error) {
	return s.ProdukRepo.GetAllProduk()
}

func (s *ProdukService) CreateProdukWithResep(input models.ProdukInput) (map[string]interface{}, error) {
	// 1. Insert Produk (harga_jual sementara = 0)
	produkID, err := s.ProdukRepo.InsertProduk(input.NamaProduk)
	if err != nil {
		return nil, err
	}

	// 2. Simpan resep
	for _, r := range input.Resep {
		err := s.ProdukRepo.InsertResep(produkID, r.IDBahan, r.Jumlah)
		if err != nil {
			return nil, err
		}
	}

	// 3. Simpan data tenaga kerja
	for _, t := range input.TenagaKerja {
		err := s.ProdukRepo.InsertWaktu(produkID, t.IDTenaga, t.JamKerja)
		if err != nil {
			return nil, err
		}
	}

	// 4. Hitung HPP bahan baku
	hppBahan, err := s.ProdukRepo.SumarizeHPP(produkID)
	if err != nil {
		return nil, err
	}

	// 5. Hitung HPP tenaga kerja
	hppPekerja, err := s.ProdukRepo.SummarizeHPPTenagaKerja(produkID)
	if err != nil {
		return nil, err
	}

	// 6. Total HPP
	totalHPP := hppBahan + hppPekerja

	// 7. Hitung harga jual (dengan margin 30%)
	hargaJual := totalHPP * 1.3

	// 8. Update produk
	err = s.ProdukRepo.UpdateHargaProduk(produkID, hargaJual, hppBahan, hppPekerja)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id_produk":      produkID,
		"harga_jual":     hargaJual,
		"total_hpp":      totalHPP,
		"hpp_bahan_baku": hppBahan,
		"hpp_pekerja":    hppPekerja,
		"margin":         "30%",
		"jumlah_bahan":   len(input.Resep),
		"jumlah_pekerja": len(input.TenagaKerja),
	}, nil
}
