package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route untuk melayani halaman utama
	r.HandleFunc("/", Loadweb)

	// Static file handler untuk gambar dan file CV
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Menjalankan server di port 8080
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

// Handler untuk halaman index
func Loadweb(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
