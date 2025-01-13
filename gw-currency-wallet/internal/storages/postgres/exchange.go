package postgres

import (
	"fmt"
	"log"
)

func (rep *Repository) Exchange(token, from, to string, rate float64, amount float64) (usd, eur, rub float64, err error) {
	tx, _ := rep.DB.Begin()
	query := fmt.Sprintf(`UPDATE wallets SET %s=%s-$1 WHERE token=$2`, from, from)
	_, err = tx.Exec(query, amount, token)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, 0, 0, fmt.Errorf("недостаточно средств")
	}
	toAmount := amount * rate
	query = fmt.Sprintf(`UPDATE wallets SET %s=%s+$1 WHERE token=$2 RETURNING USD,EUR,RUB`, to, to)
	err = tx.QueryRow(query, toAmount, token).Scan(&usd, &eur, &rub)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, 0, 0, fmt.Errorf("не удалось внести средства на %s счет - отмена операции", to)
	}

	tx.Commit()

	err = nil

	return
}
