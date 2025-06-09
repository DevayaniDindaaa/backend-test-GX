package models

import "time"

type WaktuPembuatan struct {
	IDWaktu   int        `json:"id_waktu"`
	IDProduk  int        `json:"id_produk"`
	IDTenaga  int        `json:"id_tenaga"`
	JamKerja  float64    `json:"jam_kerja"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type WaktuInput struct {
	IDTenaga int     `json:"id_tenaga"`
	JamKerja float64 `json:"jam_kerja"`
}
