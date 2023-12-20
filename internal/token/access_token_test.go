package token

import (
	"echo-hello/model"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGenerateToken(t *testing.T) {
	t.Parallel()
	secret := "test_secret_key"
	expirationTime := 1 * time.Hour
	user := model.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: "foo@bar.co",
		Name:  "foo_bar",
	}
	token, err := NewAccessToken(user, expirationTime, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestTokenExpire(t *testing.T) {
	t.Parallel()
	secret := "test_secret_key"
	expirationTime := time.Millisecond * 500
	user := model.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: "foo@bar.co",
		Name:  "foo_bar",
	}
	token, err := NewAccessToken(user, expirationTime, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	time.Sleep(time.Millisecond * 600)

	c := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	assert.Error(t, err, "Token should be expired")
	assert.False(t, parsedToken.Valid, "Token should be invalid due to expiration")
}
