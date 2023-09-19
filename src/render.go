package main

import "github.com/gen2brain/raylib-go/raylib"

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(engine.bgColor)

	rl.BeginMode2D(engine.cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}