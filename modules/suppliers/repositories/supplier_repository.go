package repositories

import (
	"database/sql"
	"storage-api/modules/suppliers/models"
)

type SupplierRepository struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

// Salvar fornecedor
func (r *SupplierRepository) Create(s *models.Supplier) error {
	query := `
		INSERT INTO suppliers (name)
		OUTPUT INSERTED.id
		VALUES (@p1)
	`
	return r.db.QueryRow(query, s.Name).Scan(&s.ID)
}

// Buscar fornecedor por ID
func (r *SupplierRepository) GetByID(id int64) (*models.Supplier, error) {
	query := `SELECT id, name FROM suppliers WHERE id = @p1`

	var s models.Supplier
	err := r.db.QueryRow(query, id).Scan(&s.ID, &s.Name)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Listar todos os fornecedores
func (r *SupplierRepository) List() ([]*models.Supplier, error) {
	query := `SELECT id, name FROM suppliers`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []*models.Supplier
	for rows.Next() {
		var s models.Supplier
		rows.Scan(&s.ID, &s.Name)
		suppliers = append(suppliers, &s)
	}

	return suppliers, nil
}

// Atualizar fornecedor
func (r *SupplierRepository) Update(s *models.Supplier) error {
	query := `UPDATE suppliers SET name = @p1 WHERE id = @p2`
	_, err := r.db.Exec(query, s.Name, s.ID)
	return err
}

// Deletar fornecedor
func (r *SupplierRepository) Delete(id int64) error {
	query := `DELETE FROM suppliers WHERE id = @p1`
	_, err := r.db.Exec(query, id)
	return err
}
