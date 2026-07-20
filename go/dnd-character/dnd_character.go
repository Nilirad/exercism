package dndcharacter

import (
	"math/rand"
	"slices"
)

const diceSides = 6
const numRolls = 4

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return (score - 10) >> 1
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := [numRolls]int{}
	for i := range numRolls {
		rolls[i] = rand.Intn(diceSides) + 1
	}

	slices.Sort(rolls[:])

	value := 0
	for i := 1; i < numRolls; i++ {
		value += rolls[i]
	}

	return value
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	hitpoints := 10 + Modifier(constitution)

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    hitpoints,
	}
}
