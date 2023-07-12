package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
}

func main() {
	// Configuración del servidor web
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// Maneja todas las solicitudes a la ruta raíz ("/") y sirve los archivos estáticos de la carpeta "./public"

	http.HandleFunc("/api/pokemon", func(w http.ResponseWriter, r *http.Request) {
		// Simulación de datos de la Pokédex
		pokemon := Pokemon{
			ID:      25,
			Name:    "Pikachu",
			Species: "Pokémon ratón",
		}

		// Convertir la estructura Pokemon a JSON
		jsonData, err := json.Marshal(pokemon)
		if err != nil {
			http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
			return
		}

		// Establecer la cabecera de respuesta para indicar el tipo de contenido JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	// Iniciar el servidor web en el puerto 8080
	fmt.Println("Servidor iniciado en http://localhost:8080/index.html")
	http.ListenAndServe(":8080", nil)
}
