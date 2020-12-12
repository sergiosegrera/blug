package mockdb

import (
	"errors"

	"github.com/sergiosegrera/blug/config"
	"github.com/sergiosegrera/blug/db"
	"github.com/sergiosegrera/blug/models"
)

type MockDB struct {
	posts []*models.Post
}

func New(conf *config.Config) (db.DB, error) {
	return &MockDB{
		posts: []*models.Post{},
	}, nil
}

func (m *MockDB) CreatePost(post *models.Post) error {
	for i, _ := range m.posts {
		if m.posts[i].Id == post.Id {
			m.posts[i] = post
			return nil
		}
	}
	m.posts = append(m.posts, post)
	return nil
}

func (m *MockDB) GetPost(id int) (*models.Post, error) {
	for _, post := range m.posts {
		if post.Id == id {
			return post, nil
		}
	}
	return nil, errors.New("Post not found")
}

func (m *MockDB) DeletePost(id int) error {
	for i, _ := range m.posts {
		if m.posts[i].Id == id {
			m.posts[i] = m.posts[len(m.posts)-1]
			m.posts = m.posts[:len(m.posts)-1]
			return nil
		}
	}
	return errors.New("No post to delete")
}
