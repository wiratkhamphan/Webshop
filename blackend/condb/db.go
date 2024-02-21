package condb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// NewDB initializes and returns a new database connection
func NewDB() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "shoplek"
	dataSourceName := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)

	db, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		return nil, err // Return the error to be handled by the caller
	}

	// Test the database connection
	if err = db.Ping(); err != nil {
		db.Close() // Attempt to close the database connection on error
		return nil, err
	}

	// fmt.Println("Connected to MySQL database")
	return db, nil
}
