package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository() (*MySQLUserRepository, error) {
	// Cargar variables de entorno
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SCHEMA"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la base de datos: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	return &MySQLUserRepository{DB: db}, nil
}


func (repo *MySQLUserRepository) SaveRequest(username string, passwordHash string) error {
	
	_, err := repo.DB.Exec("INSERT INTO registration_requests (username, password_hash, status) VALUES (?, ?, 'pending')", username, passwordHash)
	if err != nil {
		return fmt.Errorf("error al guardar solicitud de registro: %w", err)
	}
	return nil
}

func (repo *MySQLUserRepository) GetPendingRequests() ([]map[string]interface{}, error) {
	rows, err := repo.DB.Query("SELECT id, username, password_hash, status FROM registration_requests WHERE status = 'pending'")
	if err != nil {
		return nil, fmt.Errorf("error al obtener solicitudes pendientes: %w", err)
	}
	defer rows.Close()

	var requests []map[string]interface{}
	for rows.Next() {
		var id int32
		var username, passwordHash, status string
		if err := rows.Scan(&id, &username, &passwordHash, &status); err != nil {
			return nil, fmt.Errorf("error al escanear solicitudes pendientes: %w", err)
		}
		request := map[string]interface{}{
			"id":            id,
			"username":      username,
			"password_hash": passwordHash,
			"status":        status,
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (repo *MySQLUserRepository) ApproveUser(id int, macAddress string) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar transacción: %w", err)
	}

	var username, passwordHash string // Expect password_hash
	err = tx.QueryRow("SELECT username, password_hash FROM registration_requests WHERE id = ?", id).Scan(&username, &passwordHash)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error al obtener datos del usuario: %w", err)
	}

	_, err = tx.Exec("INSERT INTO users (username, password_hash, mac_address, role) VALUES (?, ?, ?, 'user')", username, passwordHash, macAddress)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error al insertar usuario aprobado: %w", err)
	}

	_, err = tx.Exec("UPDATE registration_requests SET status = 'approved' WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error al actualizar estado de solicitud: %w", err)
	}

	return tx.Commit()
}

func (repo *MySQLUserRepository) RejectUser(id int) error {
	_, err := repo.DB.Exec("UPDATE registration_requests SET status = 'rejected' WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error al rechazar solicitud de usuario: %w", err)
	}
	return nil
}