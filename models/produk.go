package models

import "time"

type Produk struct {
	IDProduk   int        `json:"id_produk"`
	NamaProduk string     `json:"nama_produk"`
	HargaJual  float64    `json:"harga_jual"`
	HPP        float64    `json:"hpp"`
	HPPPekerja float64    `json:"hpp_pekerja"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

type ProdukInput struct {
	NamaProduk  string       `json:"nama_produk"`
	Resep       []ResepInput `json:"resep"`
	TenagaKerja []WaktuInput `json:"tenaga_kerja"`
}
