package actions

import (
	"context"
	"log"
	"monsterXknight/config"
	"strconv"
)

var ctx = context.Background()

func DecreaseMonsterHealth(dmg int) {
	rdb := initRedisClient()
	monster, err := rdb.Get(ctx, "monster").Result()

	newMonsterHealth, err := strconv.Atoi(monster)
	newValue := newMonsterHealth - dmg
	err = rdb.Set(ctx, "monster", newValue, 0).Err()
	if err != nil {
		log.Println("Error setting counter:", err)
	}
}

func AttackPlayer() int {

	dmgValue := generateRandBetween(config.MonsterAttackMinDmg, config.MonsterAttackMaxDmg)

	currentPlayerHealth -= dmgValue

	return dmgValue
}
