package database

import (
	"context"
	"log"
	"os"
)

func RunMigrations() {
	sqlFile, err := os.ReadFile("database/create_tables.sql")
	if err != nil {
		log.Fatalf("Ошибка чтения файла миграции: %v", err)
	}

	_, err = DB.Exec(context.Background(), string(sqlFile))
	if err != nil {
		log.Fatalf("Ошибка выполнения миграции: %v", err)
	}

	log.Println("Миграции успешно выполнены")
}
