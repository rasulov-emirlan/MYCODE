package products

import (
	"database/sql"
	"gumSite_01/internal/app/model"
)

func Create(p model.Product, db *sql.DB) error {
	model.ValidateProduct(&p)

	return db.QueryRow(
		"INSERT INTO gumsite_products (name, description, cost) VALUES ($1, $2, $3) RETURNING id",
		p.Name,
		p.Description,
		p.Cost,
	).Scan(&p.ID)
}

func GetALL(db *sql.DB) (*[]model.Product, error) {
	var p []model.Product
	rows, err := db.Query("SELECT id, name, description, cost FROM gumsite_products")
	if err != nil {
		return nil, err
	}
	var id int
	var name string
	var description string
	var cost int
	for rows.Next() {
		rows.Scan(&id, &name, &description, &cost)
		p = append(p, model.Product{
			ID:          id,
			Name:        name,
			Description: description,
			Cost:        cost,
		})
	}
	return &p, nil
}

func FindByName(db *sql.DB, name string) (*model.Product, error) {
	p := &model.Product{}
	if err := db.QueryRow(
		"SELECT id, name, description, cost FROM gumsite_products WHERE name=$1", name,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Cost,
	); err != nil {
		return nil, err
	}

	return p, nil
}
