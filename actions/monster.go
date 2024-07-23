package actions

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"monsterXknight/config"
	"strconv"
)

var ctx = context.Background()

func InitializeMonsterHealth() (int, error) {

	rdb := initRedisClient()
	monster, err := rdb.Get(ctx, "monster").Result()

	//rdb.FlushAll(ctx)
	//os.Exit(1)
	if errors.Is(err, redis.Nil) {
		log.Println("monster is set to 100")
		err = rdb.Set(ctx, "monster", 100, 0).Err()
		if err != nil {
			log.Println("error while setting up the initial monster value: ", err.Error())
			return 0, err
		}
		return 0, err
	} else if err != nil {
		log.Println("Error getting counter:", err)
		return 0, err
	}
	newMonsterHealth, err := strconv.Atoi(monster)

	return newMonsterHealth, nil
}

func DecreaseMonsterHealth(dmg int) int {
	rdb := initRedisClient()
	monster, err := rdb.Get(ctx, "monster").Result()

	newMonsterHealth, err := strconv.Atoi(monster)
	newValue := newMonsterHealth - dmg
	err = rdb.Set(ctx, "monster", newValue, 0).Err()
	if err != nil {
		log.Println("Error setting counter:", err)
		return 0
	}
	return newValue
}

func AttackPlayer() int {

	dmgValue := generateRandBetween(config.MonsterAttackMinDmg, config.MonsterAttackMaxDmg)

	currentPlayerHealth -= dmgValue

	return dmgValue
}
