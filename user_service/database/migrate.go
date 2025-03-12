package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunMigrations(db *pgxpool.Pool) error {
	sqlFile, err := os.ReadFile("database/create_tables.sql")
	if err != nil {
		log.Printf("Ошибка чтения файла миграции: %v", err)
		return err
	}

	_, err = db.Exec(context.Background(), string(sqlFile))
	if err != nil {
		log.Printf("Ошибка выполнения миграции: %v", err)
		return err
	}

	log.Print("Миграции успешно выполнены")
	return nil
}
