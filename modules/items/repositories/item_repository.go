package repositories

import (
	"database/sql"
	"storage-api/modules/items/models"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(item *models.Item) error {
	query := `
		INSERT INTO items (name, description, price, supplier_id)
		OUTPUT INSERTED.id
		VALUES (@p1, @p2, @p3, @p4)
	`
	return r.db.QueryRow(query, item.Name, item.Description, item.Price, item.SupplierId).Scan(&item.ID)
}

func (r *ItemRepository) GetByID(id int64) (*models.Item, error) {
	query := `SELECT id, name, description, price, supplier_id FROM items WHERE id = @p1`

	var item models.Item
	err := r.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.Price,
		&item.SupplierId,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) List() ([]*models.Item, error) {
	query := `SELECT id, name, description, price, supplier_id FROM items`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.SupplierId)
		items = append(items, &item)
	}

	return items, nil
}

func (r *ItemRepository) Update(item *models.Item) error {
	query := `
		UPDATE items 
		SET name = @p1, description = @p2, price = @p3, supplier_id = @p4 
		WHERE id = @p5
	`
	_, err := r.db.Exec(query, item.Name, item.Description, item.Price, item.SupplierId, item.ID)
	return err
}

func (r *ItemRepository) Delete(id int64) error {
	query := `DELETE FROM items WHERE id = @p1`
	_, err := r.db.Exec(query, id)
	return err
}
