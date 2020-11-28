package redisdb

import (
	"encoding/json"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/sergiosegrera/blug/config"
	"github.com/sergiosegrera/blug/db"
	"github.com/sergiosegrera/blug/models"
)

type RedisDB struct {
	conn redis.Conn
}

func New(conf *config.Config) (db.DB, error) {
	conn, err := redis.Dial("tcp", conf.RedisAddress)
	if err != nil {
		return nil, err
	}

	return &RedisDB{
		conn: conn,
	}, err
}

func (r *RedisDB) CreatePost(post *models.Post) error {
	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	_, err = r.conn.Do("SET", "post:"+strconv.Itoa(post.Id), postJson)

	return err
}

func (r *RedisDB) GetPost(id int) (*models.Post, error) {
	var post models.Post

	postJSON, err := redis.Bytes(r.conn.Do("GET", "post:"+strconv.Itoa(id)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(postJSON, &post)
	if err != nil {
		return nil, err
	}

	return &post, err
}

func (r *RedisDB) DeletePost(id int) error {
	_, err := r.conn.Do("DEL", "post:"+strconv.Itoa(id))
	return err
}
