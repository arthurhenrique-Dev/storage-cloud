package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"sort"
)

//go:embed sql/*.sql
var migrationFS embed.FS

// Run executa todas as migrations pendentes em ordem alfabética (estilo Flyway no Java)
func Run(db *sql.DB) error {
	// 1. Garante que a tabela de histórico de migrations existe
	createHistoryTableQuery := `
		IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='schema_migrations' and xtype='U')
		BEGIN
			CREATE TABLE schema_migrations (
				version NVARCHAR(255) PRIMARY KEY,
				applied_at DATETIME DEFAULT GETDATE()
			);
		END;
	`
	if _, err := db.Exec(createHistoryTableQuery); err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// 2. Lê os arquivos embutidos no binário
	entries, err := migrationFS.ReadDir("sql")
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Ordena por nome (001_..., 002_...)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		// 3. Verifica se esta migration já foi aplicada
		var count int
		err := db.QueryRow("SELECT COUNT(1) FROM schema_migrations WHERE version = @p1", filename).Scan(&count)
		if err != nil {
			return fmt.Errorf("failed to check status of migration %s: %w", filename, err)
		}

		if count > 0 {
			// Já executada anteriormente, pula
			continue
		}

		// 4. Lê o conteúdo do arquivo SQL
		content, err := migrationFS.ReadFile("sql/" + filename)
		if err != nil {
			return fmt.Errorf("failed to read sql file %s: %w", filename, err)
		}

		// 5. Executa a migration
		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		// 6. Registra no histórico que foi aplicada com sucesso
		if _, err := db.Exec("INSERT INTO schema_migrations (version) VALUES (@p1)", filename); err != nil {
			return fmt.Errorf("failed to record migration %s in history: %w", filename, err)
		}

		fmt.Printf("[MIGRATION]: Applied %s successfully\n", filename)
	}

	return nil
}
