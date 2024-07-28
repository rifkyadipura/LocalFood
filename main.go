package main

import (
	"LocalFood/config"
	"LocalFood/controllers/artikelcontroller"
	"LocalFood/controllers/homecontroller"
	"LocalFood/controllers/userscontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Hompage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. List Artikel
	http.HandleFunc("/listArtikel", artikelcontroller.Index)

	// 3. form Artikel
	http.HandleFunc("/formArtikel", artikelcontroller.FormArtikel)

	// 4. Login
	http.HandleFunc("/login", userscontroller.Login)

	// 5. Registrasi
	http.HandleFunc("/registrasi", userscontroller.Registrasi)

	// 6. Dashboard
	http.HandleFunc("/dashboard", userscontroller.Dashboard)

	// 7. Logout
	http.HandleFunc("/logout", userscontroller.Logout)

	// 8. Verifikasi Form
	http.HandleFunc("/verifikasiForm", artikelcontroller.VerifikasiForm)

	// 9. Delete Artikel
	http.HandleFunc("/deleteArtikel", artikelcontroller.DeleteArtikelHandler)

	// Melayani file statis dari direktori "uploads"
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// Mengatur handler untuk URL path "/img_background/" untuk melayani file statis
	http.Handle("/img_background/", http.StripPrefix("/img_background/", http.FileServer(http.Dir("assets/img_background"))))

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
