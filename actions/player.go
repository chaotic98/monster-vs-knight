package actions

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"monsterXknight/config"
	"strconv"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = config.MonsterHealth
var currentPlayerHealth = config.PlayerHealth

func InitializePlayerHealth() (int, error) {
	rdb := initRedisClient()
	player, err := rdb.Get(ctx, "player").Result()

	if errors.Is(err, redis.Nil) {
		log.Println("player is set to 100")
		err = rdb.Set(ctx, "player", 100, 0).Err()
		if err != nil {
			log.Println("error while setting up the initial player value: ", err.Error())
			return 0, err
		}
		return 0, err
	} else if err != nil {
		log.Println("Error getting player:", err)
		return 0, err
	}
	newPlayerHealth, err := strconv.Atoi(player)

	return newPlayerHealth, nil
}

func DecreasePlayerHealth(dmg int) int {
	rdb := initRedisClient()
	player, err := rdb.Get(ctx, "player").Result()

	newMonsterHealth, err := strconv.Atoi(player)
	newValue := newMonsterHealth - dmg
	err = rdb.Set(ctx, "player", newValue, 0).Err()
	if err != nil {
		log.Println("Error setting player:", err)
		return 0
	}
	return newValue
}

func IncreasePlayerHealth(heal int) int {
	rdb := initRedisClient()
	player, err := rdb.Get(ctx, "player").Result()
	fmt.Println("player health", player)
	newPlayer, err := strconv.Atoi(player)
	newValue := newPlayer + heal

	err = rdb.Set(ctx, "player", newValue, 0).Err()
	if err != nil {
		log.Println("Error setting player:", err)
		return 0
	}
	return newValue
}
func AttackMonster() int {
	minAttackValue := config.PlayerAttackMinDmg
	macAttackValue := config.PlayerAttackMaxDmg

	dmgValue := generateRandBetween(minAttackValue, macAttackValue)

	currentMonsterHealth -= dmgValue
	return dmgValue
}

func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}

func HealPlayer() int {
	healValue := generateRandBetween(config.PlayerHealMinVal, config.PlayerHealMaxVal)

	healthDiff := config.PlayerHealth - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = config.PlayerHealth
		return healthDiff
	}

}
func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}
