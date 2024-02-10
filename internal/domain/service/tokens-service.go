package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"jwt-auth/bootstrap"
	"jwt-auth/internal/domain/model"
	"time"
)

var jwtSigningMethod = jwt.SigningMethodES512

type TokensService struct {
	TokenRepository *model.RefreshTokenRepository
	Env             *bootstrap.Env
	ContextTimeout  time.Duration
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
		AccessToken: refreshToken, RefreshToken: accessToken,
	}, nil
}

func (s *TokensService) RefreshAccessToken(req model.RefreshTokenRequest) (*model.RefreshTokenResponse, error) {
	accessClaims := jwt.MapClaims{}
	accessToken, err := jwt.ParseWithClaims(req.AccessToken, accessClaims, func(token *jwt.Token) (interface{}, error) {
		return s.Env.AccessTokenSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	refreshClaims := jwt.MapClaims{}
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return s.Env.RefreshTokenSecret, nil
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

	at, err := token.SignedString(s.Env.AccessTokenSecret)
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	rt, err := token.SignedString(s.Env.RefreshTokenSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return rt, nil
}
