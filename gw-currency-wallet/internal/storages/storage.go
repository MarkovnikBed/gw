package storages

import (
	"log"
	"os"

	"wallet/internal/storages/postgres"
)

type Storage interface {
	PrepareTable()
	Registration(username string, email string, password string) (err error)
	Login(username string, password string) (token string, err error)
	Auth(token string) (valid bool)
	Deposit(token string, amount float64, currency string) (usd, eur, rub float64, err error)
	Withdraw(token string, amount float64, currency string) (usd, eur, rub float64, err error)
	Exchange(token string, from string, to string, rate, amount float64) (usd, eur, rub float64, err error)
}

func CreateRepository() Storage {
	switch storage := os.Getenv("STORAGE"); storage {
	case "postgres":
		return &postgres.Repository{DB: postgres.ConnectDB()}
	case "sqlite":
		log.Fatal("в данной версии программы не поддерживается запрашиваемая вами СУБД(sqlite), ждите будущих обновлений")
	case "":
		log.Fatal("в конфигурационном файле нет данных о переменной STORAGE")
	default:
		log.Fatalf("указана неподдерживаемая СУБД - %v", storage)
	}

	return nil
}
