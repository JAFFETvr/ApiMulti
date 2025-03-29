package database

import "database/sql"

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{DB: db}
}

func (repo *MySQLUserRepository) SaveRequest(username string, password string) error {
	_, err := repo.DB.Exec("INSERT INTO registration_requests (username, password, status) VALUES (?, ?, 'pending')", username, password)
	return err
}

func (repo *MySQLUserRepository) GetPendingRequests() ([]map[string]interface{}, error) {
	rows, err := repo.DB.Query("SELECT id, username, password, status FROM registration_requests WHERE status = 'pending'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []map[string]interface{}
	for rows.Next() {
		var id int32
		var username, password, status string
		if err := rows.Scan(&id, &username, &password, &status); err != nil {
			return nil, err
		}
		request := map[string]interface{}{
			"id":       id,
			"username": username,
			"password": password,
			"status":   status,
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (repo *MySQLUserRepository) ApproveUser(id int, macAddress string) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	var username, password string
	err = tx.QueryRow("SELECT username, password FROM registration_requests WHERE id = ?", id).Scan(&username, &password)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO users (username, password, mac_address, role) VALUES (?, ?, ?, 'user')", username, password, macAddress)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE registration_requests SET status = 'approved' WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (repo *MySQLUserRepository) RejectUser(id int) error {
	_, err := repo.DB.Exec("UPDATE registration_requests SET status = 'rejected' WHERE id = ?", id)
	return err
}
