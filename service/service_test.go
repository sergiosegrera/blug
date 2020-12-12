package service

import (
	"context"
	"os"
	"testing"

	"github.com/sergiosegrera/blug/db/mockdb"
	"github.com/sergiosegrera/blug/models"
	"go.uber.org/zap/zaptest"
)

var service *BlugService

func TestMain(m *testing.M) {
	db, _ := mockdb.New(nil)
	service = &BlugService{
		DB:     db,
		Logger: zaptest.NewLogger(&testing.T{}),
	}

	os.Exit(m.Run())
}

func TestCreatePost(t *testing.T) {
	post := &models.Post{
		Id:       0,
		Title:    "Hello World",
		Html:     "<h1>Hello World</h1>",
		Markdown: "# Hello World",
	}

	err := service.CreatePost(context.Background(), post)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetPost(t *testing.T) {
	want := &models.Post{
		Id:    0,
		Title: "Hello World",
		// blackfriday adds a newline
		Html:     "<h1>Hello World</h1>\n",
		Markdown: "# Hello World",
	}

	got, err := service.GetPost(context.Background(), 0)

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if want.Id != got.Id {
		t.Errorf("Id: Got \"%v\" want \"%v\"", got.Id, want.Id)
	}

	if want.Title != got.Title {
		t.Errorf("Title: Got \"%v\" want \"%v\"", got.Title, want.Title)
	}

	if want.Html != got.Html {
		t.Errorf("Html: Got \"%v\" want \"%v\"", got.Html, want.Html)
	}

	if want.Markdown != got.Markdown {
		t.Errorf("Markdown: got \"%v\" want \"%v\"", got.Markdown, want.Markdown)
	}
}

func TestDeletePost(t *testing.T) {
	err := service.DeletePost(context.Background(), 0)

	if err != nil {
		t.Errorf(err.Error())
	}
}
