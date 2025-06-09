package models

import "time"

type BahanBaku struct {
	IDBahan        int        `json:"id_bahan"`
	NamaBahan      string     `json:"nama_bahan"`
	Satuan         string     `json:"satuan"`
	HargaPerSatuan float64    `json:"harga_per_satuan"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}
