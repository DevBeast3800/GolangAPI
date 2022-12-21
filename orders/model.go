package orders

import (
	"errors"
	"log"
	"time"

	"github.com/DevBeast3800/GolangAPI/config"
)

// Order represents a order instance
type Order struct {
	ID        int
	NROOrder int
	IDUser  int
	IDPromotion  int
	Status     string
	Total     int
	TotalDiscount     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateOrder creates a new order instance
func (u Order) CreateOrder() error {
	usr := `INSERT INTO 
			orders (nroorder, iduser, idpromotion, status, total, totaldiscount)
			VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := config.DB.Prepare(usr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(u.NROOrder, u.IDUser, u.IDPromotion, u.Status, u.Total, u.TotalDiscount)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := rows.RowsAffected()

	if aff != 1 {
		err = errors.New("[ERROR - ORDERS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
		return err
	}

	return nil
}

// AllOrders returns a slice of Order (all orders in orders table)
func AllOrders() ([]Order, error) {
	q := "SELECT * FROM orders ORDER BY ID ASC"

	rows, err := config.DB.Query(q)

	if err != nil {
		log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
		return nil, err
	}
	defer rows.Close()

	usrs := make([]Order, 0)

	for rows.Next() {
		usr := Order{}
		err := rows.Scan(&usr.ID, &usr.NROOrder, &usr.IDUser, &usr.IDPromotion, &usr.Status, &usr.Total, &usr.TotalDiscount, &usr.CreatedAt, &usr.UpdatedAt)

		if err != nil {
			log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
			return nil, err
		}

		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
		return nil, err
	}
	return usrs, nil
}

// FindOrder returns a order instance from the database
func (u Order) FindOrder() (Order, error) {
	q := "SELECT * FROM orders WHERE id = $1"

	row := config.DB.QueryRow(q, u.ID)

	usr := Order{}

	err := row.Scan(&usr.ID, &usr.NROOrder, &usr.IDUser, &usr.IDPromotion, &usr.Status, &usr.Total, &usr.TotalDiscount, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
		return Order{}, err
	}

	return usr, nil
}

// UpdateOrder updates the data for a order instance in the database
func (u Order) UpdateOrder() error {
	q := "UPDATE orders SET nroorder=$1, iduser=$2, idpromotion=$3, status=$4, total=$5, totaldiscount=$6, UpdatedAt=now() WHERE ID = $7"

	stmt, err := config.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.NROOrder, u.IDUser, u.IDPromotion, u.Status, u.Total, u.TotalDiscount, u.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := row.RowsAffected()
	if aff != 1 {
		err = errors.New("[ERROR - ORDERS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
		return err
	}
	return nil
}

// DeleteOrder deletes a order instance from the database
func (u Order) DeleteOrder() error {
	q := `DELETE FROM orders WHERE ID=$1`

	_, err := config.DB.Exec(q, u.ID)
	if err != nil {
		log.Printf("[ERROR - ORDERS - MODEL] => %v", err)
		return err
	}

	return nil
}
