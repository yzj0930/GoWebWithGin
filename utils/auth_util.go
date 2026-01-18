package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager struct {
    secretKey     string
    tokenDuration time.Duration
}

type UserClaims struct {
    jwt.RegisteredClaims
    UserID    uint   `json:"user_id"`
    Username  string `json:"username"`
    Email     string `json:"email,omitempty"`
    Role      string `json:"role"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
    return &JWTManager{
        secretKey:     secretKey,
        tokenDuration: tokenDuration,
    }
}

// 生成Token
func (manager *JWTManager) GenerateToken(userID uint, username, email, role string) (string, error) {
    claims := UserClaims{
        RegisteredClaims: jwt.RegisteredClaims{
            ID:        uuid.New().String(),
            Subject:   username,
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.tokenDuration)),
            Issuer:    "your-app-name",
        },
        UserID:   userID,
        Username: username,
        Email:    email,
        Role:     role,
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(manager.secretKey))
}

// 验证Token
func (manager *JWTManager) VerifyToken(tokenString string) (*UserClaims, error) {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &UserClaims{},
        func(token *jwt.Token) (interface{}, error) {
            _, ok := token.Method.(*jwt.SigningMethodHMAC)
            if !ok {
                return nil, errors.New("unexpected signing method")
            }
            return []byte(manager.secretKey), nil
        },
    )
    
    if err != nil {
        return nil, err
    }
    
    claims, ok := token.Claims.(*UserClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }
    
    return claims, nil
}

// 刷新Token
func (manager *JWTManager) RefreshToken(tokenString string) (string, error) {
    claims, err := manager.VerifyToken(tokenString)
    if err != nil {
        return "", err
    }
    
    // 如果Token即将过期（剩余时间小于一半），允许刷新
    if time.Until(claims.ExpiresAt.Time) > manager.tokenDuration/2 {
        return tokenString, nil // 无需刷新
    }
    
    return manager.GenerateToken(
        claims.UserID,
        claims.Username,
        claims.Email,
        claims.Role,
    )
}