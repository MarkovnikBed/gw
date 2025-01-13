package postgres

import (
	"fmt"
	"log"
)

func (rep *Repository) Deposit(token string, amount float64, currency string) (USD, EUR, RUB float64, err error) {
	query := fmt.Sprintf("UPDATE wallets SET %s=%s+$1 WHERE token=$2 RETURNING USD,EUR,RUB", currency, currency)
	if err = rep.DB.QueryRow(query, amount, token).Scan(&USD, &EUR, &RUB); err != nil {
		log.Println(err)
		return 0, 0, 0, fmt.Errorf("ошибка внесения средств")
	}
	return
}
