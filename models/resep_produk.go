package models

import "time"

type ResepProduk struct {
	IDResep   int        `json:"id_resep"`
	IDProduk  int        `json:"id_produk"`
	IDBahan   int        `json:"id_bahan"`
	Jumlah    float64    `json:"jumlah"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ResepInput struct {
	IDBahan int     `json:"id_bahan"`
	Jumlah  float64 `json:"jumlah"`
}
