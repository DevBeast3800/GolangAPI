package itemsorders

import (
	"errors"
	"log"
	"time"

	"github.com/DevBeast3800/GolangAPI/config"
)

// Itemsorder represents a itemsorder instance
type Itemsorder struct {
	ID        int
	IDOrder int
	IDArticulo  string
	Price  float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateItemsorder creates a new itemsorder instance
func (u Itemsorder) CreateItemsorder() error {
	usr := `INSERT INTO 
			itemsorders (idorder, idarticulo, price)
			VALUES ($1, $2, $3)`

	stmt, err := config.DB.Prepare(usr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(u.IDOrder, u.IDArticulo, u.Price)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := rows.RowsAffected()

	if aff != 1 {
		err = errors.New("[ERROR - ITEMSORDERS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
		return err
	}

	return nil
}

// AllItemsorders returns a slice of Itemsorder (all itemsorders in itemsorders table)
func AllItemsorders() ([]Itemsorder, error) {
	q := "SELECT * FROM itemsorders ORDER BY ID ASC"

	rows, err := config.DB.Query(q)

	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
		return nil, err
	}
	defer rows.Close()

	usrs := make([]Itemsorder, 0)

	for rows.Next() {
		usr := Itemsorder{}
		err := rows.Scan(&usr.ID, &usr.IDOrder, &usr.IDArticulo, &usr.Price, &usr.CreatedAt, &usr.UpdatedAt)

		if err != nil {
			log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
			return nil, err
		}

		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
		return nil, err
	}
	return usrs, nil
}

// FindItemsorder returns a itemsorder instance from the database
func (u Itemsorder) FindItemsorder() (Itemsorder, error) {
	q := "SELECT * FROM itemsorders WHERE id = $1"

	row := config.DB.QueryRow(q, u.ID)

	usr := Itemsorder{}

	err := row.Scan(&usr.ID, &usr.IDOrder, &usr.IDArticulo, &usr.Price, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
		return Itemsorder{}, err
	}

	return usr, nil
}

// UpdateItemsorder updates the data for a itemsorder instance in the database
func (u Itemsorder) UpdateItemsorder() error {
	q := "UPDATE itemsorders SET idorder=$1, idarticulo=$2, price=$3, UpdatedAt=now() WHERE ID = $4"

	stmt, err := config.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.IDOrder, u.IDArticulo, u.Price, u.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := row.RowsAffected()
	if aff != 1 {
		err = errors.New("[ERROR - ITEMSORDERS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
		return err
	}
	return nil
}

// DeleteItemsorder deletes a itemsorder instance from the database
func (u Itemsorder) DeleteItemsorder() error {
	q := `DELETE FROM itemsorders WHERE ID=$1`

	_, err := config.DB.Exec(q, u.ID)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - MODEL] => %v", err)
		return err
	}

	return nil
}
