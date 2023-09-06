package auth

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey []byte
}

func NewService(secretKeyFile string) (*jwtService, error) {
	keyBytes, err := ioutil.ReadFile(secretKeyFile)
	if err != nil {
		return nil, err
	}

	secretKey := strings.TrimSpace(string(keyBytes))

	return &jwtService{secretKey: []byte(secretKey)}, nil
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claimedUser := jwt.MapClaims{}
	claimedUser["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimedUser)

	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return s.secretKey, nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
