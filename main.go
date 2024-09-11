package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/details", detailsHandler)

	/*http.HandleFunc("/details/docker", docHandler)
	http.HandleFunc("/details/docker-compose", docComHandler)
	http.HandleFunc("/details/kubernetes", kubeHandler)
	http.HandleFunc("/details/DevOps", doHandler)
	http.HandleFunc("/details/linux", linHandler)
	http.HandleFunc("/details/gits", gitsHandler)
	http.HandleFunc("/details/languages", lanpHandler)
	http.HandleFunc("/details/oti", otiHandler)*/

	http.HandleFunc("/details/examp", exampHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "details.html")
}

func exampHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "examp.html")
}

func renderTemplate(w http.ResponseWriter, filename string) {
	path := filepath.Join("templates", filename)
	html, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(html))
}
