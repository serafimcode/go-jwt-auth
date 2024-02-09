package service

import (
	"jwt-auth/bootstrap"
	"jwt-auth/internal/domain/model"
	"time"
)

type TokensService struct {
	TokenRepository *model.RefreshTokenRepository
	Env             *bootstrap.Env
	ContextTimeout  time.Duration
}
