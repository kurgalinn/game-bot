package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type UserSets interface {
	Create(uid string, key string, values []string) (err error)
	GetRand(uid string, key string, count int) (values []string, err error)
	PopRand(uid string, key string) (value string, err error)
	Remove(uid string, key string)
}

// TODO: add key expire s.client.Expire()
type userSets struct {
	client *redis.Client
	ctx    context.Context
	expiry time.Duration
}

func NewSetsPool(rc *redis.Client, expiry time.Duration) *userSets {
	return &userSets{client: rc, ctx: context.TODO(), expiry: expiry}
}

func (s userSets) Create(uid string, key string, values []string) (err error) {
	return s.client.SAdd(s.ctx, uid+key, values).Err()
}

func (s userSets) GetRand(uid string, key string, count int) (values []string, err error) {
	return s.client.SRandMemberN(s.ctx, uid+key, int64(count)).Result()
}

func (s userSets) PopRand(uid string, key string) (value string, err error) {
	return s.client.SPop(s.ctx, uid+key).Result()
}

func (s userSets) Remove(uid string, key string) {
	s.client.Del(s.ctx, uid+key)
}
