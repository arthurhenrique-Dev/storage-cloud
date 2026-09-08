package connection

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"storage-api/shared/operations"

	_ "github.com/microsoft/go-mssqldb"
)

func Connect(connStr string) (*sql.DB, error) {
	var db *sql.DB

	err := operations.
		NewPipeline("SQL Server Connection").
		Step("Ensure Database Exists", func() error {
			return ensureDatabaseExists(connStr)
		}).
		Step("Open SQL Server", func() error {
			var err error
			db, err = sql.Open("sqlserver", connStr)
			return err
		}).
		Step("Configure Connection Pool", func() error {
			db.SetMaxOpenConns(20)
			db.SetMaxIdleConns(10)
			db.SetConnMaxLifetime(30 * time.Minute)
			db.SetConnMaxIdleTime(5 * time.Minute)
			return nil
		}).
		Step("Check SQL Server", func() error {
			return db.Ping()
		}).Execute()

	if err != nil {
		if db != nil {
			db.Close()
		}
		return nil, err
	}
	return db, nil
}

// ensureDatabaseExists conecta temporariamente no banco 'master' e cria o database alvo caso ainda não exista
func ensureDatabaseExists(connStr string) error {
	parts := strings.Split(connStr, ";")
	var dbName string
	var masterParts []string

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if strings.HasPrefix(strings.ToLower(trimmed), "database=") {
			dbName = trimmed[len("database="):]
			masterParts = append(masterParts, "database=master")
		} else if trimmed != "" {
			masterParts = append(masterParts, trimmed)
		}
	}

	if dbName == "" || strings.EqualFold(dbName, "master") {
		return nil
	}

	masterConnStr := strings.Join(masterParts, ";")
	masterDB, err := sql.Open("sqlserver", masterConnStr)
	if err != nil {
		return err
	}
	defer masterDB.Close()

	query := fmt.Sprintf("IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = '%s') CREATE DATABASE [%s];", dbName, dbName)
	_, err = masterDB.Exec(query)
	return err
}
