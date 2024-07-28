package artikelcontroller

import (
	"LocalFood/entities"
	"LocalFood/models/artikelmodel"
	"LocalFood/models/usersmodel"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	artikel := artikelmodel.GetSelesai()

	data := map[string]any{
		"Artikel": artikel,
	}
	temp, err := template.ParseFiles("views/artikel/listArtikel.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func FormArtikel(w http.ResponseWriter, r *http.Request) {
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

		// Mengambil username dari cookie
		cookie, err := r.Cookie("username")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		username := cookie.Value

		// Mendapatkan user berdasarkan username
		user, err := usersmodel.GetByUsername(username)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		formArtikel.Id_users = user.Id_users

		if ok := artikelmodel.Create(formArtikel); !ok {
			temp, err := template.ParseFiles("views/artikel/listArtikel.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			temp.Execute(w, nil)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func VerifikasiForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idArtikel := r.FormValue("id_artikel")
		action := r.FormValue("action")

		var err error
		if action == "upload" {
			err = artikelmodel.UpdateStatusToSelesai(idArtikel)
		} else if action == "ditolak" {
			err = artikelmodel.UpdateStatusToDitolak(idArtikel)
		} else if action == "menunggu" {
			err = artikelmodel.UpdateStatusToMenunggu(idArtikel)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/verifikasiForm", http.StatusSeeOther)
		return
	}

	artikel := artikelmodel.GetAll()

	data := map[string]interface{}{
		"Artikel": artikel,
	}
	temp, err := template.ParseFiles("views/dashboard/admin/verifikasi_daftar.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func DeleteArtikelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idArtikel := r.FormValue("id_artikel")

		err := artikelmodel.DeleteArtikel(idArtikel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/verifikasiForm", http.StatusSeeOther)
		return
	}
	http.NotFound(w, r)
}
