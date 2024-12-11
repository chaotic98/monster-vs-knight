package controllers

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"monsterXknight/actions"
	"net/http"
	"strconv"
)

type Player struct {
	Interaction string `param:"interaction" query:"interaction" form:"interaction" json:"interaction" validate:"interaction"`
}

var ctx = context.Background()

func Get(c echo.Context) error {
	var player Player
	c.Bind(&player)

	client := actions.GetClient()

	monsterHealth, err := client.Get(ctx, "monster").Result()
	if errors.Is(err, redis.Nil) {
		client.Set(ctx, "monster", 100, 0)
	}

	playerHealth, err := client.Get(ctx, "player").Result()
	if errors.Is(err, redis.Nil) {
		client.Set(ctx, "player", 100, 0)
	}

	var newPlayerHealth int
	if player.Interaction == "attack" {
		dmg := actions.AttackMonster()
		decrease("monster", dmg)
		monsterHealth, _ := client.Get(ctx, "monster").Result()
		newMonsterHealth, _ := strconv.Atoi(monsterHealth)

		if newMonsterHealth <= 0 {
			client.FlushAll(ctx)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "you won",
				"data":    nil,
			})
		}
	} else {
		increase("player", 20)
		playerHealth, _ = client.Get(ctx, "player").Result()
		newPlayerHealth, _ = strconv.Atoi(playerHealth)
	}

	monsterAttackDamage := actions.AttackPlayer()
	decrease("player", monsterAttackDamage)
	playerHealth, _ = client.Get(ctx, "player").Result()
	newPlayerHealth, _ = strconv.Atoi(playerHealth)

	if newPlayerHealth <= 0 {
		client.FlushAll(ctx)
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"status":  false,
			"message": "game over",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success",
		"data": map[string]interface{}{
			"monster_health": monsterHealth,
			"player_health":  playerHealth,
		},
	})
}

func decrease(key string, dmg int) {
	client := actions.GetClient()

	val, _ := client.Get(ctx, key).Result()
	newVal, _ := strconv.Atoi(val)

	newVal -= dmg
	client.Set(ctx, key, newVal, 0)
}

func increase(key string, heal int) {
	client := actions.GetClient()

	val, _ := client.Get(ctx, key).Result()
	newVal, _ := strconv.Atoi(val)

	newVal += heal
	client.Set(ctx, key, newVal, 0)
}
