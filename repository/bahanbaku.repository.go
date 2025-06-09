package repository

import (
	"fmt"

	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type BahanBakuRepository struct{}

func (r *BahanBakuRepository) GetAllBahanBaku() ([]models.BahanBaku, error) {
	query := `
		SELECT id_bahan, nama_bahan, satuan, harga_per_satuan,
		       created_at, updated_at, deleted_at
		FROM bahan_baku
		WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bahanList []models.BahanBaku

	for rows.Next() {
		var b models.BahanBaku
		err := rows.Scan(
			&b.IDBahan,
			&b.NamaBahan,
			&b.Satuan,
			&b.HargaPerSatuan,
			&b.CreatedAt,
			&b.UpdatedAt,
			&b.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		bahanList = append(bahanList, b)
	}

	return bahanList, nil
}

func (r *BahanBakuRepository) CreateBahanBaku(bahan *models.BahanBaku) error {
	query := `INSERT INTO bahan_baku (nama_bahan, satuan, harga_per_satuan, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id_bahan`

	err := db.DB.QueryRow(query, bahan.NamaBahan, bahan.Satuan, bahan.HargaPerSatuan).Scan(&bahan.IDBahan)

	if err != nil {
		return fmt.Errorf("could not create bahan: %v", err)
	}
	return nil
}
