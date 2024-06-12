package main

import (
	"LocalFood/config"
	"LocalFood/controllers/artikelcontroller"
	"LocalFood/controllers/formcontroller"
	"LocalFood/controllers/homecontroller"
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
	http.HandleFunc("/formArtikel", formcontroller.Index)

	// Melayani file statis dari direktori "uploads"
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// Mengatur handler untuk URL path "/img_background/" untuk melayani file statis
	http.Handle("/img_background/", http.StripPrefix("/img_background/", http.FileServer(http.Dir("assets/img_background"))))

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
