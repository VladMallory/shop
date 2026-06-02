package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// InitDB подключается к СУБД, проверяет связь и создает таблицы, если их нет.
func InitDB(dsn string) (*sql.DB, error) {
	// 1. Открываем соединение (драйвер "postgres")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть соединение с БД: %w", err)
	}

	// Настраиваем пул соединений (хорошая практика для продакшена)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 2. Проверяем, что база действительно доступна (выполняем пинг)
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("база данных недоступна (Ping failed): %w", err)
	}

	// 3. Запускаем создание таблиц
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("ошибка при создании структуры таблиц: %w", err)
	}

	return db, nil
}

// createTables содержит SQL-скрипты для инициализации таблиц.
func createTables(db *sql.DB) error {
	// Пишем SQL-запрос с использованием IF NOT EXISTS
	// Здесь мы создаем, например, таблицу пользователей
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`

	// Выполняем запрос в БД
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу users: %w", err)
	}

	// Если нужно создать еще таблицы, можно выполнить еще один db.Exec() ниже
	return nil
}
