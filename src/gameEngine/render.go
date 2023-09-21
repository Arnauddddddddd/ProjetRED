package gameEngine

import "github.com/gen2brain/raylib-go/raylib"

func render(engine *EngineStruct) {
	rl.BeginDrawing()
	rl.ClearBackground(engine.bgColor)

	rl.BeginMode2D(engine.cam)

	drawScene(engine)

	rl.EndMode2D()
	rl.EndDrawing()
}