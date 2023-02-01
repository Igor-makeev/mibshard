package repository

import (
	"context"
	"fmt"
	"strconv"

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

func (rv *RedisVault) ChangeWalletBalance(ctx context.Context, key int, value int) error {

	strKey := strconv.Itoa(key)

	if err := rv.Client.Set(ctx, strKey, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (rv *RedisVault) GetWalletBalance(ctx context.Context, key int) (int, error) {
	strKey := strconv.Itoa(key)
	record, err := rv.Client.Get(ctx, strKey).Result()
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(record)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (rv *RedisVault) CreateWallet(ctx context.Context, key int, value int) error {
	strKey := strconv.Itoa(key)
	res, err := rv.Client.Exists(ctx, strKey).Result()
	if err != nil {
		return err
	}
	if res != 0 {
		err = &IdAlreadyExistError{strKey}
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
