package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"monsterXknight/actions"
	"net/http"
)

type Player struct {
	Interaction string `param:"interaction" query:"interaction" form:"interaction" json:"interaction" validate:"interaction"`
}

func Get(c echo.Context) error {
	var player Player
	c.Bind(&player)

	monsterHealth, err := actions.Get("monster")
	if errors.Is(err, redis.Nil) {
		actions.Set("monster", 100)
	}

	playerHealth, err := actions.Get("player")
	if errors.Is(err, redis.Nil) {
		actions.Set("player", 100)
	}

	if player.Interaction == "attack" {
		dmg := actions.AttackMonster()
		actions.DecreaseMonsterHealth(dmg)
		monsterHealth, _ = actions.Get("monster")

		if monsterHealth <= 0 {
			actions.Reset()
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "you won",
				"data":    nil,
			})
		}
	} else {
		heal := 20
		playerHealth = actions.IncreasePlayerHealth(heal)
	}
	monsterAttackDamage := actions.AttackPlayer()

	playerHealth = actions.DecreasePlayerHealth(monsterAttackDamage)

	if playerHealth <= 0 {
		actions.Reset()
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
