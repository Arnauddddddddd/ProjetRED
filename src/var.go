package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1000
	screenHeight = 480
)

var (
	run = true
	bgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle

	playerMoving bool
	playerDir int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame int

	framCount int

	playerSpeed float32 = 3

	musicPaused bool
	music rl.Music

	cam rl.Camera2D

    buttonMenu rl.Texture2D
    buttonMenuPressed rl.Texture2D

)