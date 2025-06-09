package models

import "time"

type TransaksiPenjualan struct {
	IDTransaksi int        `json:"id_transaksi"`
	Tanggal     time.Time  `json:"tanggal"`
	IDProduk    int        `json:"id_produk"`
	Jumlah      float32    `json:"jumlah"`
	TotalHarga  float64    `json:"total_harga"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
