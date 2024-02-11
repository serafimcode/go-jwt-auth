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

func (s *TokensService) SaveRefreshToken(guid string, refreshToken string) error {
	hashedToken, err := s.CryptoService.HashToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to hash refresh token: %w", err)
	}

	err = s.TokenRepository.Create(context.Background(), guid, hashedToken)

	if err != nil {
		return fmt.Errorf("failed to persist refresh token: %w", err)
	}

	return nil
}

func (s *TokensService) RetrieveToken(refreshToken string) (*model.TokenEntity, error) {
	guid, err := s.extractGuidFromToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to extract guid from refresh token: %w", err)
	}

	te, err := s.TokenRepository.GetByGuid(context.Background(), guid)
	if err != nil {
		return nil, fmt.Errorf("token does not exist: %w", err)
	}

	return te, nil
}

func (s *TokensService) CreateTokens(req model.GetTokenRequest) (*model.GetTokensResponse, error) {
	accessToken, err := s.createAccessToken(req.Guid)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createRefreshToken(req.Guid)
	if err != nil {
		return nil, err
	}

	return &model.GetTokensResponse{
		AccessToken: accessToken, RefreshToken: refreshToken,
	}, nil
}

func (s *TokensService) RefreshAccessToken(req model.RefreshTokenRequest) (*model.RefreshTokenResponse, error) {
	accessClaims := jwt.MapClaims{}
	accessToken, err := jwt.ParseWithClaims(req.AccessToken, &accessClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Env.AccessTokenSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	refreshClaims := jwt.MapClaims{}
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, &refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Env.RefreshTokenSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if !accessToken.Valid || !refreshToken.Valid {
		return nil, errors.New("invalid tokens")
	}

	refreshSub := refreshClaims["sub"]
	accessSub := accessClaims["sub"]
	if refreshSub != accessSub {
		return nil, errors.New("refresh token does not match access token")
	}

	at, err := s.createAccessToken(accessSub.(string))
	if err != nil {
		return nil, err
	}

	return &model.RefreshTokenResponse{AccessToken: at}, nil
}

func (s *TokensService) createAccessToken(guid string) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(s.Env.AccessTokenExpiryHour))

	claims := jwt.MapClaims{
		"sub": guid,
		"exp": exp,
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

	claims := jwt.MapClaims{
		"sub": guid,
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)

	rt, err := token.SignedString([]byte(s.Env.RefreshTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return rt, nil
}

func (s *TokensService) extractGuidFromToken(refreshToken string) (string, error) {
	refreshClaims := jwt.MapClaims{}

	parsedToken, err := jwt.ParseWithClaims(refreshToken, &refreshClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.Env.RefreshTokenSecret), nil
	})

	if err != nil {
		fmt.Println("ERROR HERE")
		return "", err
	}

	if !parsedToken.Valid {
		return "", fmt.Errorf("token invalid")
	}

	return refreshClaims["sub"].(string), nil
}
