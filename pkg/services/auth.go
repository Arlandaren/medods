package services

import (
	"context"
	"time"

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

func (s *AuthService) init() error {
	query := `CREATE TABLE IF NOT EXISTS tokens (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR NOT NULL,
		token VARCHAR NOT NULL UNIQUE,
		exp_time TIMESTAMP
	);`

	return s.pool.AcquireFunc(context.Background(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
}

func (s *AuthService) SaveRefreshToken(userID string, token string, expTime time.Time) error {
	query := `INSERT INTO tokens (user_id, token, exp_time) VALUES ($1,$2,$3)`
	return s.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, userID, token, expTime)
		return err
	})
}

func (s *AuthService) GetRefreshToken(tokenHash,userID string) (string, error) {
	query := `SELECT token FROM tokens WHERE user_id = $1 AND exp_time > $2 AND token = $3 ORDER BY exp_time DESC LIMIT 1`

	var token string
	now := time.Now()

	err := s.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		err := c.QueryRow(context.TODO(), query, userID, now, tokenHash).Scan(&token)
		return err
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

