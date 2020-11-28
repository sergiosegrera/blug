package service

import (
	"context"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/sergiosegrera/blug/db"
	"github.com/sergiosegrera/blug/models"
	"go.uber.org/zap"
)

type Service interface {
	// TODO: Auth
	// Login(context.Context, string) (string, error)
	CreatePost(context.Context, *models.Post) error
	GetPost(context.Context, int) (*models.Post, error)
	DeletePost(context.Context, int) error
	// TODO: Returns list of id and title?
	// GetAll(context.Context)
}

type BlugService struct {
	DB     db.DB
	Logger *zap.Logger
}

func (s *BlugService) CreatePost(ctx context.Context, post *models.Post) (err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
			"blug",
			zap.String("method", "createpost"),
			zap.Int("id", post.Id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	unsafe := blackfriday.Run(
		[]byte(post.Markdown),
		blackfriday.WithNoExtensions(),
	)
	post.Html = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	err = s.DB.CreatePost(post)

	return err
}

func (s *BlugService) GetPost(ctx context.Context, id int) (post *models.Post, err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
			"blug",
			zap.String("method", "getpost"),
			zap.Int("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	post, err = s.DB.GetPost(id)

	return post, err
}

func (s *BlugService) DeletePost(ctx context.Context, id int) (err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
			"blug",
			zap.String("method", "deletepost"),
			zap.Int("id", id),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = s.DB.DeletePost(id)

	return err
}
