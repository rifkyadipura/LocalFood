package userscontroller

import (
	"LocalFood/entities"
	"LocalFood/models/artikelmodel"
	"LocalFood/models/usersmodel"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"text/template"
	"time"
)

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/auth/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, err := usersmodel.GetByUsername(username)
		if err != nil || hashPassword(password) != user.Password {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Simpan data pengguna dalam cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   user.Username,
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		})

		// Redirect to dashboard after successful login
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func Registrasi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/auth/registrasi.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		var formRegistrasi entities.Users

		formRegistrasi.Email = r.FormValue("email")
		formRegistrasi.Username = r.FormValue("username")
		formRegistrasi.Password = hashPassword(r.FormValue("password"))
		formRegistrasi.CreatedAt = time.Now()
		formRegistrasi.UpdatedAt = time.Now()

		if ok := usersmodel.Create(formRegistrasi); !ok {
			temp, err := template.ParseFiles("views/auth/login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			temp.Execute(w, nil)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// Dapatkan data pengguna dari cookie
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := cookie.Value

	user, err := usersmodel.GetByUsername(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	// Memeriksa role pengguna
	switch user.Role {
	case "admin":
		// Ambil data artikel
		data_users := usersmodel.GetAll()
		data := struct {
			Username   string
			Data_users interface{}
		}{
			Username:   username,
			Data_users: data_users,
		}

		temp, err := template.ParseFiles("views/dashboard/admin/dashboard_admin.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, data)
		// http.Redirect(w, r, "/admin", http.StatusSeeOther)
	default:
		// Ambil data artikel
		artikel := artikelmodel.GetFormMenunggu(int(user.Id_users))

		// Siapkan data untuk template
		data := struct {
			Username string
			Artikel  interface{} // Menggunakan interface{} untuk mendukung tipe data yang bervariasi
		}{
			Username: username,
			Artikel:  artikel,
		}

		// Muat template
		temp, err := template.ParseFiles("views/dashboard/pengguna/dashboard_pengguna.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Jalankan template dengan data
		err = temp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Hapus cookie dengan mengatur waktu kadaluarsa ke waktu yang sudah berlalu
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	})

	// Redirect ke halaman login setelah logout
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
