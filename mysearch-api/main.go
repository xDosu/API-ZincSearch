package main

import (
	"encoding/json"
	"fmt"
	"mysearch-api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Decodificar la solicitud
		var searchRequest handlers.SearchRequest
		err := json.NewDecoder(r.Body).Decode(&searchRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error al decodificar la solicitud: %v", err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Realizar la búsqueda
		searchResponse, err := handlers.Search(searchRequest, "admin", "admin")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error en la búsqueda: %v", err), http.StatusInternalServerError)
			return
		}

		// Enviar la respuesta al cliente
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchResponse)
	})

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
