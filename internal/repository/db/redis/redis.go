package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisVault struct {
	Client *redis.Client
}

func NewRedisVault(address string) *RedisVault {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	return &RedisVault{Client: client}
}

func (rv *RedisVault) SetNote(ctx context.Context, key string, value int) error {
	if err := rv.Client.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (rv *RedisVault) GetNote(ctx context.Context, key string) (string, error) {
	result, err := rv.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func (rv *RedisVault) CreateNote(ctx context.Context, key string, value int) error {
	res, err := rv.Client.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if res != 0 {
		err = &IdAlreadyExistError{key}
		return err
	}

	return nil
}

type IdAlreadyExistError struct {
	id string
}

func (iaee *IdAlreadyExistError) Error() string {
	return fmt.Sprintf("error: Id:%v has already exist", iaee.id)
}
