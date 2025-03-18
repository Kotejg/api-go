package security

import "golang.org/x/crypto/bcrypt"

// recebe um string e coloca o hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// compara uma ssenha e um hash e retorna a validacao delas
func Verify(senhaComHash string, senhaStr string) error {
	err := bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaStr))
	return err
}
