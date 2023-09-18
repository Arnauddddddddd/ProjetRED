package main

import "github.com/gen2brain/raylib-go/raylib"

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bgColor)

	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}