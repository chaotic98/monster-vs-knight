package controllers

import (
	"github.com/labstack/echo/v4"
	"log"
	"monsterXknight/actions"
	"net/http"
)

type Player struct {
	Interaction string `param:"interaction" query:"interaction" form:"interaction" json:"interaction" validate:"interaction"`
}

func Get(c echo.Context) error {
	var player Player
	c.Bind(&player)

	monsterHealth, err := actions.InitializeMonsterHealth()
	playerHealth, err := actions.InitializePlayerHealth()

	if err != nil {
		log.Println(err)
	}

	if player.Interaction == "attack" {
		dmg := actions.AttackMonster()
		monsterHealth = actions.DecreaseMonsterHealth(dmg)

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
