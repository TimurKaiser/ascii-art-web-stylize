package main

import (
	"ascii-art-web/controllers"
	"fmt"
	"net/http"
)

func main() {
	// Définir le gestionnaire pour servir le fichier HTML
	http.HandleFunc("/", controllers.HomePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur démarré sur http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
	