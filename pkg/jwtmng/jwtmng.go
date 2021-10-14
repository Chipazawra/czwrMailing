package jwtmng

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Mng struct {
	signingKey string
}

func NewManager(signingKey string) (*Mng, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Mng{signingKey: signingKey}, nil
}

func (m *Mng) NewJWT(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl * time.Second).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Mng) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *Mng) ValidToken(val string) (jwt.Claims, error) {

	token, err := jwt.Parse(val, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})

	return token.Claims, err
}

func (m *Mng) ParseToken(val string) (jwt.Claims, error) {

	token, err := jwt.ParseWithClaims(val, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(m.signingKey), nil
	})

	return token.Claims, err
}
