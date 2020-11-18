package main

import (
	"fmt"
)

const stackSize = 128

var stack [stackSize]int
var stackPointer int = 0

// Bytecode instructions
const (
	// General instructions
	INST_LITERAL = iota
	// Game instructions
	INST_SPAWN_PLAYER
	INST_SET_HEALTH
	INST_GET_HEALTH
	INST_SET_AMMO
	INST_GET_AMMO
	// Arithmetic instructions
	INST_ADD
	INST_MINUS
	INST_MULTIPLY
	INST_DIVIDE
)

// Player definition
type Player struct {
	health int
	ammo   int
}

var players []Player

// High-level API example
func spawnPlayer(health int, ammo int) {
	player := Player{
		health: health,
		ammo:   ammo,
	}
	players = append(players, player)
	fmt.Printf("New player #%v spawned!\n", len(players)-1)
}
func getHealth(playerID int) int {
	health := players[playerID].health
	fmt.Printf("Player #%v current health is: %v\n", playerID, health)
	return health
}
func setHealth(playerID int, health int) {
	fmt.Printf("Player #%v health set to: %v\n", playerID, health)
	players[playerID].health = health
}
func getAmmo(playerID int) int {
	ammo := players[playerID].ammo
	fmt.Printf("Player #%v current ammo is: %v\n", playerID, ammo)
	return ammo
}
func setAmmo(playerID int, ammo int) {
	fmt.Printf("Player #%v ammo set to: %v\n", playerID, ammo)
	players[playerID].ammo = ammo
}

// Stack control functions
func push(value int) {
	if stackPointer < stackSize {
		stack[stackPointer] = value
		stackPointer++
		return
	}
	panic("Cannot push() to full stack!")
}
func pop() int {
	if stackPointer > 0 {
		stackPointer--
		return stack[stackPointer]
	}
	panic("Cannot pop() from empty stack!")
}

func interpret(bytecode []byte) {
	for i := 0; i < len(bytecode); i++ {
		instruction := bytecode[i]
		switch instruction {
		case INST_LITERAL:
			i++
			value := bytecode[i]
			push(int(value))
			break
		case INST_SPAWN_PLAYER:
			health := pop()
			ammo := pop()
			spawnPlayer(health, ammo)
			break
		case INST_SET_HEALTH:
			playerID := pop()
			health := pop()
			setHealth(playerID, health)
			break
		case INST_SET_AMMO:
			playerID := pop()
			ammo := pop()
			setAmmo(playerID, ammo)
			break
		case INST_GET_HEALTH:
			playerID := pop()
			push(getHealth(playerID))
			break
		case INST_GET_AMMO:
			playerID := pop()
			push(getAmmo(playerID))
			break
		case INST_ADD:
			val1 := pop()
			val2 := pop()
			push(val1 + val2)
			break
		case INST_MINUS:
			val1 := pop()
			val2 := pop()
			push(val2 - val1)
			break
		case INST_MULTIPLY:
			val1 := pop()
			val2 := pop()
			push(val1 * val2)
			break
		case INST_DIVIDE:
			val1 := pop()
			val2 := pop()
			push(int(val2 / val1))
			break
		}
	}
}

func main() {
	fmt.Println("Starting virtual machine ..")
	// Sample program
	bytecode := []byte{
		INST_LITERAL, 100, // Health 100
		INST_LITERAL, 100, // Ammo 100
		INST_SPAWN_PLAYER,
		INST_LITERAL, 70, // New health 70
		INST_LITERAL, 0, // Player ID #0
		INST_SET_HEALTH,
		INST_LITERAL, 0, // Player ID #0
		INST_GET_AMMO,
		INST_LITERAL, 2, // literal value 2
		INST_DIVIDE,
		INST_LITERAL, 0, // Player ID #0
		INST_SET_AMMO, // Set ammo to 50 (100/2)
	}
	interpret(bytecode)
	fmt.Println("Program execution complete.")
}
