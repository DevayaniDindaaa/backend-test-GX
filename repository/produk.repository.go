package repository

import (
	"github.com/DevayaniDindaaa/backend-test-GX/db"
	"github.com/DevayaniDindaaa/backend-test-GX/models"
)

type ProdukRepository struct{}

func (r *ProdukRepository) GetAllProduk() ([]models.Produk, error) {
	query := `
		SELECT id_produk, nama_produk, harga_jual, hpp,
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
			&b.HPP,
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

func (r *ProdukRepository) InsertProduk(nama string) (int, error) {
	var id int

	err := db.DB.QueryRow(`
		INSERT INTO produk (nama_produk, harga_jual, hpp, hpp_pekerja, created_at, updated_at)
		VALUES ($1, 0, 0, 0, NOW(), NOW())
		RETURNING id_produk
	`, nama).Scan(&id)
	return id, err
}

func (r *ProdukRepository) InsertResep(idProduk, idBahan int, jumlah float64) error {
	_, err := db.DB.Exec(`
		INSERT INTO resep_produk (id_produk, id_bahan, jumlah, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
	`, idProduk, idBahan, jumlah)
	return err
}

func (r *ProdukRepository) InsertWaktu(idProduk, idTenaga int, jamKerja float64) error {
	_, err := db.DB.Exec(`
		INSERT INTO waktu_pembuatan (id_produk, id_tenaga, jam_kerja, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
	`, idProduk, idTenaga, jamKerja)
	return err
}

func (r *ProdukRepository) SumarizeHPP(idProduk int) (float64, error) {
	var hpp float64

	err := db.DB.QueryRow(`
		 SELECT SUM(
            CASE 
                WHEN bb.satuan = 'gram' THEN rp.jumlah * bb.harga_per_satuan / 1000
                WHEN bb.satuan = 'kg' THEN rp.jumlah * bb.harga_per_satuan
                WHEN bb.satuan = 'ml' THEN rp.jumlah * bb.harga_per_satuan / 1000
                WHEN bb.satuan = 'liter' THEN rp.jumlah * bb.harga_per_satuan
                ELSE rp.jumlah * bb.harga_per_satuan
            END
        )
		FROM resep_produk rp
		JOIN bahan_baku bb ON rp.id_bahan = bb.id_bahan
		WHERE rp.id_produk = $1
	`, idProduk).Scan(&hpp)
	return hpp, err
}

func (r *ProdukRepository) SummarizeHPPTenagaKerja(idProduk int) (float64, error) {
	var hppPekerja float64

	err := db.DB.QueryRow(`
		SELECT COALESCE(SUM(wk.jam_kerja * tk.biaya_per_jam), 0)
		FROM waktu_pembuatan wk
		JOIN tenaga_kerja tk ON wk.id_tenaga = tk.id_tenaga
		WHERE wk.id_produk = $1
	`, idProduk).Scan(&hppPekerja)
	return hppPekerja, err
}

func (r *ProdukRepository) UpdateHargaProduk(idProduk int, harga float64, hpp float64, hppPekerja float64) error {
	_, err := db.DB.Exec(`
		UPDATE produk SET harga_jual = $1, hpp = $2, hpp_pekerja = $3, updated_at = NOW()
		WHERE id_produk = $4
	`, harga, hpp, hppPekerja, idProduk)
	return err
}
