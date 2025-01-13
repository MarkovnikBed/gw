package postgres

import "log"

func (rep *Repository) PrepareTable() {
	query := `CREATE TABLE IF NOT EXISTS wallets (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	username VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	token VARCHAR(255),
	usd DECIMAL(10,2) NOT NULL CHECK (usd>=0),
	eur DECIMAL(10,2) NOT NULL CHECK (eur>=0),
	rub DECIMAL(10,2) NOT NULL CHECK (rub>=0))`
	_, err := rep.DB.Exec(query)
	if err != nil {
		log.Fatalf("не удалось создать таблицу wallets: %v", err)
	}
}
