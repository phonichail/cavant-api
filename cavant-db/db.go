package cavantdb

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Database *sql.DB
}

func (d *DB) InitDB() {
	db, err := sql.Open("sqlite3", "./data/database/cavant.db")
	if err != nil {
		log.Fatal(err)
	}
	d.Database = db
}

func (d *DB) AddNewTable(tableName string, persistantRowId bool) error {
	statement := "CREATE TABLE IF NOT EXISTS " + "" + tableName + " ("
	if persistantRowId {
		statement += "ID INTEGER PRIMARY KEY AUTOINCREMENT,"
	}
	statement += "CREATED_DATETIME TIMESTAMP DEFAULT CURRENT_TIMESTAMP, DEACTIVATED INTEGER)" // Row is deactivated if value is not null or 0?
	_, err := d.Database.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func validateDatabaseStructureInput(input string) error {
	if len(input) == 0 {
		return errors.New("ValidateDatabaseStructureInput Error: Length of input must be more than 0.")
	}

	if input[len(input)-1] == '_' {
		return errors.New("ValidateDatabaseStructureInput Error: Input cannot end with \"_\".")
	}

	if (input[0] < 'A' || input[0] > 'Z') && (input[0] < 'a' || input[0] > 'z') {
		return errors.New("ValidateDatabaseStructureInput Error: First character of input must be an alphabetic character A-Z, a-z")
	}

	for _, value := range input {
		if (value < 'A' || value > 'Z') && (value < 'a' || value > 'z') && (value < '0' || value > '9') {
			return errors.New("ValidateDatabaseStructureInput Error: Input must only contain A-Z, a-z, 0-9")
		}
	}

	return nil
}

func (d *DB) ValidateTableNameAndAddNewTable(tableName string, persistantRowId bool) error {
	err := validateDatabaseStructureInput(tableName)
	if err != nil {
		return err
	}

	d.AddNewTable(tableName, persistantRowId)
	return nil
}

func (d *DB) AddColumnToTable(columnName string) error {
	return errors.New("Not implemented yet.")
}

func (d *DB) AddDataToTable(tableName string, columnName string, data string) error {
	transaction, err := d.Database.Begin()
	if err != nil {
		return err
	}

	statement, err := transaction.Prepare("INSERT INTO " + tableName + " DEFAULT VALUES")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetInitialTable() ([]string, error) {
	noResult := []string{}

	result, err := d.Database.Query("SELECT * FROM FishTable;")
	if err != nil {
		return noResult, err
	}

	columns, err := result.Columns()
	if err != nil {
		return noResult, err
	}
	return columns, err
}
