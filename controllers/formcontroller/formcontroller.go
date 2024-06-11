package formcontroller

import (
	"LocalFood/entities"
	"LocalFood/models/artikelmodel"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/form/form.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		var formArtikel entities.Artikel

		formArtikel.Judul = r.FormValue("judul")
		formArtikel.Isi = r.FormValue("isi")
		formArtikel.Alamat = r.FormValue("alamat")
		formArtikel.Maps = r.FormValue("maps")
		formArtikel.CreatedAt = time.Now()
		formArtikel.UpdatedAt = time.Now()

		// Menghandle upload file
		file, header, err := r.FormFile("foto")
		if err == nil {
			defer file.Close()

			// Tentukan direktori tempat penyimpanan file
			uploadDir := "./uploads"
			// Membuat direktori jika belum ada
			err := os.MkdirAll(uploadDir, os.ModePerm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Tentukan path lengkap untuk menyimpan file
			filePath := filepath.Join(uploadDir, header.Filename)

			dst, err := os.Create(filePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			formArtikel.Foto = "/uploads/" + header.Filename // Path penyimpanan file
		} else if err != http.ErrMissingFile {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if ok := artikelmodel.Create(formArtikel); !ok {
			temp, err := template.ParseFiles("views/Artikel/artikel.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			temp.Execute(w, nil)
			return
		}

		http.Redirect(w, r, "/listArtikel", http.StatusSeeOther)
	}
}
