package redis

import (
	"bws_test/app/models"
	"fmt"
	"github.com/go-redis/redis"
)

type HackersRepo struct {
	client *redis.Client
	key    string
}

func CreateHackersRepo(client *redis.Client, key string) *HackersRepo {
	return &HackersRepo{client: client, key: key}
}

func (r *HackersRepo) Create(hacker models.Hacker) error {
	res := r.client.ZAdd(r.key, redis.Z{Score: float64(hacker.Score), Member: hacker.Name})
	if res.Err() != nil {
		return fmt.Errorf("redisHackerRepo Create() error: %w", res.Err())
	}
	return nil
}

func (r *HackersRepo) GetAll() ([]models.Hacker, error) {
	redisResult, err := r.client.ZRangeWithScores(r.key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("redisHackerRepo GetAll() error: %w", err)
	}
	result := make([]models.Hacker, len(redisResult))
	var ok bool
	for i, item := range redisResult {
		result[i].Score = int(item.Score)
		result[i].Name, ok = item.Member.(string)
		if !ok {
			return nil, fmt.Errorf("redisHackerRepo GetAll() error: member converting to string error %v", item)
		}
	}
	return result, nil
}
