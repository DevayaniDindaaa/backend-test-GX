package repository

import (
	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type WaktuKerjaRepository struct{}

func (r *WaktuKerjaRepository) GetAllWaktuKerja() ([]models.WaktuPembuatan, error) {
	query := `
		SELECT id_waktu, id_produk, jam_kerja, id_tenaga, 
		       created_at, updated_at, deleted_at
		FROM waktu_pembuatan
		WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bahanList []models.WaktuPembuatan

	for rows.Next() {
		var b models.WaktuPembuatan
		err := rows.Scan(
			&b.IDWaktu,
			&b.IDProduk,
			&b.JamKerja,
			&b.IDTenaga,
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
