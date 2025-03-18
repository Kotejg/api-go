package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// preparar chama os metodos para validar e formaatar os metodos
func (u *Usuario) PrepararUsuario(etapa string) error {
	if erro := u.Validar(etapa); erro != nil {
		return erro
	}
	if erro := u.Formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (u *Usuario) Validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("Nome is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("Formato de email invalido")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("Senha is required")
	}
	return nil
}

func (u *Usuario) Formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)

	if etapa == "cadastro" {
		senhaComHash, err := security.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(senhaComHash)
	}
	return nil
}
