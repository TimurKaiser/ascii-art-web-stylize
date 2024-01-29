package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodPost {
			// Récupérer les données du formulaire
			text := r.FormValue("inputText")
			font := r.FormValue("fontSelect")
			//check si les caracteres font partie des printables
			for i := 0; i < len(text); i++ {
				if text[i] < ' ' && text[i] != '\n' && text[i] != '\r'|| text[i] > '~' {
					http.Error(w, "Bad Request", http.StatusBadRequest)
					return
				}
			}

			// Obtenir les lignes de la police spécifiée
			lines, err := getFont(font)
			if err != nil {
				http.Error(w, "Page Not Found", http.StatusNotFound)
				return
			}

			// Convertir le texte
			convertedText := convertInput(text, lines)

			// Créer une structure de données pour stocker les variables de la page
			pageVariables := PageVariables{
				ConvertedText: convertedText,
			}

			// Utiliser le modèle pour générer la page HTML
			tmpl, err := template.ParseFiles("static/html/index.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Remplacer la partie {{.ConvertedText}} dans le fichier HTML par le texte converti
			err = tmpl.Execute(w, pageVariables)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			// Si la méthode HTTP n'est pas POST, simplement servir la page HTML
			//http.ServeFile(w, r, "index.html")
			tmpl, err := template.ParseFiles("static/html/index.html")
			if err != nil {
				fmt.Println(err)
				tmpl, _ := template.ParseFiles("static/html/error.html")
				tmpl.Execute(w, "Error 500 : Internal server error")
				return
			}
			tmpl.Execute(w, nil)
		}
	} else {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}
