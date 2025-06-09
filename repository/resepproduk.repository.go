package repository

import (
	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type ResepRepository struct{}

func (r *ResepRepository) GetAllResep() ([]models.ResepProduk, error) {
	query := `
		SELECT id_produk, id_resep, id_bahan, jumlah,
		       created_at, updated_at, deleted_at
		FROM resep_produk
		WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resepList []models.ResepProduk

	for rows.Next() {
		var b models.ResepProduk
		err := rows.Scan(
			&b.IDProduk,
			&b.IDResep,
			&b.IDBahan,
			&b.Jumlah,
			&b.CreatedAt,
			&b.UpdatedAt,
			&b.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		resepList = append(resepList, b)
	}

	return resepList, nil
}
