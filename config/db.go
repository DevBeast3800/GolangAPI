package config

import (
	"database/sql"
	"fmt"
	"log"

	// db driver
	_ "github.com/lib/pq"
)

// DB variable of type pointer to DB type in sql package
var DB *sql.DB

func init() {

	var err error

	const (
		host     = "api-candidates.cljbtn7puujc.us-east-1.rds.amazonaws.com"
		port     = 5432
		user     = "amarcano"
		password = "amarcanotest"
		dbname   = "amarcano_db"
	)
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		log.Printf("[warning] => %v", err)
	} else {
		fmt.Println("Database ready to GO!")
	}

	usersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id          SERIAL PRIMARY KEY NOT NULL,
			name        VARCHAR(255) NOT NULL,
			active 	    BOOLEAN NOT NULL,
			type	    VARCHAR(255) NOT NULL,
			createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
			updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`
	_, err = DB.Exec(usersTable)

	itemsTable := `
		CREATE TABLE IF NOT EXISTS items (
			id          SERIAL PRIMARY KEY NOT NULL,
			available   VARCHAR(255) NOT NULL,
			price 	    FLOAT NOT NULL,
			name	    VARCHAR(255) NOT NULL,
			createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
			updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`
	_, err = DB.Exec(itemsTable)

	promotionsTable := `
		CREATE TABLE IF NOT EXISTS promotions (
			id          SERIAL PRIMARY KEY NOT NULL,
			used        BOOLEAN NOT NULL,
			code 	    VARCHAR(255) NOT NULL,
			name	    VARCHAR(255) NOT NULL,
			createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
			updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`
	_, err = DB.Exec(promotionsTable)

	ordersTable := `
		CREATE TABLE IF NOT EXISTS orders (
			id          SERIAL PRIMARY KEY NOT NULL,
			nroorder    INT NOT NULL,
			iduser 	    INT NOT NULL,
			idpromotion INT NOT NULL,
			status 		VARCHAR(255) NOT NULL,
			total 		INT NOT NULL,
			totaldiscount INT NOT NULL,
			createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
			updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`
	_, err = DB.Exec(ordersTable)

	itemsordersTable := `
		CREATE TABLE IF NOT EXISTS itemsorders (
			id          SERIAL PRIMARY KEY NOT NULL,
			idorder     INT NOT NULL,
			idarticulo 	VARCHAR(255) NOT NULL,
			price	    FLOAT NOT NULL,
			createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
			updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`
	_, err = DB.Exec(itemsordersTable)
	
	if err != nil {
		log.Printf("[WARNING - CONFIG - DB] => %v", err)
	}

}
