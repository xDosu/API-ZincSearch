package utils

import(
	"encoding/json"
	"bytes"
	"os"
	"fmt"
	"io"
)


func ReadJSON(filepath string) (string, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(jsonFile)
	var bufferContent bytes.Buffer

	for {
		var doc map[string]interface{}
		err := decoder.Decode(&doc)
		if err == io.EOF {
			break // Fin del archivo
		}
		if err != nil {
			return "", fmt.Errorf("error al decodificar NDJSON: %w", err)
		}
		// Añadir el documento JSON como línea
		docJSON, err := json.Marshal(doc)
		if err != nil {
			return "", fmt.Errorf("error al codificar documento JSON: %w", err)
		}
		bufferContent.Write(docJSON)
		bufferContent.WriteString("\n")
	}
	return bufferContent.String(), nil
}