package repository

import (
	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type TransaksiRepository struct{}

func (r *TransaksiRepository) GetAllTransaksi() ([]models.TransaksiPenjualan, error) {
	query := `
		SELECT id_transaksi, tanggal, id_produk, jumlah, total_harga,
		       created_at, updated_at, deleted_at
		FROM transaksi_penjualan
		WHERE deleted_at IS NULL
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transaksiList []models.TransaksiPenjualan

	for rows.Next() {
		var b models.TransaksiPenjualan
		err := rows.Scan(
			&b.IDTransaksi,
			&b.Tanggal,
			&b.IDProduk,
			&b.Jumlah,
			&b.TotalHarga,
			&b.CreatedAt,
			&b.UpdatedAt,
			&b.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		transaksiList = append(transaksiList, b)
	}

	return transaksiList, nil
}
