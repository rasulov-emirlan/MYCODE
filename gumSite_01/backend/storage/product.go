package storage

// Product describes a row of gumsite_roducts from postgresql
type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Image       string `json:"image"`
}

// SelectAllProducts returns a pointer to an array of type Product
// which is a copy of all rows from gumsite_products
func SelectAllProducts(store Storage) (*[]Product, error) {
	p := []Product{}
	var id int
	var name string
	var desc string
	var cost int
	var image string
	rows, err := store.db.Query("SELECT id, name, description, cost, image FROM gumsite_products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &name, &desc, &cost, &image)
		p = append(p, Product{
			ID:          id,
			Name:        name,
			Description: desc,
			Cost:        cost,
			Image:       image,
		})
	}
	return &p, nil
}

// FindByName selects a row from gumsite_products
// where column name=(second argument of FindByName of type string)
func FindProductByName(store Storage, name string) (*Product, error) {
	row, err := store.db.Query("SELECT id, name, description, cost, image FROM gumsite_products WHERE name=$1", name)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		p := &Product{}
		row.Scan(&p.ID, &p.Name, &p.Description, &p.Cost, &p.Image)
		return p, nil
	}
	return nil, nil
}

// InsertProduct inserts a row that is descibed by an argument of type Product
// into gumsite_products
func InserProduct(store Storage, p Product) error {
	if _, err := store.db.Query("INSERT INTO gumsite_products(name, description, cost, image) VALUES($1,$2,$3, $4)", p.Name, p.Description, p.Cost, p.Image); err != nil {
		return err
	}
	return nil
}

// DeleteProduct deletes a row from gumsite_products
// where a name=(second argument from DelectProduct of type string)
func DeleteProduct(store Storage, name string) error {
	if _, err := store.db.Query("DELETE FROM gumsite_products WHERE name=$1", name); err != nil {
		return err
	}
	return nil
}
