package postgres

import (
	"fmt"
	"log"
)

func (rep *Repository) Registration(username string, email string, password string) (err error) {
	var exists bool

	tx, _ := rep.DB.Begin()

	tx.QueryRow("SELECT EXISTS(SELECT 1 FROM wallets WHERE email=$1 OR username=$2)", email, username).Scan(&exists)
	if exists {
		tx.Commit()
		return fmt.Errorf("пользователь с таким логином или именем уже существует")
	}

	if _, err = tx.Exec("INSERT INTO wallets (username, email,password,usd, eur, rub) VALUES ($1,$2,$3,0,0,0)", username, email, password); err != nil {
		log.Println(err)
		tx.Rollback()
		return fmt.Errorf("не удалось выполнить регистрацию")
	}
	tx.Commit()

	return nil
}
