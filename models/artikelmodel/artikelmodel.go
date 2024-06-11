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
		if err := rows.Scan(&Article.Id, &Article.Judul, &Article.Isi, &Article.Foto, &Article.Alamat, &Article.Maps, &Article.CreatedAt, &Article.UpdatedAt); err != nil {
			panic(err)
		}

		artikel = append(artikel, Article)
	}

	return artikel
}

func Create(formArtikel entities.Artikel) bool {
	result, err := config.DB.Exec(`
	INSERT INTO artikel (judul, isi, alamat, maps, foto, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		formArtikel.Judul, formArtikel.Isi, formArtikel.Alamat, formArtikel.Maps, formArtikel.Foto, formArtikel.CreatedAt, formArtikel.UpdatedAt,
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
