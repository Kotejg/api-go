package authorization

import (
	"api/src/config"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 5).Unix()
	permissoes["usuarioID"] = usuarioID

	//secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

func ValidarToken(r *http.Request) error {
	tokenStr := obterToken(r)
	token, err := jwt.Parse(tokenStr, RetornaChaveVerificacao)
	if err != nil {
		return err
	}
	fmt.Println(token)
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token invalidado")
}

// obtem o token do header da requisicao e ritira po bearer
func obterToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// retira o ID de Usuario do campo de token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenStr := obterToken(r)
	token, err := jwt.Parse(tokenStr, RetornaChaveVerificacao)
	if err != nil {
		return 0, err
	}
	if permissaos, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// se der erro bou precisar converter o permissoes para o tipo float e depois converter ele para tipo unti64
		UsuarioID, erro := strconv.ParseUint(fmt.Sprint(permissaos["usuarioID"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return UsuarioID, nil

	}
	return 0, errors.New("Token invalidado")
}

func RetornaChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado: %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
