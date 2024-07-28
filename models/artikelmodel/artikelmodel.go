package artikelmodel

import (
	"LocalFood/config"
	"LocalFood/entities"
)

func GetAll() []entities.Artikel {
	rows, err := config.DB.Query("SELECT * FROM artikel ORDER BY created_at DESC")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var artikel []entities.Artikel

	for rows.Next() {
		var Article entities.Artikel
		if err := rows.Scan(&Article.Id_artikel, &Article.Id_users, &Article.Judul, &Article.Isi, &Article.Foto, &Article.Alamat, &Article.Maps, &Article.Status, &Article.CreatedAt, &Article.UpdatedAt); err != nil {
			panic(err)
		}

		artikel = append(artikel, Article)
	}

	return artikel
}

func GetFormMenunggu(idUser int) []entities.Artikel {
	// Query SQL dengan filter berdasarkan id_users
	query := `SELECT * FROM artikel WHERE id_users = ? ORDER BY created_at DESC`

	// Menjalankan query
	rows, err := config.DB.Query(query, idUser)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var artikel []entities.Artikel

	for rows.Next() {
		var Article entities.Artikel
		if err := rows.Scan(&Article.Id_artikel, &Article.Id_users, &Article.Judul, &Article.Isi, &Article.Foto, &Article.Alamat, &Article.Maps, &Article.Status, &Article.CreatedAt, &Article.UpdatedAt); err != nil {
			panic(err)
		}

		artikel = append(artikel, Article)
	}

	return artikel
}

func GetSelesai() []entities.Artikel {
	rows, err := config.DB.Query("SELECT * FROM artikel WHERE status = 'selesai' ORDER BY created_at DESC")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var artikel []entities.Artikel

	for rows.Next() {
		var Article entities.Artikel
		if err := rows.Scan(&Article.Id_artikel, &Article.Id_users, &Article.Judul, &Article.Isi, &Article.Foto, &Article.Alamat, &Article.Maps, &Article.Status, &Article.CreatedAt, &Article.UpdatedAt); err != nil {
			panic(err)
		}

		artikel = append(artikel, Article)
	}

	return artikel
}

func Create(formArtikel entities.Artikel) bool {
	result, err := config.DB.Exec(`
	INSERT INTO artikel (id_users, judul, isi, alamat, maps, foto, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		formArtikel.Id_users, formArtikel.Judul, formArtikel.Isi, formArtikel.Alamat, formArtikel.Maps, formArtikel.Foto, formArtikel.CreatedAt, formArtikel.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertID > 0
}

func UpdateStatusToSelesai(id string) error {
	_, err := config.DB.Exec("UPDATE artikel SET status = 'selesai' WHERE id_artikel = ?", id)
	return err
}

func UpdateStatusToDitolak(id string) error {
	db := config.DB
	_, err := db.Exec("UPDATE artikel SET status = 'ditolak' WHERE id_artikel = ?", id)
	return err
}

func UpdateStatusToMenunggu(id string) error {
	db := config.DB
	_, err := db.Exec("UPDATE artikel SET status = 'menunggu' WHERE id_artikel = ?", id)
	return err
}

func DeleteArtikel(id string) error {
	_, err := config.DB.Exec("DELETE FROM artikel WHERE id_artikel = ?", id)
	return err
}
