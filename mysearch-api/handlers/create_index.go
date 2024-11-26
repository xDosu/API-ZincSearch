package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

const zincBaseURL = "http://localhost:4080"

func CreateIndex(jsonContent, user, password string) error {
	// Crear la solicitud HTTP
	zincURL := fmt.Sprintf("%s/api/index/", zincBaseURL)
	req, err := http.NewRequest("POST", zincURL, strings.NewReader(jsonContent))

	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Configurar encabezados y autenticación
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, password)
	req.Close = true

	// Enviar la solicitud
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to ZincSearch: %w", err)
	}
	defer resp.Body.Close()

	// Validar el código de respuesta
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create index, status code: %d", resp.StatusCode)
	}
	return nil
}