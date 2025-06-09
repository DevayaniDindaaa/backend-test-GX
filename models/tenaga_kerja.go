package models

import "time"

type TenagaKerja struct {
	IDTenaga    int        `json:"id_tenaga"`
	NamaTenaga  string     `json:"nama_tenaga"`
	BiayaPerJam float64    `json:"biaya_per_jam"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
