package main

import (
	"mysearch-api/handlers"
	"mysearch-api/utils"
)

func main() {
	// // Cargar datos del índice desde un archivo JSON
	// indexData, err := handlers.CreateIndexerFromJsonFile("output.ndjson")

	// if err != nil {
	// 	log.Fatalf("Error loading index data: %v", err)
	// }

	// // Subir el índice a ZincSearch
	// if err := handlers.UploadIndexOnZincSearch(indexData, "admin", "admin"); err != nil {
	// 	log.Fatalf("Error uploading index: %v", err)
	// }

	// index, err := utils.ReadJSON("index.json")

	// if(err == nil) {
	// 	handlers.CreateIndex(index, "admin", "admin")
	// }

	//handlers.DeleteIndex("admin", "admin", "email")


	ndjsonContent, err := utils.ReadJSON("output.ndjson")
	if (err == nil) {
		handlers.BulkInsertion(ndjsonContent, "admin", "admin")
	}
}
