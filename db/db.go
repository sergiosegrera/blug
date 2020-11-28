package db

import "github.com/sergiosegrera/blug/models"

type DB interface {
	CreatePost(*models.Post) error
	GetPost(int) (*models.Post, error)
	DeletePost(int) error
}
