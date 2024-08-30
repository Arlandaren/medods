package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
	"server/pkg/services"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
	jwt.StandardClaims
}

type TokenPair struct{
	Access string
	Refresh string
	ExpTime time.Time
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
	access, err := token.SignedString(os.Getenv("JWT_KEY"))

	if err != nil{
		return "", err
	}

	return access,nil
}

func GenerateTokens(userId, ip string, s * services.Services) (*TokenPair, error){
	refreshExptime := time.Now().Add(time.Hour * 720)
	
	access,err := GenerateAccessToken(userId,ip)
	if err != nil{
		return nil, err
	}

	refreshToken := generateRefreshToken()
	refreshHash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = s.AuthS.SaveRefreshToken(userId, string(refreshHash),refreshExptime)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		Access:  access,
		Refresh: refreshToken,
	}, nil
}

// func GenerateNewAccess(tokens *TokenPair, userId, ip string) (*TokenPair,error){
// 	access,err := GenerateAccessToken(userId,ip)
// 	if err != nil{
// 		return nil, err
// 	}
// 	return access, nil
// }

func generateRefreshToken() string {
	randomData := fmt.Sprintf("%d", time.Now().UnixNano())
	hash := sha512.Sum512([]byte(randomData))
	return base64.URLEncoding.EncodeToString(hash[:])
}

func ValidateRefreshToken(userID, refreshToken string, s * services.Services) (bool, error) {
	refreshHash, err := s.AuthS.GetRefreshToken(userID)
	if err != nil {
		return false, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(refreshHash), []byte(refreshToken))
	if err != nil {
		return false, err
	}

	return true, nil
}