package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// carregar as variaveis de ambiente

var (
	//string de conexao com o mysql
	StringConexaoDB = ""
	//porta q api vai esta rodando
	Port = 0
	//chave de assinatura de token
	SecretKey []byte
)

func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 900
	}
	StringConexaoDB = fmt.Sprintf("%s:%s@/%s?parseTime=True&charset=utf8&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
