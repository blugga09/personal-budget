package config

import (
	"encoding/json"
	"log"
	"os"
)

func loadFile(file string) map[string]string {

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	var values map[string]string

	err = json.Unmarshal(content, &values)
	if err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %v", err)
	}

	return values
}
