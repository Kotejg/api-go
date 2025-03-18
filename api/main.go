package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//func init() {
//	chave := make([]byte, 64)
//	if _, err := rand.Read(chave); err != nil {
//		log.Fatal(err)
//	}
//
//	strBase64 := base64.StdEncoding.EncodeToString(chave)
//	fmt.Println(strBase64)
//}

func main() {
	config.Carregar()
	fmt.Println(config.Port)
	fmt.Println(config.StringConexaoDB)
	fmt.Println("Rodando api")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
