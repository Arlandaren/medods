package main

import (
	"context"
	"errors"
	"fmt"
	"server/pkg/config"
	"server/pkg/handlers"
	"server/pkg/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB(cfg *config.Config, max_retries int)(*pgxpool.Pool, error){
	var retryCount int

	pool, err := pgxpool.New(context.TODO(), cfg.ConnStr)
	for err != nil && retryCount < max_retries{
		fmt.Println("емае переподключение к бд")
		time.Sleep(time.Second * 6)
		pool, err = pgxpool.New(context.TODO(), cfg.ConnStr)
		retryCount ++
	}
	
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось подключиться к бд за %s подключений", retryCount))
	}

	retryCount = 0
	err = pool.Ping(context.TODO())
	
	for err != nil && retryCount < max_retries{
		fmt.Println(err, "пытаюсь подключиться...")
		time.Sleep(time.Second * 5)
		err = pool.Ping(context.TODO())
		retryCount++
	}

	if err != nil {
		return nil, err
	}

	return pool, nil
}


func main(){
	cfg, err := config.Get()
	if err != nil{
		fmt.Println("ошибка:", err)
		return
	}

	pool, err := ConnectToDB(cfg, 5)
	
	if err != nil{
		fmt.Println("ошибка:", err)
		return
	}
	defer pool.Close()

	services, err := services.InitServices(pool)
	if err != nil{
		fmt.Println("ошибка:", err)
		return
	}

	r := gin.Default()

	handlers.InitRoutes(r, services)
	err = r.Run(cfg.Addr)
	if err != nil{
		fmt.Println("ошибка:",err)
		return
	}
	
}