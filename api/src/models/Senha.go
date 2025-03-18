package models

// representa a requisicão de formato de alteração de senha
type Senha struct {
	NewPassword string `json:"nova,omitempty"`
	OldPassword string `json:"atual,omitempty"`
}
