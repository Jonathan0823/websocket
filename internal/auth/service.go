package auth

import (
	"os"
	"time"
	"websocket/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Register(user models.User) error
	Login(user models.User) (string, error)
}

type authservice struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *authservice {
	return &authservice{
		repo: repo,
	}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID string) (string, error) {
	// Define claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *authservice) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func (s *authservice) Register(user models.User) error {
	return s.repo.Register(user)
}

func (s *authservice) Login(user models.User) (string, error) {
	return s.repo.Login(user)
}
