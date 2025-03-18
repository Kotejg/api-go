package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// cria novo repositprio de usuario\
func NovoRepositoryDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// insere no banco de dados
func (repo usuarios) Criar(u models.Usuario) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES(?, ?, ?, ?) ",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(u.Nome, u.Nick, u.Email, u.Senha)
	if err != nil {
		return 0, err
	}
	ultimoID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoID), nil

}

func (repo usuarios) Buscar(filter string) ([]models.Usuario, error) {
	filtro := fmt.Sprintf("%%%s%%", filter)

	linhas, err := repo.db.Query(
		`select id, nome, nick, email, criadoEm
				from usuarios
				where nome LIKE ? or nick LIKE ?`,
		filtro,
		filtro,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []models.Usuario
	for linhas.Next() {
		var u models.Usuario
		if err = linhas.Scan(
			&u.ID,
			&u.Nome,
			&u.Nick,
			&u.Email,
			&u.CriadoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

func (repo usuarios) BuscarPorId(id uint) (models.Usuario, error) {
	row, err := repo.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		id)
	if err != nil {
		return models.Usuario{}, err
	}
	defer row.Close()
	var u models.Usuario
	if row.Next() {
		if err = row.Scan(
			&u.ID,
			&u.Nome,
			&u.Nick,
			&u.Email,
			&u.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}

	}
	return u, nil
}

func (repo *usuarios) Atualizar(ID int, u models.Usuario) error {
	statement, err := repo.db.Prepare(
		`update usuarios set nome = ?, nick = ?, email = ? where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(u.Nome, u.Nick, u.Email, ID); err != nil {
		return err
	}
	return nil
}

func (repo *usuarios) Deletar(id int) error {
	statement, err := repo.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(id); err != nil {
		return err
	}
	return nil
}

func (repo *usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	row, erro := repo.db.Query("select id,senha from usuarios where email = ? ", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer row.Close()
	var u models.Usuario
	if row.Next() {
		if erro := row.Scan(&u.ID, &u.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return u, nil

}

// seguir realizar o ate de um usuario seguir outro
func (repo *usuarios) Seguir(IDToken, IDSeguidor uint64) error {
	statement, err := repo.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?,?) ",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(IDToken, IDSeguidor); err != nil {
		return err
	}
	return nil
}

func (repo *usuarios) PararDeSeguir(IDToken, IDSeguidor uint64) error {
	statement, err := repo.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(IDToken, IDSeguidor); err != nil {
		return err
	}
	return nil
}

func (repo *usuarios) BuscarSeguidores(usuarioID int) ([]models.Usuario, error) {
	var usuarios []models.Usuario

	statement, err := repo.db.Prepare(
		`select u.id, u.nome, u.nick, u.email, u.criadoEm
					from usuarios u 
				inner join seguidores s on u.id = s.seguidor_id 
					where s.usuario_id = ?`,
	)
	if err != nil {
		return usuarios, err
	}
	defer statement.Close()
	rows, err := statement.Query(usuarioID)
	if err != nil {
		return usuarios, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.Usuario
		if err = rows.Scan(&u.ID, &u.Nome, &u.Nick, &u.Email, &u.CriadoEm); err != nil {
			return usuarios, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

func (repo *usuarios) BuscarSeguindo(usuarioID int) ([]models.Usuario, error) {
	var usuarios []models.Usuario

	statement, err := repo.db.Prepare(
		`select u.id, u.nome, u.nick, u.email, u.criadoEm
					from usuarios u 
				inner join seguidores s on u.id = s.usuario_id 
					where s.seguidor_id = ?`,
	)
	if err != nil {
		return usuarios, err
	}
	defer statement.Close()
	rows, err := statement.Query(usuarioID)
	if err != nil {
		return usuarios, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.Usuario
		if err = rows.Scan(&u.ID, &u.Nome, &u.Nick, &u.Email, &u.CriadoEm); err != nil {
			return usuarios, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

func (repo *usuarios) BuscarSenhaPorID(UsuarioID int) (string, error) {
	row, err := repo.db.Query(`
	select senha from usuarios where id = ?
	`, UsuarioID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var u models.Usuario

	if row.Next() {
		if err = row.Scan(&u.Senha); err != nil {
			return "", err
		}
	}
	return u.Senha, nil
}

func (repo *usuarios) AtualizarSenha(usuarioID int, hash string) error {
	stt, err := repo.db.Query(`
	update usuarios set senha = ? where id = ?
	`, hash, usuarioID)
	if err != nil {
		return err
	}
	defer stt.Close()

	return nil

}
