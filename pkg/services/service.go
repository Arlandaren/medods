package services

import "github.com/jackc/pgx/v5/pgxpool"

type Services struct {
	AuthS *AuthService
}

func InitServices(pool *pgxpool.Pool) (*Services,error){
	as,err := NewAuthService(pool)
	if err != nil{
		return nil,err
	}

	return &Services{AuthS: as}, nil
}