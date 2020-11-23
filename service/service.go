package service

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	Login(context.Context, string) (string, error)
	PostPost(context.Context, string) error
	GetPost(context.Context, int) (string, error)
	DeletePost(context.Context, int) error
	// TODO: Returns list of id and title?
	// GetAll(context.Context)
}

type BlugService struct {
	DB     db.DB
	Logger *zap.Logger
}
