package query

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/osag1e/gtc/internal/model"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestProductStore(t *testing.T) {
	container, sqlDB, err := SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer container.Terminate(context.Background())
	// Ensuring container is stopped after the test

	bookStore := &BookStore{
		DB: sqlDB,
	}

	book := model.Books{
		Title:  "gtc",
		Author: "Osagie",
		Price:  20.29,
	}

	insertedBook, err := bookStore.InsertBook(&book)
	if err != nil {
		t.Fatalf("InsertBook returned an unexpected error: %v", err)
	}

	if insertedBook.ID == uuid.Nil {
		t.Errorf("Expected book ID to be set, but it was empty")
	}
}

func SetupTestDatabase() (testcontainers.Container, *sql.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, nil, err
	}

	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, nil, err
	}

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		return nil, nil, err
	}

	connStr := fmt.Sprintf("user=postgres password=postgres dbname=testdb host=%s port=%s sslmode=disable", host, port.Port())

	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	migrationsDir := filepath.Join(currentDir, "db_test_script")

	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	n, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	log.Printf("Applied %d migrations!\n", n)

	return dbContainer, sqlDB, nil
}
