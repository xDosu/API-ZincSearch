package handlers

import (
	"fmt"
	"net/http"
	"io"
	"strings"
)

func BulkInsertion(jsonContent, user, password string) error {
	// Crear la solicitud HTTP
	zincURL := fmt.Sprintf("%s/api/_bulk", zincBaseURL)

	req, err := http.NewRequest("POST", zincURL,  strings.NewReader(jsonContent))

	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Configurar encabezados y autenticación
	req.Header.Set("Content-Type", "application/x-ndjson")
	req.SetBasicAuth(user, password)
	req.Close = true

	// Enviar la solicitud
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to ZincSearch: %w", err)
	}
	defer resp.Body.Close()

	// Validar el código de respuesta
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error en la inserción masiva: %d - %s", resp.StatusCode, string(body))
	}
	return nil
}
