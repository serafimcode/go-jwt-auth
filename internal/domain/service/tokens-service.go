package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"jwt-auth/bootstrap"
	"jwt-auth/internal/domain/model"
	"time"
)

var jwtSigningMethod = jwt.SigningMethodHS512

type TokensService struct {
	TokenRepository model.RefreshTokenRepository
	CryptoService   CryptoService
	Env             *bootstrap.Env
}

func (s *TokensService) CreateTokens(req model.GetTokenRequest) (*model.TokensPair, error) {
	accessToken, err := s.createAccessToken(req.Guid)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createRefreshToken(req.Guid)
	if err != nil {
		return nil, err
	}

	err = s.saveRefreshToken(req.Guid, refreshToken)
	if err != nil {
		return nil, err
	}

	return &model.TokensPair{
		AccessToken: accessToken, RefreshToken: refreshToken,
	}, nil
}

func (s *TokensService) RefreshTokens(req model.TokensPair) (*model.TokensPair, error) {
	accessClaims := jwt.RegisteredClaims{}
	accessToken, err := jwt.ParseWithClaims(req.AccessToken, &accessClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Env.AccessTokenSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	refreshClaims := jwt.RegisteredClaims{}
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, &refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Env.RefreshTokenSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if !accessToken.Valid || !refreshToken.Valid {
		return nil, errors.New("invalid tokens")
	}

	refreshSub := refreshClaims.Subject
	accessSub := accessClaims.Subject
	if refreshSub != accessSub {
		return nil, errors.New("refresh token does not match access token")
	}

	tp, err := s.CreateTokens(model.GetTokenRequest{Guid: accessSub})
	if err != nil {
		return nil, err
	}

	return &model.TokensPair{AccessToken: tp.AccessToken, RefreshToken: tp.RefreshToken}, nil
}

func (s *TokensService) ExtractExpiryTime(refreshToken string) (time.Time, error) {
	refreshTokenClaims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, &refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.Env.RefreshTokenSecret), nil
	})

	if err != nil {
		return time.Time{}, err
	}

	if !token.Valid {
		return time.Time{}, fmt.Errorf("invalid Token")
	}

	return refreshTokenClaims.ExpiresAt.Time, nil
}

func (s *TokensService) saveRefreshToken(guid string, refreshToken string) error {
	hashedToken, err := s.CryptoService.HashToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to hash refresh token: %w", err)
	}

	if err != nil {
		return fmt.Errorf("token does not exist: %w", err)
	}

	err = s.TokenRepository.UpsertToken(context.Background(), guid, hashedToken)

	if err != nil {
		return fmt.Errorf("failed to persist refresh token: %w", err)
	}

	return nil
}

func (s *TokensService) createAccessToken(guid string) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(s.Env.AccessTokenExpiryHour))

	claims := jwt.RegisteredClaims{
		Subject:   guid,
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)

	at, err := token.SignedString([]byte(s.Env.AccessTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}

	return at, nil
}

func (s *TokensService) createRefreshToken(guid string) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(s.Env.RefreshTokenExpiryHour))

	claims := jwt.RegisteredClaims{
		Subject:   guid,
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)

	rt, err := token.SignedString([]byte(s.Env.RefreshTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return rt, nil
}
