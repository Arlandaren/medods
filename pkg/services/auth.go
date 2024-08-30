package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct {
	pool *pgxpool.Pool
}

func NewAuthService(pool *pgxpool.Pool) (*AuthService, error){
	AuthS := &AuthService{pool: pool}
	if err := AuthS.init(); err != nil{
		return nil,err
	}
	return AuthS,nil
}

func (s *AuthService) init() error{
	query := ""

	return s.pool.AcquireFunc(context.Background(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
}

func (as *AuthService) CreateUser(){
	
}

