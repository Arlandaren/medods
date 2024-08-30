package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"server/pkg/services"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)


type Claims struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
	jwt.StandardClaims
}

type TokenPair struct{
	Access string
	Refresh string
}

func GenerateAccessToken(userId, ip string) (string,error){
	expTime := time.Now().Add(time.Minute * 30)
	claims := &Claims{
		UserID: userId,
		IP: ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	access, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil{
		return "", err
	}

	return access,nil
}

func GenerateTokens(userId, ip string, s * services.Services) (*TokenPair, error){
	refreshExptime := time.Now().Add(time.Hour * 720)
	// refreshExptime := time.Now().Add(time.Second *5)
	
	access,err := GenerateAccessToken(userId,ip)
	if err != nil{
		fmt.Println("1")
		return nil, err
	}

	refreshToken := generateRefreshToken()
	refreshHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("2")
		return nil, err
	}
	err = s.AuthS.SaveRefreshToken(userId, string(refreshHash),refreshExptime)
	if err != nil {
		fmt.Println("3")
		return nil, err
	}

	return &TokenPair{
		Access:  access,
		Refresh: refreshToken,
	}, nil
}

func generateRefreshToken() string {
	randomData := make([]byte, 32)
	_, err := rand.Read(randomData)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(randomData)
}

func ValidateRefreshToken(userID, refreshToken string, s * services.Services) (bool, error) {
	refreshHash, err := s.AuthS.GetRefreshToken(refreshToken,userID)
	if err != nil {
		return false, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(refreshHash), []byte(refreshToken))
	if err != nil {
		return false, err
	}

	return true, nil
}