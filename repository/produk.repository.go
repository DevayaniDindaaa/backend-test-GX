package repository

import (
	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type ProdukRepository struct{}

func (r *ProdukRepository) GetAllProduk() ([]models.Produk, error) {
	query := `
		SELECT id_produk, nama_produk, harga_jual,
		       created_at, updated_at, deleted_at
		FROM produk
		WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produkList []models.Produk

	for rows.Next() {
		var b models.Produk
		err := rows.Scan(
			&b.IDProduk,
			&b.NamaProduk,
			&b.HargaJual,
			&b.CreatedAt,
			&b.UpdatedAt,
			&b.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		produkList = append(produkList, b)
	}

	return produkList, nil
}
