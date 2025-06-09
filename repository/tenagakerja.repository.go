package repository

import (
	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type TenagaKerjaRepository struct{}

func (r *TenagaKerjaRepository) GetAllTenagaKerja() ([]models.TenagaKerja, error) {
	query := `
		SELECT id_tenaga, nama_tenaga, biaya_per_jam,
		       created_at, updated_at, deleted_at
		FROM tenaga_kerja
		WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bahanList []models.TenagaKerja

	for rows.Next() {
		var b models.TenagaKerja
		err := rows.Scan(
			&b.IDTenaga,
			&b.NamaTenaga,
			&b.BiayaPerJam,
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
