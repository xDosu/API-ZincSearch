package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SearchRequest define la estructura de la solicitud de búsqueda.
type SearchRequest struct {
	Query  string                 `json:"query"`  // Búsqueda general
	Filters map[string]string     `json:"filters"` // Filtros por criterio, como tema o país
}

// SearchResponse define la estructura de la respuesta de búsqueda.
type SearchResponse struct {
	Hits []map[string]interface{} `json:"hits"` // Documentos encontrados
}

// Search realiza una búsqueda general o con filtros en ZincSearch.
func Search(searchRequest SearchRequest, user, password string) (*SearchResponse, error) {
	// Construir la consulta base
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"content": searchRequest.Query, // Cambia "content" por el campo relevante
						},
					},
				},
			},
		},
	}

	// Agregar filtros a la consulta (si existen)
	if len(searchRequest.Filters) > 0 {
		filterClauses := []map[string]interface{}{}
		for field, value := range searchRequest.Filters {
			filterClauses = append(filterClauses, map[string]interface{}{
				"term": map[string]interface{}{
					field: value,
				},
			})
		}
		// Insertar filtros en el bloque bool
		searchQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = filterClauses
	}

	// Codificar la consulta como JSON
	queryBody, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, fmt.Errorf("error al codificar la consulta: %w", err)
	}

	// Hacer la solicitud al motor de búsqueda
	searchURL := fmt.Sprintf("%s/api/_search", zincBaseURL)
	req, err := http.NewRequest("POST", searchURL, bytes.NewReader(queryBody))
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud de búsqueda: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la búsqueda: %w", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta del servidor
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %w", err)
	}

	// Validar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la búsqueda: %d - %s", resp.StatusCode, string(body))
	}

	// Decodificar la respuesta en SearchResponse
	var searchResponse SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %w", err)
	}

	return &searchResponse, nil
}
