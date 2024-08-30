package config

import (
	"errors"
	"os"
)

type Config struct{
	Addr string
	ConnStr string
}

func Get() (*Config, error){
	addr := os.Getenv("ADDRESS")
	if addr == ""{
		return nil, errors.New("Не найден адрес")
	}
	pgConn := os.Getenv("PG_CONN")
	if pgConn == ""{
		return nil, errors.New("Не найдет подкл стринг")
	}
	return &Config{
		Addr: addr,
		ConnStr: pgConn,
	}, nil
}