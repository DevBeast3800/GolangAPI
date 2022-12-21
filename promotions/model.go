package promotions

import (
	"errors"
	"log"
	"time"

	"github.com/DevBeast3800/GolangAPI/config"
)

// Promotion represents a promotion instance
type Promotion struct {
	ID    int
	Used  bool
	Code  string
	Name  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreatePromotion creates a new promotion instance
func (u Promotion) CreatePromotion() error {
	usr := `INSERT INTO 
			promotions (used, code, name)
			VALUES ($1, $2, $3)`

	stmt, err := config.DB.Prepare(usr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(u.Used, u.Code, u.Name)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := rows.RowsAffected()

	if aff != 1 {
		err = errors.New("[ERROR - PROMOTIONS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
		return err
	}

	return nil
}

// AllPromotions returns a slice of Promotion (all promotions in promotions table)
func AllPromotions() ([]Promotion, error) {
	q := "SELECT * FROM promotions ORDER BY ID ASC"

	rows, err := config.DB.Query(q)

	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
		return nil, err
	}
	defer rows.Close()

	usrs := make([]Promotion, 0)

	for rows.Next() {
		usr := Promotion{}
		err := rows.Scan(&usr.ID, &usr.Used, &usr.Code, &usr.Name, &usr.CreatedAt, &usr.UpdatedAt)

		if err != nil {
			log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
			return nil, err
		}

		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
		return nil, err
	}
	return usrs, nil
}

// FindPromotion returns a promotion instance from the database
func (u Promotion) FindPromotion() (Promotion, error) {
	q := "SELECT * FROM promotions WHERE id = $1"

	row := config.DB.QueryRow(q, u.ID)

	usr := Promotion{}

	err := row.Scan(&usr.ID, &usr.Used, &usr.Code, &usr.Name, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
		return Promotion{}, err
	}

	return usr, nil
}

// UpdatePromotion updates the data for a promotion instance in the database
func (u Promotion) UpdatePromotion() error {
	q := "UPDATE promotions SET used=$1, code=$2, name=$3, UpdatedAt=now() WHERE ID = $4"

	stmt, err := config.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.Used, u.Code, u.Name, u.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := row.RowsAffected()
	if aff != 1 {
		err = errors.New("[ERROR - PROMOTIONS - MODEL] => More than 1 rows where affected")
	}

	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
		return err
	}
	return nil
}

// DeletePromotion deletes a promotion instance from the database
func (u Promotion) DeletePromotion() error {
	q := `DELETE FROM promotions WHERE ID=$1`

	_, err := config.DB.Exec(q, u.ID)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - MODEL] => %v", err)
		return err
	}

	return nil
}
