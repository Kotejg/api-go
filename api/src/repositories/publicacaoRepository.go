package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (r *Publicacoes) Criar(publi models.Publicacoes) (uint64, error) {
	stt, err := r.db.Prepare(`
	insert into publicacoes (titulo, conteudo, autor_id) 
	values (?,?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stt.Close()
	result, err := stt.Exec(publi.Titulo, publi.Conteudo, publi.AutorID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
