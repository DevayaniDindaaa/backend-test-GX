package repository

import (
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
