// Package mockdb is a non-persistent test db.DB implementation
package mockdb

import (
	"errors"

	"github.com/sergiosegrera/blug/config"
	"github.com/sergiosegrera/blug/db"
	"github.com/sergiosegrera/blug/models"
)

// MockDB implements db.DB
type MockDB struct {
	posts []*models.Post
}

// New creates a new instance of MockDB
func New(conf *config.Config) (db.DB, error) {
	return &MockDB{
		posts: []*models.Post{},
	}, nil
}

// CreatePost inserts a post into the db
func (m *MockDB) CreatePost(post *models.Post) error {
	for i := range m.posts {
		if m.posts[i].ID == post.ID {
			m.posts[i] = post
			return nil
		}
	}
	m.posts = append(m.posts, post)
	return nil
}

// GetPost requests a post from the db
func (m *MockDB) GetPost(id int) (*models.Post, error) {
	for _, post := range m.posts {
		if post.ID == id {
			return post, nil
		}
	}
	return nil, errors.New("Post not found")
}

// DeletePost delets a post from the db
func (m *MockDB) DeletePost(id int) error {
	for i := range m.posts {
		if m.posts[i].ID == id {
			m.posts[i] = m.posts[len(m.posts)-1]
			m.posts = m.posts[:len(m.posts)-1]
			return nil
		}
	}
	return errors.New("No post to delete")
}
