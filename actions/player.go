package actions

import (
	"math/rand"
	"monsterXknight/config"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = config.MonsterHealth
var currentPlayerHealth = config.PlayerHealth

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
