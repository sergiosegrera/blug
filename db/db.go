// Package db contains the DB interface that is used by the service
package db

import "github.com/sergiosegrera/blug/models"

// DB is the database manipulation interface
type DB interface {
	CreatePost(*models.Post) error
	GetPost(int) (*models.Post, error)
	DeletePost(int) error
}
