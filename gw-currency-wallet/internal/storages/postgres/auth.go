package postgres

func (r Repository) Auth(token string) (valid bool) {
	r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM wallets WHERE token=$1)", token).Scan(&valid)
	return
}
