package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accessSecret  = []byte("access_secret_Key")
	refreshSecret = []byte("refresh_secret_Key")
)

type MyClaim struct {
	UserID uint   `json:"uid"`
	Role   string `json:"role"` //区分管理员和普通用户
	Type   string `json:"type"`
	jwt.StandardClaims
}

func GenerateTokens(userID uint, role string) (accessToken string, refreshToken string, err error) {
	accessClaim := MyClaim{
		UserID: userID,
		Role:   role,
		Type:   "access",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: int64(time.Minute) * 15,
			Issuer:    "lzl",
		},
	}
	accsssTok := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)
	accessToken, err1 := accsssTok.SignedString(accessSecret)
	if err1 != nil {
		return "", "", fmt.Errorf("sign access token: %w", err1)
	}
	refreshClaim := MyClaim{
		UserID: userID,
		Role:   role,
		Type:   "refresh",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: int64(time.Hour) * 7 * 24,
			Issuer:    "lzl",
		},
	}
	refreshTok := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshToken, err2 := refreshTok.SignedString(refreshSecret)
	if err2 != nil {
		return "", "", fmt.Errorf("sign refresh token: %w", err2)
	}
	return accessToken, refreshToken, nil
}

func VerifyAccessToken(accessToken string) (claims *MyClaim, err error) {
	raw := stripBearer(accessToken)
	Token, err := jwt.ParseWithClaims(raw, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() !=
			jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v",
				t.Header["alg"])
		}
		return accessSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := Token.Claims.(*MyClaim)
	if !ok || !Token.Valid {
		return nil, errors.New("invalid access token")
	}
	if claim.Type != "access" {
		return nil, errors.New("token type mismatch: not an access token")
	}
	return claim, nil
}

func VerifyRefreshToken(refreshToken string) (claims *MyClaim, err error) {
	raw := stripBearer(refreshToken)
	Token, err := jwt.ParseWithClaims(raw, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() !=
			jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v",
				t.Header["alg"])
		}
		return refreshSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := Token.Claims.(*MyClaim)
	if !ok || !Token.Valid {
		return nil, errors.New("invalid refresh token")
	}
	if claim.Type != "refresh" {
		return nil, errors.New("token type mismatch: not an refresh token")
	}
	return claim, nil
}

func RefreshToken(refreshToken string) (string, string, error) {
	claim, err := VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	Atoken, Rtoken, err := GenerateTokens(claim.UserID, claim.Role)
	if err != nil {
		return "", "", err
	}
	return Atoken, Rtoken, nil
}

func stripBearer(token string) string {
	if len(token) > 7 && token[0:7] == "Bearer " {
		return token[7:]
	}
	return token
}
