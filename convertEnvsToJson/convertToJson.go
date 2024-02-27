package convertenvstojson

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ParseToEnvs(jsonInput string) {

	// Um mapa para armazenar as chaves e valores convertidos
	var envVars map[string]string

	// Decodificando a string JSON para o mapa
	err := json.Unmarshal([]byte(strings.TrimSpace(jsonInput)), &envVars)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		os.Exit(1)
	}

	// Iterando sobre o mapa e imprimindo no formato CHAVE=VALOR
	for key, value := range envVars {
		fmt.Printf("%s=%s\n", key, value)
	}
}

func ParseToJson(envInput string) {

}
