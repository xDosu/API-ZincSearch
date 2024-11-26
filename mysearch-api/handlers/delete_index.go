package handlers

import (
	"net/http"
	"fmt"
	"io"
	"errors"
)

func DeleteIndex(user, password, index string) error {
	// Crear la solicitud HTTP
	zincURL := fmt.Sprintf("%s/api/index/%s", zincBaseURL, index)

	req, err := http.NewRequest("DELETE", zincURL, nil)
	if err != nil {
		return err
	}

	// Configurar encabezados y autenticación
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Close = true

	// Enviar la solicitud
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}