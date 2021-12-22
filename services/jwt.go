package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Structures
type jwtService struct {
	secretKey string
	issure    string
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

// Functions
func NewJwtService() *jwtService {

	return &jwtService{
		secretKey: GetSecretKey(),
		issure:    "user-api",
	}
}

func (s *jwtService) GenerateToken(id uint) (string, error) {

	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func GetSecretKey() string {

	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "secret"
	}

	return secret
}
