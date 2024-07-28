package actions

import (
	"context"
	"monsterXknight/config"
)

var ctx = context.Background()

func AttackPlayer() int {

	dmgValue := generateRandBetween(config.MonsterAttackMinDmg, config.MonsterAttackMaxDmg)

	currentPlayerHealth -= dmgValue

	return dmgValue
}
