package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Initialize initializes the environment variables
//
// Business logic:
// - Loads .env file (by default, no need to provide path)
// - Loads env files from the provided paths, if any
//
// Parameters:
// - envFilePath: slice of strings representing the paths to the .env files to load
//
// Returns:
// - None
func Initialize(envFilePath ...string) {
	paths := []string{".env"}

	paths = append(paths, envFilePath...)

	for _, path := range paths {
		if fileExists(path) {
			err := godotenv.Load(path)
			if err != nil {
				log.Fatal("Error loading " + path + " file")
			}
		}
	}
}
