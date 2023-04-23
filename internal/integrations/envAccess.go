package integrations

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVar() error {
	envPath, err := getEnvFile()
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(envPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(absPath)
	if err != nil {
		return err
	}

	return nil
}

func getEnvFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Busca la raíz del proyecto
	projectRoot, err := findProjectRoot(currentDir)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Construye la ruta del archivo .env en la raíz del proyecto
	envFilePath := filepath.Join(projectRoot, ".env")
	return envFilePath, nil

}

func findProjectRoot(currentDir string) (string, error) {
	for {
		goModPath := filepath.Join(currentDir, "go.mod")

		// Verifica si el archivo go.mod existe en el directorio actual
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir, nil
		}

		// Si hemos llegado al directorio raíz y no encontramos go.mod
		if currentDir == filepath.Dir(currentDir) {
			break
		}

		// Sube un nivel en la estructura de directorios
		currentDir = filepath.Dir(currentDir)
	}

	return "", fmt.Errorf("no se encontró el archivo go.mod")
}
