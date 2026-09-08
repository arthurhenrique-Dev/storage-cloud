package repositories

import (
	"database/sql"
	"storage-api/modules/users/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Salvar usuário
func (r *UserRepository) Create(u *models.User) error {
	query := `
		INSERT INTO users (email, first_name, last_name, password_hash)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2, @p3, @p4)
	`
	return r.db.QueryRow(query, u.Email, u.FirstName, u.LastName, u.PasswordHash).Scan(&u.ID)
}

// Buscar por ID
func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash FROM users WHERE id = @p1`

	var u models.User
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Buscar por Email (muito comum em autenticação/login)
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash FROM users WHERE email = @p1`

	var u models.User
	err := r.db.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Listar todos os usuários
func (r *UserRepository) List() ([]*models.User, error) {
	query := `SELECT id, email, first_name, last_name FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName)
		users = append(users, &u)
	}

	return users, nil
}

// Deletar usuário
func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = @p1`
	_, err := r.db.Exec(query, id)
	return err
}
