package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Addr string
	ConnStr string
}

func Get() (*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	addr := os.Getenv("ADDRESS")
	if addr == ""{
		return nil, errors.New("Не найден адрес")
	}
	pgConn := os.Getenv("PG_STRING")
	if pgConn == ""{
		return nil, errors.New("Не найдет подкл стринг")
	}
	return &Config{
		Addr: addr,
		ConnStr: pgConn,
	}, nil
}