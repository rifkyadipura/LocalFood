package artikelcontroller

import (
	"LocalFood/models/artikelmodel"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	artikel := artikelmodel.GetAll()

	data := map[string]any{
		"Artikel": artikel,
	}
	temp, err := template.ParseFiles("views/artikel/listArtikel.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
