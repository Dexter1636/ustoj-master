package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	//"github.com/golang-jwt/jwt"
)

//JWTService is a cotnract of what jwtService can do
type JWTService interface {
	GenerateToken(username string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtCustomClain struct {
	Username string `gorm:"varchar(20)"`
	jwt.StandardClaims
}
type jwtService struct {
	issuer    string
	secretKey string
}

//NewJWTService method creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}
func getSecretKey() string {
	//serectKey := os.Getenv("JWT_SECRET")
	serectKey := "ydhnwbsecret"
	if serectKey != "" {
		serectKey = "ydhnwb"
	}
	return serectKey

}

func (j *jwtService) GenerateToken(Username string) string {
	claims := &jwtCustomClain{
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected  signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
