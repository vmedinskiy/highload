package jwt

import (
	"time"

	"github.com/go-faster/errors"
	"github.com/golang-jwt/jwt/v4"

	"github.com/vmedinskiy/highload/internal/pkg/user"
)

// Manager is a JSON web token manager
type Manager struct {
	secretKey     string
	tokenDuration time.Duration
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"userID"`
}

// NewManager returns a new JWT manager
func NewManager(secretKey string, tokenDuration time.Duration) *Manager {
	return &Manager{secretKey, tokenDuration}
}

func (manager *Manager) Generate(user *user.User) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(manager.secretKey))
	if err != nil {
		return "", errors.Wrap(err, "generate jwt token")
	}
	return ss, nil
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *Manager) Verify(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return -1, errors.Wrap(err, "invalid token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return -1, errors.New("invalid token claims")
	}

	return claims.UserID, nil
}
