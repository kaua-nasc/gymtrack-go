package config

import "github.com/joho/godotenv"

func LoadEnvironmentVariable() {
	err := godotenv.Load()
	if err != nil {
		panic("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
}
