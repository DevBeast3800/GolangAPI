package items

import (
	"errors"
	"log"
	"time"

	"github.com/DevBeast3800/GolangAPI/config"
)

// Item represents a item instance
type Item struct {
	ID        int
	Available bool
	Price  float64
	Name  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateItem creates a new item instance
func (u Item) CreateItem() error {
	usr := `INSERT INTO 
			items (available, price, name)
			VALUES ($1, $2, $3)`

	stmt, err := config.DB.Prepare(usr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(u.Available, u.Price, u.Name)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := rows.RowsAffected()

	if aff != 1 {
		err = errors.New("[ERROR - ITEMS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
		return err
	}

	return nil
}

// AllItems returns a slice of ITEM (all items in items table)
func AllItems() ([]Item, error) {
	q := "SELECT * FROM items ORDER BY ID ASC"

	rows, err := config.DB.Query(q)

	if err != nil {
		log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
		return nil, err
	}
	defer rows.Close()

	usrs := make([]Item, 0)

	for rows.Next() {
		usr := Item{}
		err := rows.Scan(&usr.ID, &usr.Available, &usr.Price, &usr.Name, &usr.CreatedAt, &usr.UpdatedAt)

		if err != nil {
			log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
			return nil, err
		}

		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
		return nil, err
	}
	return usrs, nil
}

// FindItem returns a item instance from the database
func (u Item) FindItem() (Item, error) {
	q := "SELECT * FROM items WHERE id = $1"

	row := config.DB.QueryRow(q, u.ID)

	usr := Item{}

	err := row.Scan(&usr.ID, &usr.Available, &usr.Price, &usr.Name, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
		return Item{}, err
	}

	return usr, nil
}

// UpdateItem updates the data for a item instance in the database
func (u Item) UpdateItem() error {
	q := "UPDATE items SET available=$1, price=$2, name=$3, UpdatedAt=now() WHERE ID = $4"

	stmt, err := config.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.Available, u.Price, u.Name, u.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := row.RowsAffected()
	if aff != 1 {
		err = errors.New("[ERROR - ITEMS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
		return err
	}
	return nil
}

// DeleteItem deletes a item instance from the database
func (u Item) DeleteItem() error {
	q := `DELETE FROM items WHERE ID=$1`

	_, err := config.DB.Exec(q, u.ID)
	if err != nil {
		log.Printf("[ERROR - ITEMS - MODEL] => %v", err)
		return err
	}

	return nil
}
